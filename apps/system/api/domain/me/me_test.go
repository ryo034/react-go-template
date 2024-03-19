package me

import (
	"reflect"
	"testing"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"

	"github.com/google/uuid"

	"github.com/ryo034/react-go-template/apps/system/api/domain/me/provider"
	user "github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	member "github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
)

func TestMe_UpdateWorkspace_OK(t *testing.T) {
	wID := workspace.NewIDFromUUID(uuid.MustParse("018e5650-c7e5-77ff-b019-1ff43e19ccea"))
	wName, _ := workspace.NewName("test")
	wSubdomain, _ := workspace.NewSubdomain("test")
	w := workspace.NewWorkspace(wID, workspace.NewDetail(wName, wSubdomain))

	uID, _ := account.NewID("018e5650-c7e5-7702-8a0a-ee9e7365fda6")
	email, _ := account.NewEmail("test@example.com")
	name, _ := account.NewName("test")
	self := user.NewUser(uID, email, &name, nil, nil)

	mID, _ := member.NewID("018e5650-c7e5-7d58-985c-234da258296a")
	mp := member.NewProfile(member.NewDisplayName("test"), nil, member.NewAsEmptyBio())

	mem := member.NewMember(mID, self, mp, member.RoleOwner, member.MembershipStatusActive)

	joinedWs := workspace.NewWorkspaces([]*workspace.Workspace{w})

	apID := provider.NewIDFromUUID(uuid.MustParse("018e5650-c7e5-7d74-a6af-2283571bd935"))
	apUID, _ := provider.NewUID("018e5650-c7e5-71a9-bc91-37cc7ce66e1a")
	ap := provider.NewProvider(apID, provider.Email, provider.ProvidedByFirebase, apUID)
	prs := provider.NewProviders([]*provider.Provider{ap})

	updatedName, _ := workspace.NewName("updated name")
	updatedSubdomain, _ := workspace.NewSubdomain("updated-subdomain")

	want := workspace.NewWorkspace(wID, workspace.NewDetail(updatedName, updatedSubdomain))

	type fields struct {
		self                *user.User
		workspace           *workspace.Workspace
		member              *member.Member
		joinedWorkspaces    workspace.Workspaces
		receivedInvitations ReceivedInvitations
		providers           provider.Providers
	}
	type args struct {
		wID       workspace.ID
		name      workspace.Name
		subdomain workspace.Subdomain
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *workspace.Workspace
		wantErr bool
	}{
		{
			"Update workspace with valid input",
			fields{
				self:                self,
				workspace:           w,
				member:              mem,
				joinedWorkspaces:    joinedWs,
				receivedInvitations: nil,
				providers:           prs,
			},
			args{
				wID:       wID,
				name:      updatedName,
				subdomain: updatedSubdomain,
			},
			want,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Me{tt.fields.self, tt.fields.workspace, tt.fields.member, tt.fields.joinedWorkspaces, tt.fields.receivedInvitations, tt.fields.providers}
			got, err := m.UpdateWorkspace(tt.args.wID, tt.args.name, tt.args.subdomain)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateWorkspace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateWorkspace() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMe_UpdateWorkspace_Error(t *testing.T) {
	wID := workspace.NewIDFromUUID(uuid.MustParse("018e5650-c7e5-77ff-b019-1ff43e19ccea"))
	wName, _ := workspace.NewName("test")
	wSubdomain, _ := workspace.NewSubdomain("test")
	w := workspace.NewWorkspace(wID, workspace.NewDetail(wName, wSubdomain))

	uID, _ := account.NewID("018e5650-c7e5-7702-8a0a-ee9e7365fda6")
	email, _ := account.NewEmail("test@example.com")
	name, _ := account.NewName("test")
	self := user.NewUser(uID, email, &name, nil, nil)

	mID, _ := member.NewID("018e5650-c7e5-7d58-985c-234da258296a")
	mp := member.NewProfile(member.NewDisplayName("test"), nil, member.NewAsEmptyBio())

	joinedWs := workspace.NewWorkspaces([]*workspace.Workspace{w})

	apID := provider.NewIDFromUUID(uuid.MustParse("018e5650-c7e5-7d74-a6af-2283571bd935"))
	apUID, _ := provider.NewUID("018e5650-c7e5-71a9-bc91-37cc7ce66e1a")
	ap := provider.NewProvider(apID, provider.Email, provider.ProvidedByFirebase, apUID)
	prs := provider.NewProviders([]*provider.Provider{ap})

	updatedName, _ := workspace.NewName("updated name")
	updatedSubdomain, _ := workspace.NewSubdomain("updated-subdomain")

	diffWorkspaceID := workspace.NewIDFromUUID(uuid.MustParse("018e5650-c7e5-77be-892b-4445beed5da5"))

	type fields struct {
		self                *user.User
		workspace           *workspace.Workspace
		member              *member.Member
		joinedWorkspaces    workspace.Workspaces
		receivedInvitations ReceivedInvitations
		providers           provider.Providers
	}
	type args struct {
		wID       workspace.ID
		name      workspace.Name
		subdomain workspace.Subdomain
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			"Can not update different workspace ID",
			fields{
				self:                self,
				workspace:           w,
				member:              member.NewMember(mID, self, mp, member.RoleOwner, member.MembershipStatusActive),
				joinedWorkspaces:    joinedWs,
				receivedInvitations: nil,
				providers:           prs,
			},
			args{
				wID:       diffWorkspaceID,
				name:      updatedName,
				subdomain: updatedSubdomain,
			},
			domainErr.NewForbidden("Cannot update workspace"),
		},
		{
			"Can not update workspace with Member role, receive forbidden error",
			fields{
				self:                self,
				workspace:           w,
				member:              member.NewMember(mID, self, mp, member.RoleMember, member.MembershipStatusActive),
				joinedWorkspaces:    joinedWs,
				receivedInvitations: nil,
				providers:           prs,
			},
			args{
				wID:       wID,
				name:      updatedName,
				subdomain: updatedSubdomain,
			},
			domainErr.NewForbidden("Can update only owner or admin"),
		},
		{
			"Can not update workspace with Guest role, receive forbidden error",
			fields{
				self:                self,
				workspace:           w,
				member:              member.NewMember(mID, self, mp, member.RoleGuest, member.MembershipStatusActive),
				joinedWorkspaces:    joinedWs,
				receivedInvitations: nil,
				providers:           prs,
			},
			args{
				wID:       wID,
				name:      updatedName,
				subdomain: updatedSubdomain,
			},
			domainErr.NewForbidden("Can update only owner or admin"),
		},
		{
			"Can not update Gone member, receive Gone error",
			fields{
				self:                self,
				workspace:           w,
				member:              member.NewMember(mID, self, mp, member.RoleOwner, member.MembershipStatusLeave),
				joinedWorkspaces:    joinedWs,
				receivedInvitations: nil,
				providers:           prs,
			},
			args{
				wID:       wID,
				name:      updatedName,
				subdomain: updatedSubdomain,
			},
			domainErr.NewGone("MemberID 018e5650-c7e5-7d58-985c-234da258296a"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Me{tt.fields.self, tt.fields.workspace, tt.fields.member, tt.fields.joinedWorkspaces, tt.fields.receivedInvitations, tt.fields.providers}
			got, err := m.UpdateWorkspace(tt.args.wID, tt.args.name, tt.args.subdomain)
			if err == nil {
				t.Errorf("UpdateWorkspace() got = %v, want %v", got, tt.wantErr)
				return
			} else {
				if !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("UpdateWorkspace() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestMe_ValidateCanLeave_Error(t *testing.T) {
	wID := workspace.NewIDFromUUID(uuid.MustParse("018e5650-c7e5-77ff-b019-1ff43e19ccea"))
	wName, _ := workspace.NewName("test")
	wSubdomain, _ := workspace.NewSubdomain("test")
	w := workspace.NewWorkspace(wID, workspace.NewDetail(wName, wSubdomain))

	uID, _ := account.NewID("018e5650-c7e5-7702-8a0a-ee9e7365fda6")
	email, _ := account.NewEmail("test@example.com")
	name, _ := account.NewName("test")
	self := user.NewUser(uID, email, &name, nil, nil)

	mID, _ := member.NewID("018e5650-c7e5-7d58-985c-234da258296a")
	mp := member.NewProfile(member.NewDisplayName("test"), nil, member.NewAsEmptyBio())

	joinedWs := workspace.NewWorkspaces([]*workspace.Workspace{w})

	apID := provider.NewIDFromUUID(uuid.MustParse("018e5650-c7e5-7d74-a6af-2283571bd935"))
	apUID, _ := provider.NewUID("018e5650-c7e5-71a9-bc91-37cc7ce66e1a")
	ap := provider.NewProvider(apID, provider.Email, provider.ProvidedByFirebase, apUID)
	prs := provider.NewProviders([]*provider.Provider{ap})

	type fields struct {
		self                *user.User
		workspace           *workspace.Workspace
		member              *member.Member
		joinedWorkspaces    workspace.Workspaces
		receivedInvitations ReceivedInvitations
		providers           provider.Providers
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			"ValidateCanLeave with Gone member, receive Gone error",
			fields{
				self:                self,
				workspace:           w,
				member:              member.NewMember(mID, self, mp, member.RoleOwner, member.MembershipStatusLeave),
				joinedWorkspaces:    joinedWs,
				receivedInvitations: nil,
				providers:           prs,
			},
			domainErr.NewGone("MemberID 018e5650-c7e5-7d58-985c-234da258296a"),
		},
		{
			"If not joined, receive Unauthenticated error",
			fields{self, nil, nil, nil, nil, prs},
			domainErr.NewUnauthenticated("Not joined"),
		},
		{
			"If Owner, receive Forbidden error",
			fields{
				self:                self,
				workspace:           w,
				member:              member.NewMember(mID, self, mp, member.RoleOwner, member.MembershipStatusActive),
				joinedWorkspaces:    joinedWs,
				receivedInvitations: nil,
				providers:           prs,
			},
			domainErr.NewForbidden("Cannot leave owner"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Me{tt.fields.self, tt.fields.workspace, tt.fields.member, tt.fields.joinedWorkspaces, tt.fields.receivedInvitations, tt.fields.providers}
			err := m.ValidateCanLeave()
			if err == nil {
				t.Errorf("ValidateCanLeave() got = %v, want %v", err, tt.wantErr)
				return
			} else {
				if !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("ValidateCanLeave() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
