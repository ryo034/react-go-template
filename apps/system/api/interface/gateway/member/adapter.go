package member

import (
	"fmt"

	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	userGw "github.com/ryo034/react-go-template/apps/system/api/interface/gateway/user"
)

type Adapter interface {
	Adapt(m *models.Member) (*member.Member, error)
	AdaptAll(ms models.Members) (member.Members, error)
}

type adapter struct {
	uga userGw.Adapter
}

func NewAdapter(uga userGw.Adapter) Adapter {
	return &adapter{uga}
}

func (a *adapter) adaptRole(r *models.MemberRole) (member.Role, error) {
	switch r.Role {
	case "owner":
		return member.RoleOwner, nil
	case "admin":
		return member.RoleAdmin, nil
	case "member":
		return member.RoleMember, nil
	case "guest":
		return member.RoleGuest, nil
	}
	return "", fmt.Errorf("invalid role: %s", r.Role)
}

func (a *adapter) Adapt(m *models.Member) (*member.Member, error) {
	u, err := a.uga.AdaptTmp(m.Account)
	if err != nil {
		return nil, err
	}
	idNumber, err := member.NewIDNumber(m.Profile.MemberProfile.MemberIDNumber)
	if err != nil {
		return nil, err
	}

	dn := member.NewDisplayName(m.Profile.MemberProfile.DisplayName)
	if err != nil {
		return nil, err
	}
	id := member.NewIDFromUUID(m.MemberID)
	bio, err := member.NewBio(m.Profile.MemberProfile.Bio)
	if err != nil {
		return nil, err
	}

	pro := member.NewProfile(dn, &idNumber, bio)
	ar, err := a.adaptRole(m.Role.MemberRole)
	if err != nil {
		return nil, err
	}

	membershipEventType := member.MembershipStatusActive
	if m.MembershipEvent.MembershipEvent.EventType == "leave" {
		membershipEventType = member.MembershipStatusLeave
	}

	return member.NewMember(id, u, pro, ar, membershipEventType), nil
}

func (a *adapter) AdaptAll(ms models.Members) (member.Members, error) {
	mws := make([]*member.Member, 0, len(ms))
	for _, m := range ms {
		aw, err := a.Adapt(m)
		if err != nil {
			return nil, err
		}
		mws = append(mws, aw)
	}
	return member.NewMembers(mws), nil
}
