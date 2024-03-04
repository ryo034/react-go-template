package auth

import (
	"context"
	"testing"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/datetime"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"

	"github.com/ryo034/react-go-template/apps/system/api/domain/me/provider"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/phone"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"

	"go.uber.org/mock/gomock"

	"github.com/ryo034/react-go-template/apps/system/api/domain/auth"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/notification"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
)

func Test_useCase_AuthByOAuth_OK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockTxProvider := core.NewMockTransactionProvider(ctrl)
	mockDbProvider := core.NewMockProvider(ctrl)
	mockAuthRepo := auth.NewMockRepository(ctrl)
	mockMeRepo := me.NewMockRepository(ctrl)
	mockInvRepo := invitation.NewMockRepository(ctrl)
	mockWRepo := workspace.NewMockRepository(ctrl)
	mockNotificationRepo := notification.NewMockRepository(ctrl)
	mockOutputPort := NewMockOutputPort(ctrl)

	uc := NewUseCase(mockTxProvider, mockDbProvider, mockAuthRepo, mockMeRepo, mockInvRepo, mockWRepo, mockNotificationRepo, mockOutputPort)

	aID, _ := account.NewID("018e060f-faec-73f0-a25f-ed83f41347ea")
	email, _ := account.NewEmail("test@example.com")
	name, _ := account.NewName("test")
	iph, _ := phone.NewInternationalPhoneNumber("+819012345678", "JP")
	usr := user.NewUser(aID, email, &name, &iph, nil)

	prID := provider.NewIDFromUUID(uuid.MustParse("018e060f-faec-73f0-a25f-ed83f41347ea"))
	prUID, _ := provider.NewUID("test")
	prv := provider.NewProvider(prID, provider.Email, "test", prUID)

	mockInput := ByOAuthInput{aID, CreateInfo{User: usr, Provider: prv}}

	tests := []struct {
		name    string
		input   ByOAuthInput
		setup   func()
		wantErr bool
	}{
		{
			"ログイン履歴がない場合、新規登録される",
			mockInput,
			func() {
				mockMe := me.NewMe(usr, nil, nil, nil, nil, provider.NewProviders([]*provider.Provider{prv}))

				mockDbProvider.EXPECT().GetExecutor(gomock.Any(), gomock.Any()).Return(nil)
				mockMeRepo.EXPECT().FindLastLogin(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil)
				mockAuthRepo.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(usr, nil)
				mockMeRepo.EXPECT().UpdateProfile(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				mockMeRepo.EXPECT().FindBeforeOnboard(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockMe, nil)
				mockOutputPort.EXPECT().AuthByAuth(gomock.Any())
			},
			false,
		},
		{
			"ログイン履歴があり、ワークスペース未所属の場合、ログイン履歴が記録されない",
			mockInput,
			func() {
				mockLastLogin := me.NewMe(usr, nil, nil, nil, nil, provider.NewProviders([]*provider.Provider{prv}))

				mockDbProvider.EXPECT().GetExecutor(gomock.Any(), gomock.Any()).Return(nil)
				mockMeRepo.EXPECT().FindLastLogin(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockLastLogin, nil)
				mockOutputPort.EXPECT().AuthByAuth(gomock.Any())
			},
			false,
		},
		{
			"ログイン履歴があり、ワークスペース所属済の場合、ログイン履歴が記録される",
			mockInput,
			func() {
				wID := workspace.NewIDFromUUID(uuid.MustParse("018e0622-57a9-755b-b55a-8ce20a8161e1"))
				wSubdomain, _ := workspace.NewSubdomain("test")
				wName, _ := workspace.NewName("test")
				wd := workspace.NewDetail(wName, wSubdomain)
				mockWorkspace := workspace.NewWorkspace(wID, wd)
				mID := member.NewIDFromUUID(uuid.MustParse("018e0624-2e0f-71e9-8296-2ba7f19063ef"))
				mockMember := member.NewMember(mID, usr, member.NewEmptyProfile(), member.RoleOwner)
				mockJoinedWorkspaces := workspace.NewWorkspaces([]*workspace.Workspace{mockWorkspace})
				mockLastLogin := me.NewMe(usr, mockWorkspace, mockMember, mockJoinedWorkspaces, nil, provider.NewProviders([]*provider.Provider{prv}))

				mockDbProvider.EXPECT().GetExecutor(gomock.Any(), gomock.Any()).Return(nil)
				mockMeRepo.EXPECT().FindLastLogin(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockLastLogin, nil)
				mockMeRepo.EXPECT().SetMe(gomock.Any(), gomock.Any()).Return(nil)
				mockMeRepo.EXPECT().RecordLogin(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				mockOutputPort.EXPECT().AuthByAuth(gomock.Any())
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			_, err := uc.AuthByOAuth(context.Background(), tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthByOAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_useCase_ProcessInvitationEmail_OK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockTxProvider := core.NewMockTransactionProvider(ctrl)
	mockDbProvider := core.NewMockProvider(ctrl)
	mockAuthRepo := auth.NewMockRepository(ctrl)
	mockMeRepo := me.NewMockRepository(ctrl)
	mockInvRepo := invitation.NewMockRepository(ctrl)
	mockWRepo := workspace.NewMockRepository(ctrl)
	mockNotificationRepo := notification.NewMockRepository(ctrl)
	mockOutputPort := NewMockOutputPort(ctrl)

	uc := NewUseCase(mockTxProvider, mockDbProvider, mockAuthRepo, mockMeRepo, mockInvRepo, mockWRepo, mockNotificationRepo, mockOutputPort)

	mockInputToken := invitation.NewToken(uuid.MustParse("018e062b-e4e7-7065-bf2a-926216c4700e"))
	mockEmail, _ := account.NewEmail("test@example.com")

	t.Run("Invalid input token", func(t *testing.T) {
		ctx := context.Background()
		invID := invitation.NewID(uuid.MustParse("018e062e-b742-706e-be1f-2827a5d73fbb"))
		mockToken := invitation.NewToken(uuid.MustParse("018e062f-334d-78c6-b008-6a66f553c67f"))
		events := invitation.NewEvents(make([]invitation.Event, 0))
		expiredAt := invitation.NewExpiredAt(datetime.Now())
		inviteeEmail, _ := account.NewEmail("test_invitee@exampel.com")
		displayName := member.NewDisplayName("test")
		mockFindActiveByEmail := invitation.NewInvitation(invID, mockToken, events, expiredAt, inviteeEmail, displayName)

		mockDbProvider.EXPECT().GetExecutor(gomock.Any(), gomock.Any()).Return(nil)
		mockInvRepo.EXPECT().FindActiveByEmail(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockFindActiveByEmail, nil)

		res, err := uc.ProcessInvitationEmail(ctx, ProcessInvitationEmailInput{Token: mockInputToken, Email: mockEmail})

		assert.Nil(t, res)
		assert.Error(t, err)
		assert.Equal(t, invitation.NewInvalidInviteToken(mockInputToken.Value()), err)
	})

	t.Run("Invalid input token can not verify", func(t *testing.T) {
		ctx := context.Background()
		invID := invitation.NewID(uuid.MustParse("018e062e-b742-706e-be1f-2827a5d73fbb"))
		events := invitation.NewEvents([]invitation.Event{invitation.NewEvent(invitation.Accepted, datetime.Now())})
		expiredAt := invitation.NewExpiredAt(datetime.Now())
		inviteeEmail, _ := account.NewEmail("test_invitee@exampel.com")
		displayName := member.NewDisplayName("test")
		mockFindActiveByEmail := invitation.NewInvitation(invID, mockInputToken, events, expiredAt, inviteeEmail, displayName)

		mockDbProvider.EXPECT().GetExecutor(gomock.Any(), gomock.Any()).Return(nil)
		mockInvRepo.EXPECT().FindActiveByEmail(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockFindActiveByEmail, nil)

		res, err := uc.ProcessInvitationEmail(ctx, ProcessInvitationEmailInput{Token: mockInputToken, Email: mockEmail})

		assert.Nil(t, res)
		assert.Error(t, err)
		assert.Equal(t, invitation.NewAlreadyAcceptedInvitation(invID, mockInputToken.Value()), err)
	})
}
