package workspace

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/me"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/member"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
)

type UseCase interface {
	Create(ctx context.Context, i *CreateInput) (openapi.APIV1WorkspacesPostRes, error)
}

type useCase struct {
	txp    core.TransactionProvider
	dbp    core.Provider
	repo   workspace.Repository
	meRepo me.Repository
	op     OutputPort
}

func NewUseCase(txp core.TransactionProvider, dbp core.Provider, repo workspace.Repository, meRepo me.Repository, op OutputPort) UseCase {
	return &useCase{txp, dbp, repo, meRepo, op}
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
		_, err = u.repo.AddMember(pr, p, wres, m)
		if err != nil {
			return nil, err
		}

		reMeRes, err := u.meRepo.Find(pr, p, meRes.Member().ID())
		if err != nil {
			return nil, err
		}
		if err = u.meRepo.LastLogin(pr, p, reMeRes); err != nil {
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
