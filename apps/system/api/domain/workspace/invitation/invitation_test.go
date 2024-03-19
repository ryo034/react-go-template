package invitation

import (
	"reflect"
	"testing"
	"time"

	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/datetime"
	"github.com/ryo034/react-go-template/apps/system/api/util/test"

	"github.com/ryo034/react-go-template/apps/system/api/domain/user"

	"github.com/google/uuid"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
)

func TestInvitation_ValidateCanAccept_Error(t *testing.T) {
	dt := test.GetDefaultTime()
	defaultTime := datetime.NewDatetime(dt)

	tm, _ := time.Parse("2006-01-02 15:04:05", "2200-01-10 12:00:00")
	expTime := NewExpiredAt(datetime.NewDatetime(tm))

	iID := NewID(uuid.MustParse("018e5650-c7e5-7a6c-9aea-c969edd257d5"))
	token := NewToken(uuid.MustParse("018e5650-c7e5-768e-b776-caf501074e41"))
	inviteeEmail, _ := account.NewEmail("test@example.com")
	dn := member.NewDisplayName("test")

	uID, _ := account.NewID("018e5650-c7e5-7702-8a0a-ee9e7365fda6")
	email, _ := account.NewEmail("test@example.com")
	name, _ := account.NewName("test")
	usr := user.NewUser(uID, email, &name, nil, nil)
	mID, _ := member.NewID("018e5650-c7e5-7d58-985c-234da258296a")
	mp := member.NewProfile(member.NewDisplayName("test"), nil, member.NewAsEmptyBio())
	inviter := member.NewMember(mID, usr, mp, member.RoleOwner, member.MembershipStatusActive)

	odtm, _ := time.Parse("2006-01-02 15:04:05", "2000-01-10 12:00:00")
	overDeadline := NewExpiredAt(datetime.NewDatetime(odtm))

	type fields struct {
		id           ID
		token        Token
		event        *Event
		expiredAt    ExpiredAt
		inviteeEmail account.Email
		displayName  *member.DisplayName
		inviter      *member.Member
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			"Already accepted, receive NewAlreadyAcceptedInvitation error",
			fields{iID, token, &Event{Accepted, defaultTime}, expTime, inviteeEmail, dn, inviter},
			NewAlreadyAcceptedInvitation(iID, token.Value()),
		},
		{
			"Already revoked, receive NewAlreadyRevokedInvitation error",
			fields{iID, token, &Event{Revoked, defaultTime}, expTime, inviteeEmail, dn, inviter},
			NewAlreadyRevokedInvitation(iID, token.Value()),
		},
		{
			"Already expired, receive NewAlreadyExpiredInvitation error",
			fields{iID, token, nil, overDeadline, inviteeEmail, dn, inviter},
			NewAlreadyExpiredInvitation(iID, token.Value()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Invitation{tt.fields.id, tt.fields.token, tt.fields.event, tt.fields.expiredAt, tt.fields.inviteeEmail, tt.fields.displayName, tt.fields.inviter}
			err := i.ValidateCanAccept()
			if err == nil {
				t.Errorf("ValidateCanAccept() got = %v, want %v", err, tt.wantErr)
				return
			} else {
				if !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("ValidateCanAccept() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestInvitation_ValidateCanRevoke_Error(t *testing.T) {
	dt := test.GetDefaultTime()
	defaultTime := datetime.NewDatetime(dt)

	diffAccountID, _ := account.NewID("018e5650-c7e5-79f3-812a-9b56bf9b024e")

	tm, _ := time.Parse("2006-01-02 15:04:05", "2200-01-10 12:00:00")
	expTime := NewExpiredAt(datetime.NewDatetime(tm))

	iID := NewID(uuid.MustParse("018e5650-c7e5-7a6c-9aea-c969edd257d5"))
	token := NewToken(uuid.MustParse("018e5650-c7e5-768e-b776-caf501074e41"))
	inviteeEmail, _ := account.NewEmail("test@example.com")
	dn := member.NewDisplayName("test")

	uID, _ := account.NewID("018e5650-c7e5-7702-8a0a-ee9e7365fda6")
	email, _ := account.NewEmail("test@example.com")
	name, _ := account.NewName("test")
	usr := user.NewUser(uID, email, &name, nil, nil)
	mID, _ := member.NewID("018e5650-c7e5-7d58-985c-234da258296a")
	mp := member.NewProfile(member.NewDisplayName("test"), nil, member.NewAsEmptyBio())
	inviter := member.NewMember(mID, usr, mp, member.RoleOwner, member.MembershipStatusActive)

	odtm, _ := time.Parse("2006-01-02 15:04:05", "2000-01-10 12:00:00")
	overDeadline := NewExpiredAt(datetime.NewDatetime(odtm))

	type fields struct {
		id           ID
		token        Token
		event        *Event
		expiredAt    ExpiredAt
		inviteeEmail account.Email
		displayName  *member.DisplayName
		inviter      *member.Member
	}
	tests := []struct {
		name    string
		fields  fields
		args    account.ID
		wantErr error
	}{
		{
			"If the inviter is different, receive NewForbidden error",
			fields{iID, token, nil, expTime, inviteeEmail, dn, inviter},
			diffAccountID,
			domainErr.NewForbidden("revoke can only be done by the inviter"),
		},
		{
			"Already accepted, receive NewAlreadyAcceptedInvitation error",
			fields{iID, token, &Event{Accepted, defaultTime}, expTime, inviteeEmail, dn, inviter},
			diffAccountID,
			domainErr.NewForbidden("revoke can only be done by the inviter"),
		},
		{
			"Already revoked, receive NewAlreadyRevokedInvitation error",
			fields{iID, token, &Event{Revoked, defaultTime}, expTime, inviteeEmail, dn, inviter},
			uID,
			NewAlreadyRevokedInvitation(iID, token.Value()),
		},
		{
			"Already expired, receive NewAlreadyExpiredInvitation error",
			fields{iID, token, nil, overDeadline, inviteeEmail, dn, inviter},
			uID,
			NewAlreadyExpiredInvitation(iID, token.Value()),
		},
		{
			"Already verified, receive NewAlreadyAcceptedInvitation error",
			fields{iID, token, &Event{Verified, defaultTime}, expTime, inviteeEmail, dn, inviter},
			uID,
			NewAlreadyVerifiedInvitation(iID, token.Value()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Invitation{tt.fields.id, tt.fields.token, tt.fields.event, tt.fields.expiredAt, tt.fields.inviteeEmail, tt.fields.displayName, tt.fields.inviter}
			err := i.ValidateCanRevoke(tt.args)
			if err == nil {
				t.Errorf("ValidateCanRevoke() got = %v, want %v", err, tt.wantErr)
				return
			} else {
				if !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("ValidateCanRevoke() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestInvitation_ValidateCanResend_Error(t *testing.T) {
	dt := test.GetDefaultTime()
	defaultTime := datetime.NewDatetime(dt)

	diffAccountID, _ := account.NewID("018e5650-c7e5-79f3-812a-9b56bf9b024e")

	tm, _ := time.Parse("2006-01-02 15:04:05", "2200-01-10 12:00:00")
	expTime := NewExpiredAt(datetime.NewDatetime(tm))

	iID := NewID(uuid.MustParse("018e5650-c7e5-7a6c-9aea-c969edd257d5"))
	token := NewToken(uuid.MustParse("018e5650-c7e5-768e-b776-caf501074e41"))
	inviteeEmail, _ := account.NewEmail("test@example.com")
	dn := member.NewDisplayName("test")

	uID, _ := account.NewID("018e5650-c7e5-7702-8a0a-ee9e7365fda6")
	email, _ := account.NewEmail("test@example.com")
	name, _ := account.NewName("test")
	usr := user.NewUser(uID, email, &name, nil, nil)
	mID, _ := member.NewID("018e5650-c7e5-7d58-985c-234da258296a")
	mp := member.NewProfile(member.NewDisplayName("test"), nil, member.NewAsEmptyBio())
	inviter := member.NewMember(mID, usr, mp, member.RoleOwner, member.MembershipStatusActive)

	type fields struct {
		id           ID
		token        Token
		event        *Event
		expiredAt    ExpiredAt
		inviteeEmail account.Email
		displayName  *member.DisplayName
		inviter      *member.Member
	}
	tests := []struct {
		name    string
		fields  fields
		args    account.ID
		wantErr error
	}{
		{
			"If the inviter is different, receive NewForbidden error",
			fields{iID, token, nil, expTime, inviteeEmail, dn, inviter},
			diffAccountID,
			domainErr.NewForbidden("resend can only be done by the inviter"),
		},
		{
			"Already accepted, receive NewAlreadyAcceptedInvitation error",
			fields{iID, token, &Event{Accepted, defaultTime}, expTime, inviteeEmail, dn, inviter},
			uID,
			NewAlreadyAcceptedInvitation(iID, token.Value()),
		},
		{
			"Already revoked, receive NewAlreadyRevokedInvitation error",
			fields{iID, token, &Event{Revoked, defaultTime}, expTime, inviteeEmail, dn, inviter},
			uID,
			NewAlreadyRevokedInvitation(iID, token.Value()),
		},
		{
			"Already verified, receive NewAlreadyAcceptedInvitation error",
			fields{iID, token, &Event{Verified, defaultTime}, expTime, inviteeEmail, dn, inviter},
			uID,
			NewAlreadyVerifiedInvitation(iID, token.Value()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Invitation{tt.fields.id, tt.fields.token, tt.fields.event, tt.fields.expiredAt, tt.fields.inviteeEmail, tt.fields.displayName, tt.fields.inviter}
			err := i.ValidateCanResend(tt.args)
			if err == nil {
				t.Errorf("ValidateCanResend() got = %v, want %v", err, tt.wantErr)
				return
			} else {
				if !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("ValidateCanResend() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestInvitation_ValidateCanResend_OK(t *testing.T) {
	iID := NewID(uuid.MustParse("018e5650-c7e5-7a6c-9aea-c969edd257d5"))
	token := NewToken(uuid.MustParse("018e5650-c7e5-768e-b776-caf501074e41"))
	inviteeEmail, _ := account.NewEmail("test@example.com")
	dn := member.NewDisplayName("test")

	uID, _ := account.NewID("018e5650-c7e5-7702-8a0a-ee9e7365fda6")
	email, _ := account.NewEmail("test@example.com")
	name, _ := account.NewName("test")
	usr := user.NewUser(uID, email, &name, nil, nil)
	mID, _ := member.NewID("018e5650-c7e5-7d58-985c-234da258296a")
	mp := member.NewProfile(member.NewDisplayName("test"), nil, member.NewAsEmptyBio())
	inviter := member.NewMember(mID, usr, mp, member.RoleOwner, member.MembershipStatusActive)

	odtm, _ := time.Parse("2006-01-02 15:04:05", "2000-01-10 12:00:00")
	overDeadline := NewExpiredAt(datetime.NewDatetime(odtm))

	type fields struct {
		id           ID
		token        Token
		event        *Event
		expiredAt    ExpiredAt
		inviteeEmail account.Email
		displayName  *member.DisplayName
		inviter      *member.Member
	}
	tests := []struct {
		name   string
		fields fields
		args   account.ID
	}{
		{
			"Already expired, can resend",
			fields{iID, token, nil, overDeadline, inviteeEmail, dn, inviter},
			uID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Invitation{tt.fields.id, tt.fields.token, tt.fields.event, tt.fields.expiredAt, tt.fields.inviteeEmail, tt.fields.displayName, tt.fields.inviter}
			err := i.ValidateCanResend(tt.args)
			if err != nil {
				t.Errorf("ValidateCanResend() got = %v, want nil", err)
			}
		})
	}
}

func TestInvitation_ValidateCanVerify_Error(t *testing.T) {
	dt := test.GetDefaultTime()
	defaultTime := datetime.NewDatetime(dt)

	diffToken := NewToken(uuid.MustParse("018e5650-c7e5-71c9-918d-0c8bf0017a9e"))

	tm, _ := time.Parse("2006-01-02 15:04:05", "2200-01-10 12:00:00")
	expTime := NewExpiredAt(datetime.NewDatetime(tm))

	iID := NewID(uuid.MustParse("018e5650-c7e5-7a6c-9aea-c969edd257d5"))
	token := NewToken(uuid.MustParse("018e5650-c7e5-768e-b776-caf501074e41"))
	inviteeEmail, _ := account.NewEmail("test@example.com")
	dn := member.NewDisplayName("test")

	uID, _ := account.NewID("018e5650-c7e5-7702-8a0a-ee9e7365fda6")
	email, _ := account.NewEmail("test@example.com")
	name, _ := account.NewName("test")
	usr := user.NewUser(uID, email, &name, nil, nil)
	mID, _ := member.NewID("018e5650-c7e5-7d58-985c-234da258296a")
	mp := member.NewProfile(member.NewDisplayName("test"), nil, member.NewAsEmptyBio())
	inviter := member.NewMember(mID, usr, mp, member.RoleOwner, member.MembershipStatusActive)

	odtm, _ := time.Parse("2006-01-02 15:04:05", "2000-01-10 12:00:00")
	overDeadline := NewExpiredAt(datetime.NewDatetime(odtm))

	type fields struct {
		id           ID
		token        Token
		event        *Event
		expiredAt    ExpiredAt
		inviteeEmail account.Email
		displayName  *member.DisplayName
		inviter      *member.Member
	}
	tests := []struct {
		name    string
		fields  fields
		args    Token
		wantErr error
	}{
		{
			"If the different token, receive NewForbidden error",
			fields{iID, token, nil, expTime, inviteeEmail, dn, inviter},
			diffToken,
			NewInvalidInviteToken(diffToken.Value()),
		},
		{
			"Already accepted, receive NewAlreadyAcceptedInvitation error",
			fields{iID, token, &Event{Accepted, defaultTime}, expTime, inviteeEmail, dn, inviter},
			token,
			NewAlreadyAcceptedInvitation(iID, token.Value()),
		},
		{
			"Already revoked, receive NewAlreadyRevokedInvitation error",
			fields{iID, token, &Event{Revoked, defaultTime}, expTime, inviteeEmail, dn, inviter},
			token,
			NewAlreadyRevokedInvitation(iID, token.Value()),
		},
		{
			"Already verified, receive NewAlreadyAcceptedInvitation error",
			fields{iID, token, &Event{Verified, defaultTime}, expTime, inviteeEmail, dn, inviter},
			token,
			NewAlreadyVerifiedInvitation(iID, token.Value()),
		},
		{
			"Already expired, receive NewAlreadyExpiredInvitation error",
			fields{iID, token, nil, overDeadline, inviteeEmail, dn, inviter},
			token,
			NewAlreadyExpiredInvitation(iID, token.Value()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Invitation{tt.fields.id, tt.fields.token, tt.fields.event, tt.fields.expiredAt, tt.fields.inviteeEmail, tt.fields.displayName, tt.fields.inviter}
			err := i.ValidateCanVerify(tt.args)
			if err == nil {
				t.Errorf("ValidateCanVerify() got = %v, want %v", err, tt.wantErr)
				return
			} else {
				if !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("ValidateCanVerify() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestInvitation_ValidateCanGetByToken_Error(t *testing.T) {
	dt := test.GetDefaultTime()
	defaultTime := datetime.NewDatetime(dt)

	tm, _ := time.Parse("2006-01-02 15:04:05", "2200-01-10 12:00:00")
	expTime := NewExpiredAt(datetime.NewDatetime(tm))

	iID := NewID(uuid.MustParse("018e5650-c7e5-7a6c-9aea-c969edd257d5"))
	token := NewToken(uuid.MustParse("018e5650-c7e5-768e-b776-caf501074e41"))
	inviteeEmail, _ := account.NewEmail("test@example.com")
	dn := member.NewDisplayName("test")

	uID, _ := account.NewID("018e5650-c7e5-7702-8a0a-ee9e7365fda6")
	email, _ := account.NewEmail("test@example.com")
	name, _ := account.NewName("test")
	usr := user.NewUser(uID, email, &name, nil, nil)
	mID, _ := member.NewID("018e5650-c7e5-7d58-985c-234da258296a")
	mp := member.NewProfile(member.NewDisplayName("test"), nil, member.NewAsEmptyBio())
	inviter := member.NewMember(mID, usr, mp, member.RoleOwner, member.MembershipStatusActive)

	odtm, _ := time.Parse("2006-01-02 15:04:05", "2000-01-10 12:00:00")
	overDeadline := NewExpiredAt(datetime.NewDatetime(odtm))

	type fields struct {
		id           ID
		token        Token
		event        *Event
		expiredAt    ExpiredAt
		inviteeEmail account.Email
		displayName  *member.DisplayName
		inviter      *member.Member
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			"Already accepted, receive NewAlreadyAcceptedInvitation error",
			fields{iID, token, &Event{Accepted, defaultTime}, expTime, inviteeEmail, dn, inviter},
			NewAlreadyAcceptedInvitation(iID, token.Value()),
		},
		{
			"Already revoked, receive NewAlreadyRevokedInvitation error",
			fields{iID, token, &Event{Revoked, defaultTime}, expTime, inviteeEmail, dn, inviter},
			NewAlreadyRevokedInvitation(iID, token.Value()),
		},
		{
			"Already expired, receive NewAlreadyExpiredInvitation error",
			fields{iID, token, nil, overDeadline, inviteeEmail, dn, inviter},
			NewAlreadyExpiredInvitation(iID, token.Value()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Invitation{tt.fields.id, tt.fields.token, tt.fields.event, tt.fields.expiredAt, tt.fields.inviteeEmail, tt.fields.displayName, tt.fields.inviter}
			err := i.ValidateCanGetByToken()
			if err == nil {
				t.Errorf("ValidateCanGetByToken() got = %v, want %v", err, tt.wantErr)
				return
			} else {
				if !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("ValidateCanGetByToken() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
