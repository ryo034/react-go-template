package shared

import (
	"github.com/ryo034/react-go-template/packages/go/domain/shared/address"
	models "github.com/ryo034/react-go-template/packages/go/infrastructure/database/sqlboiler/api"
	"github.com/spf13/cast"
)

type Adapter interface {
	AdaptAddress(ad *models.Address) (*address.Address, error)
}

type adapter struct {
}

func NewAdapter() Adapter {
	return &adapter{}
}

func (a *adapter) AdaptAddress(ad *models.Address) (*address.Address, error) {
	bID, err := address.NewID(ad.R.Building.BuildingID)
	if err != nil {
		return nil, err
	}
	bu, err := address.NewBuilding(bID, ad.R.Building.Name)
	if err != nil {
		return nil, err
	}
	cID, err := address.NewID(ad.R.Building.R.City.CityID)
	if err != nil {
		return nil, err
	}
	ci, err := address.NewCity(cID, ad.R.Building.R.City.Name)
	if err != nil {
		return nil, err
	}
	var st *address.Street = nil
	if ad.R.Streets != nil {
		stID, err := address.NewID(ad.R.Streets[0].StreetID)
		if err != nil {
			return nil, err
		}
		tmpSt, err := address.NewStreet(stID, ad.R.Streets[0].Name)
		if err != nil {
			return nil, err
		}
		st = &tmpSt
	}
	preID, err := address.NewID(cast.ToString(ad.R.Building.R.City.R.PrefectureCodePrefecture.PrefectureCode))
	if err != nil {
		return nil, err
	}
	pre, err := address.NewPrefecture(preID, ad.R.Building.R.City.R.PrefectureCodePrefecture.Name)
	if err != nil {
		return nil, err
	}

	coID, err := address.NewID(cast.ToString(ad.R.Building.R.City.R.PrefectureCodePrefecture.R.CountryCodeCountry.CountryCode))
	if err != nil {
		return nil, err
	}
	co, err := address.NewCountry(coID, ad.R.Building.R.City.R.PrefectureCodePrefecture.R.CountryCodeCountry.Name)
	if err != nil {
		return nil, err
	}
	zc, err := address.NewZipCode(ad.R.Building.R.City.R.ZipCodes[0].ZipCode)
	if err != nil {
		return nil, err
	}
	aID, err := address.NewID(ad.AddressID)
	if err != nil {
		return nil, err
	}
	return address.NewAddress(aID, zc, co, pre, ci, st, bu), nil
}
