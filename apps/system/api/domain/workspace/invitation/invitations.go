package invitation

import "sort"

type Invitations interface {
	Size() int
	AsSlice() []*Invitation
	IsEmpty() bool
	IsNotEmpty() bool
	ExcludeRevoked() Invitations
	ExcludeVerified() Invitations
	OnlyAccepted() Invitations
	SortByExpiryAt() Invitations
}

type workspaces struct {
	wrapped []*Invitation
}

func NewInvitations(wrapped []*Invitation) Invitations {
	return &workspaces{wrapped}
}

func (ws *workspaces) IsEmpty() bool {
	return len(ws.wrapped) == 0
}

func (ws *workspaces) IsNotEmpty() bool {
	return !ws.IsEmpty()
}

func (ws *workspaces) Size() int {
	return len(ws.wrapped)
}

func (ws *workspaces) AsSlice() []*Invitation {
	return append(make([]*Invitation, 0, ws.Size()), ws.wrapped...)
}

func (ws *workspaces) ExcludeRevoked() Invitations {
	filtered := make([]*Invitation, 0)
	for _, w := range ws.wrapped {
		if !w.IsRevoked() {
			filtered = append(filtered, w)
		}
	}
	return NewInvitations(filtered)
}

func (ws *workspaces) ExcludeVerified() Invitations {
	filtered := make([]*Invitation, 0)
	for _, w := range ws.wrapped {
		if !w.IsVerified() {
			filtered = append(filtered, w)
		}
	}
	return NewInvitations(filtered)
}

func (ws *workspaces) OnlyAccepted() Invitations {
	filtered := make([]*Invitation, 0)
	for _, w := range ws.wrapped {
		if w.IsAccepted() {
			filtered = append(filtered, w)
		}
	}
	return NewInvitations(filtered)
}

func (ws *workspaces) SortByExpiryAt() Invitations {
	sorted := append(make([]*Invitation, 0, ws.Size()), ws.wrapped...)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].ExpiredAt().Value().ToTime().Before(sorted[j].ExpiredAt().Value().ToTime())
	})
	return NewInvitations(sorted)
}
