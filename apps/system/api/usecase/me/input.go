package me

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type UpdateInput struct {
	me *me.Me
}

type UpdateProfileInput struct {
	user *user.User
}

type UpdateMemberProfileInput struct {
	AccountID account.ID
	Profile   member.Profile
}

func NewUpdateInput(i openapi.Me) (*UpdateInput, error) {
	aID := account.NewIDFromUUID(i.Self.UserId)
	email, err := account.NewEmail(i.Self.Email)
	if err != nil {
		return nil, err
	}
	var na *account.Name
	if i.Self.Name.Set {
		tmpNa, err := account.NewName(i.Self.Name.Value)
		if err != nil {
			return nil, err
		}
		na = &tmpNa
	}
	u := user.NewUser(aID, email, na, nil)

	var w *workspace.Workspace
	if i.CurrentWorkspace.Set {
		wID := workspace.NewIDFromUUID(i.CurrentWorkspace.Value.WorkspaceId)
		wn, err := workspace.NewName(i.CurrentWorkspace.Value.Name)
		if err != nil {
			return nil, err
		}
		wsd, err := workspace.NewSubdomain(i.CurrentWorkspace.Value.Subdomain)
		if err != nil {
			return nil, err
		}
		d := workspace.NewDetail(wn, wsd)
		w = workspace.NewWorkspace(wID, d)
	}

	var mem *member.Member
	if i.Member.Set {
		mID, err := member.NewID(i.Member.Value.ID)
		if err != nil {
			return nil, err
		}
		mdn := member.NewDisplayName(i.Member.Value.Profile.DisplayName)
		var mid *member.IDNumber
		if i.Member.Value.Profile.IdNumber.Set {
			tmpMid, err := member.NewIDNumber(i.Member.Value.Profile.IdNumber.Value)
			if err != nil {
				return nil, err
			}
			mid = &tmpMid
		}
		bio := member.NewAsEmptyBio()
		if i.Member.Value.Profile.Bio.Set {
			bio, err = member.NewBio(i.Member.Value.Profile.Bio.Value)
			if err != nil {
				return nil, err
			}
		}

		pr := member.NewProfile(mdn, mid, bio)
		mem = member.NewMember(mID, u, pr)
	}

	jws := make([]*workspace.Workspace, 0, len(i.JoinedWorkspaces))
	for _, jw := range i.JoinedWorkspaces {
		wID := workspace.NewIDFromUUID(jw.WorkspaceId)
		wn, err := workspace.NewName(jw.Name)
		if err != nil {
			return nil, err
		}
		wsd, err := workspace.NewSubdomain(jw.Subdomain)
		if err != nil {
			return nil, err
		}
		d := workspace.NewDetail(wn, wsd)
		jws = append(jws, workspace.NewWorkspace(wID, d))
	}

	m := me.NewMe(u, w, mem, workspace.NewWorkspaces(jws), nil)
	return &UpdateInput{me: m}, nil
}

func NewUpdateProfileInput(i openapi.User) (*UpdateProfileInput, error) {
	aID := account.NewIDFromUUID(i.UserId)
	email, err := account.NewEmail(i.Email)
	if err != nil {
		return nil, err
	}
	var na *account.Name
	if i.Name.Set {
		tmpNa, err := account.NewName(i.Name.Value)
		if err != nil {
			return nil, err
		}
		na = &tmpNa
	}
	u := user.NewUser(aID, email, na, nil)
	return &UpdateProfileInput{user: u}, nil
}

type AcceptInvitationInput struct {
	AccountID    account.ID
	InvitationID invitation.ID
}
