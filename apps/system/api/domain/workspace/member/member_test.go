package member

import (
	"reflect"
	"testing"

	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"

	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
)

func TestMember_UpdateRole_OK(t *testing.T) {
	type fields struct {
		id               ID
		u                *user.User
		profile          Profile
		role             Role
		membershipStatus MembershipStatus
	}
	type args struct {
		role Role
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Member
		wantErr bool
	}{
		{
			name:   "Admin can update to Member Role",
			fields: fields{id: ID{}, u: &user.User{}, profile: Profile{}, role: RoleAdmin, membershipStatus: MembershipStatusActive},
			args:   args{RoleMember},
			want:   &Member{ID{}, &user.User{}, Profile{}, RoleMember, MembershipStatusActive},
		},
		{
			name:   "Admin can update to Guest Role",
			fields: fields{id: ID{}, u: &user.User{}, profile: Profile{}, role: RoleAdmin, membershipStatus: MembershipStatusActive},
			args:   args{RoleGuest},
			want:   &Member{ID{}, &user.User{}, Profile{}, RoleGuest, MembershipStatusActive},
		},
		{
			name:   "Member can update to Admin Role",
			fields: fields{id: ID{}, u: &user.User{}, profile: Profile{}, role: RoleMember, membershipStatus: MembershipStatusActive},
			args:   args{RoleAdmin},
			want:   &Member{ID{}, &user.User{}, Profile{}, RoleAdmin, MembershipStatusActive},
		},
		{
			name:   "Member can update to Guest Role",
			fields: fields{id: ID{}, u: &user.User{}, profile: Profile{}, role: RoleMember, membershipStatus: MembershipStatusActive},
			args:   args{RoleGuest},
			want:   &Member{ID{}, &user.User{}, Profile{}, RoleGuest, MembershipStatusActive},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Member{tt.fields.id, tt.fields.u, tt.fields.profile, tt.fields.role, tt.fields.membershipStatus}
			got, err := w.UpdateRole(tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateRole() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMember_UpdateRole_Validate(t *testing.T) {
	type fields struct {
		id               ID
		u                *user.User
		profile          Profile
		role             Role
		membershipStatus MembershipStatus
	}

	tests := []struct {
		name    string
		fields  fields
		value   Role
		wantErr error
	}{
		{"Can not update to Same Role",
			fields{id: ID{}, u: &user.User{}, profile: Profile{}, role: RoleAdmin, membershipStatus: MembershipStatusActive},
			RoleAdmin,
			domainErr.NewBadRequest("the role is already the same"),
		},
		{"Can not update to Owner Role if Owner",
			fields{id: ID{}, u: &user.User{}, profile: Profile{}, role: RoleOwner, membershipStatus: MembershipStatusActive},
			RoleOwner,
			domainErr.NewForbidden("cannot change the role to owner"),
		},
		{"Can not update to Admin Role if Owner",
			fields{id: ID{}, u: &user.User{}, profile: Profile{}, role: RoleOwner, membershipStatus: MembershipStatusActive},
			RoleAdmin,
			domainErr.NewForbidden("cannot change the role"),
		},
		{"Can not update to Admin Role if Owner",
			fields{id: ID{}, u: &user.User{}, profile: Profile{}, role: RoleOwner, membershipStatus: MembershipStatusActive},
			RoleGuest,
			domainErr.NewForbidden("cannot change the role"),
		},
		{"Can not update to Member Role if Member",
			fields{id: ID{}, u: &user.User{}, profile: Profile{}, role: RoleOwner, membershipStatus: MembershipStatusActive},
			RoleMember,
			domainErr.NewForbidden("cannot change the role"),
		},
		{"Can not update to Guest Role if Guest",
			fields{id: ID{}, u: &user.User{}, profile: Profile{}, role: RoleOwner, membershipStatus: MembershipStatusActive},
			RoleGuest,
			domainErr.NewForbidden("cannot change the role"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Member{tt.fields.id, tt.fields.u, tt.fields.profile, tt.fields.role, tt.fields.membershipStatus}
			_, err := w.UpdateRole(tt.value)
			if err == nil {
				t.Errorf("UpdateRole() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				if !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("UpdateRole() got = %v, want %v", err, tt.wantErr)
				}
			}
		})
	}
}
