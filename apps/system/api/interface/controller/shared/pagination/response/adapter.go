package response

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/pagination/offset"
	paginationPb "github.com/ryo034/react-go-template/apps/system/api/schema/pb/pagination/v1"
)

type Adapter interface {
	AdaptOffset(p *offset.Pagination) (*paginationPb.OffsetPaginationResponse, error)
}

func NewAdapter() Adapter {
	return &adapter{}
}

type adapter struct {
}

func (a *adapter) AdaptOffset(p *offset.Pagination) (*paginationPb.OffsetPaginationResponse, error) {
	return &paginationPb.OffsetPaginationResponse{
		Limit:       p.Limit().ToUint32(),
		Page:        p.NextPage().ToUint32(),
		Total:       p.Total().ToUint32(),
		HasNextPage: p.HasNextPage(),
	}, nil
}
