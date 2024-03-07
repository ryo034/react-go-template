package me

import (
	"context"
	"fmt"
	"io"
	"mime"
	"net/textproto"
	"path/filepath"

	domainErr "github.com/ryo034/react-go-template/apps/system/api/domain/shared/error"

	"github.com/ryo034/react-go-template/apps/system/api/domain/user"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/media"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"

	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"

	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	infraShared "github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"github.com/ryo034/react-go-template/apps/system/api/interface/presenter/shared"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	meUc "github.com/ryo034/react-go-template/apps/system/api/usecase/me"
)

type Controller interface {
	Find(ctx context.Context) (openapi.APIV1MeGetRes, error)
	AcceptInvitation(ctx context.Context, i AcceptInvitationInput) (openapi.AcceptInvitationRes, error)
	UpdateName(ctx context.Context, i UpdateProfileInput) (openapi.APIV1MeProfilePutRes, error)
	UpdateMemberProfile(ctx context.Context, i UpdateMemberProfileInput) (openapi.APIV1MeMemberProfilePutRes, error)
	UpdateProfilePhoto(ctx context.Context, i UpdateProfilePhotoInput) (openapi.APIV1MeProfilePhotoPutRes, error)
	RemoveProfilePhoto(ctx context.Context) (openapi.APIV1MeProfilePhotoDeleteRes, error)
}

type controller struct {
	uc   meUc.UseCase
	resl shared.Resolver
	co   infraShared.ContextOperator
}

type AcceptInvitationInput struct {
	InvitationID uuid.UUID
}

type UpdateProfileInput struct {
	Name string
}

type UpdateMemberProfileInput struct {
	DisplayName string
	IdNumber    string
	Bio         string
}

type UpdateProfilePhotoInput struct {
	File   io.Reader
	Name   string
	Header textproto.MIMEHeader
}

func NewController(uc meUc.UseCase, resl shared.Resolver, co infraShared.ContextOperator) Controller {
	return &controller{uc, resl, co}
}

func (c *controller) Find(ctx context.Context) (openapi.APIV1MeGetRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return nil, err
	}
	res, err := c.uc.Find(ctx, aID)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1MeGetRes), nil
	}
	return res, nil
}

func (c *controller) AcceptInvitation(ctx context.Context, i AcceptInvitationInput) (openapi.AcceptInvitationRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.AcceptInvitationRes), nil
	}
	in := meUc.AcceptInvitationInput{AccountID: aID, InvitationID: invitation.NewID(i.InvitationID)}
	res, err := c.uc.AcceptInvitation(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.AcceptInvitationRes), nil
	}
	return res, nil
}

func NewUpdateNameInput(i UpdateProfileInput, aID account.ID) (meUc.UpdateNameInput, error) {
	na, err := account.NewName(i.Name)
	if err != nil {
		return meUc.UpdateNameInput{}, err
	}
	return meUc.UpdateNameInput{Name: na, AccountID: aID}, nil
}

func (c *controller) UpdateName(ctx context.Context, i UpdateProfileInput) (openapi.APIV1MeProfilePutRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1MeProfilePutRes), nil
	}
	in, err := NewUpdateNameInput(i, aID)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1MeProfilePutRes), nil
	}
	res, err := c.uc.UpdateName(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1MeProfilePutRes), nil
	}
	return res, nil
}

func NewUpdateMemberProfileInput(i UpdateMemberProfileInput, aID account.ID) (meUc.UpdateMemberProfileInput, error) {
	var err error = nil
	var dn *member.DisplayName
	if i.DisplayName != "" {
		dn = member.NewDisplayName(i.DisplayName)
	}
	var mid *member.IDNumber
	if i.IdNumber != "" {
		tmpMID, err := member.NewIDNumber(i.IdNumber)
		if err != nil {
			return meUc.UpdateMemberProfileInput{}, err
		}
		mid = &tmpMID
	}
	bio := member.NewAsEmptyBio()
	if i.Bio != "" {
		bio, err = member.NewBio(i.Bio)
		if err != nil {
			return meUc.UpdateMemberProfileInput{}, err
		}
	}
	pr := member.NewProfile(dn, mid, bio)
	return meUc.UpdateMemberProfileInput{AccountID: aID, Profile: pr}, nil
}

func (c *controller) UpdateMemberProfile(ctx context.Context, i UpdateMemberProfileInput) (openapi.APIV1MeMemberProfilePutRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1MeMemberProfilePutRes), nil
	}
	in, err := NewUpdateMemberProfileInput(i, aID)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1MeMemberProfilePutRes), nil
	}
	res, err := c.uc.UpdateMemberProfile(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1MeMemberProfilePutRes), nil
	}
	return res, nil
}

func (c *controller) UpdateProfilePhoto(ctx context.Context, i UpdateProfilePhotoInput) (openapi.APIV1MeProfilePhotoPutRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1MeProfilePhotoPutRes), nil
	}

	_, params, err := mime.ParseMediaType(i.Header.Get("Content-Disposition"))
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1MeProfilePhotoPutRes), nil
	}
	filename := params["filename"]
	ext := filepath.Ext(filename)

	tmpPhoID, _ := uuid.NewV7()
	phoID := media.NewIDFromUUID(tmpPhoID)

	var avatarExt user.AvatarExt
	switch ext {
	case ".jpeg":
		avatarExt = user.AvatarExtJpeg
	case ".jpg":
		avatarExt = user.AvatarExtJpg
	case ".png":
		avatarExt = user.AvatarExtPng
	}

	if avatarExt == "" {
		return c.resl.Error(ctx, domainErr.NewBadRequest(fmt.Sprintf("invalid file type: %s", ext))).(openapi.APIV1MeProfilePhotoPutRes), nil
	}

	in := meUc.UpdateProfilePhotoInput{
		AccountID: aID,
		PhotoID:   phoID,
		File:      i.File,
		Ext:       avatarExt,
		Size:      -1,
	}
	res, err := c.uc.UpdateProfilePhoto(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1MeProfilePhotoPutRes), nil
	}
	return res, nil
}

func (c *controller) RemoveProfilePhoto(ctx context.Context) (openapi.APIV1MeProfilePhotoDeleteRes, error) {
	aID, err := c.co.GetUID(ctx)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1MeProfilePhotoDeleteRes), nil
	}
	in := meUc.RemoveProfilePhotoInput{AccountID: aID}
	res, err := c.uc.RemoveProfilePhoto(ctx, in)
	if err != nil {
		return c.resl.Error(ctx, err).(openapi.APIV1MeProfilePhotoDeleteRes), nil
	}
	return res, nil
}
