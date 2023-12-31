package address

import (
	sharedPb "github.com/ryo034/react-go-template/apps/system/api/schema/pb/shared/v1"
	"github.com/ryo034/react-go-template/packages/go/domain/shared/address"
)

type Adapter interface {
	Adapt(ad *address.Address) *sharedPb.Address
}

type adapter struct {
}

func NewAdapter() Adapter {
	return &adapter{}
}

func (a *adapter) Adapt(ad *address.Address) *sharedPb.Address {
	return &sharedPb.Address{
		ZipCode:    ad.ZipCode().ToString(),
		Country:    ad.Country().ToString(),
		Prefecture: ad.Prefecture().ToString(),
		City:       ad.City().ToString(),
		Street:     ad.Street().ToString(),
		Building:   ad.Building().ToString(),
	}
}
