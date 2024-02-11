package member

type InvitedBy struct {
	id   ID
	name DisplayName
}

func NewInvitedBy(id ID, name DisplayName) InvitedBy {
	return InvitedBy{id, name}
}

func (m *InvitedBy) ID() ID {
	return m.id
}

func (m *InvitedBy) Name() DisplayName {
	return m.name
}
