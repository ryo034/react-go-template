package workspace

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	fbDr "github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type UseCase interface {
	Create(ctx context.Context, i *CreateInput) (openapi.APIV1WorkspacesPostRes, error)
	FindAllMembers(ctx context.Context, i *FindAllMembersInput) (openapi.APIV1MembersGetRes, error)
}

type useCase struct {
	txp      core.TransactionProvider
	dbp      core.Provider
	repo     workspace.Repository
	meRepo   me.Repository
	fbDriver fbDr.Driver
	op       OutputPort
}

func NewUseCase(txp core.TransactionProvider, dbp core.Provider, repo workspace.Repository, meRepo me.Repository, fbDriver fbDr.Driver, op OutputPort) UseCase {
	return &useCase{txp, dbp, repo, meRepo, fbDriver, op}
}

func (u *useCase) Create(ctx context.Context, i *CreateInput) (openapi.APIV1WorkspacesPostRes, error) {
	p := u.dbp.GetExecutor(ctx, false)
	pr, err := u.txp.Provide(ctx)
	if err != nil {
		return nil, err
	}
	fn := func() (*workspace.Workspace, error) {
		w := i.Workspace()
		meRes, err := u.meRepo.FindBeforeOnboard(pr, p, i.AccountID())
		if err != nil {
			return nil, err
		}
		wres, err := u.repo.Create(pr, p, w)
		if err != nil {
			return nil, err
		}
		meID, err := member.GenerateID()
		if err != nil {
			return nil, err
		}
		dn, err := member.NewDisplayName(meRes.Self().Name().ToString())
		if err != nil {
			return nil, err
		}

		m := member.NewMember(meID, meRes.Self(), dn, nil)
		memRes, err := u.repo.AddMember(pr, p, wres, m)
		if err != nil {
			return nil, err
		}

		meRes, err = u.meRepo.Find(pr, p, memRes.ID())
		if err != nil {
			return nil, err
		}
		if err = u.meRepo.LastLogin(pr, p, meRes); err != nil {
			return nil, err
		}
		return wres, nil
	}
	result := pr.Transactional(fn)()
	if err = result.Error(); err != nil {
		return nil, err
	}
	res := result.Value(0).(*workspace.Workspace)
	return u.op.Create(res), nil
}

func (u *useCase) FindAllMembers(ctx context.Context, i *FindAllMembersInput) (openapi.APIV1MembersGetRes, error) {
	exec := u.dbp.GetExecutor(ctx, true)
	currentWorkspaceID, err := u.fbDriver.GetCurrentWorkspaceFromCustomClaim(ctx, i.accountID)
	if err != nil {
		return nil, err
	}
	ms, err := u.repo.FindAllMembers(ctx, exec, *currentWorkspaceID)
	if err != nil {
		return nil, err
	}
	return u.op.FindAllMembers(ms), nil
}
