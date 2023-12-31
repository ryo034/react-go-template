package request

import (
	"fmt"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/pagination"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/pagination/offset"
	paginationPb "github.com/ryo034/react-go-template/apps/system/api/schema/pb/pagination/v1"
)

type Adapter interface {
	AdaptOffset(p *paginationPb.OffsetPaginationRequest) (*offset.Pagination, error)
}

func NewAdapter() Adapter {
	return &adapter{}
}

type adapter struct {
}

func (a *adapter) AdaptOffset(p *paginationPb.OffsetPaginationRequest) (*offset.Pagination, error) {
	if p.Page < 0 {
		return nil, fmt.Errorf("page must be greater than or equal to 0")
	}
	if p.Limit < 0 {
		return nil, fmt.Errorf("limit must be greater than or equal to 0")
	}

	l, err := pagination.NewLimit(p.Limit)
	if err != nil {
		return nil, err
	}
	o, err := offset.NewPage(p.Page)
	if err != nil {
		return nil, err
	}
	return offset.NewPagination(o, l), nil
}
