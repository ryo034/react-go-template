package address

import (
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/address"
	models "github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/sqlboiler/api"
	dbErr "github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/sqlboiler/error"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type (
	Driver interface {
		Save(ctx context.Context, exec boil.ContextExecutor, a *address.Address) error
	}
)

type driver struct{}

func NewDriver() Driver {
	return &driver{}
}

func (d *driver) Find(ctx context.Context, exec boil.ContextExecutor, aID address.ID) (*models.Address, error) {
	return models.Addresses(
		models.AddressWhere.AddressID.EQ(aID.ToString()),
	).One(ctx, exec)
}

func (d *driver) Save(ctx context.Context, exec boil.ContextExecutor, a *address.Address) error {
	resCi, err := models.Cities(models.CityWhere.Name.EQ(a.City().ToString())).One(ctx, exec)
	if err != nil {
		if !dbErr.IsNoSuchDataError(err) {
			return err
		}
	}

	if resCi == nil {
		resCi.CityID = a.City().ID().ToString()
		resCi.PrefectureCode = a.Prefecture().ID().ToString()
		resCi.Name = a.City().ToString()
		if err := resCi.Insert(ctx, exec, boil.Infer()); err != nil {
			return err
		}
	}

	resZi, err := models.ZipCodes(models.ZipCodeWhere.ZipCode.EQ(a.ZipCode().ToString())).One(ctx, exec)
	if err != nil {
		if !dbErr.IsNoSuchDataError(err) {
			return err
		}
	}
	if resZi == nil {
		resZi.ZipCode = a.ZipCode().ToString()
		resZi.CityID = a.City().ID().ToString()
		resZi.ZipCode = a.ZipCode().ToString()
		if err := resZi.Insert(ctx, exec, boil.Infer()); err != nil {
			return err
		}
	}

	resSt, err := models.Streets(models.StreetWhere.Name.EQ(a.Street().ToString())).One(ctx, exec)
	if err != nil {
		if !dbErr.IsNoSuchDataError(err) {
			return err
		}
	}
	if resSt == nil {
		resSt.StreetID = a.Street().ID().ToString()
		resSt.Name = a.Street().ToString()
		if err := resSt.Insert(ctx, exec, boil.Infer()); err != nil {
			return err
		}
	}

	resBu, err := models.Buildings(models.BuildingWhere.Name.EQ(a.Building().ToString())).One(ctx, exec)
	if err != nil {
		if !dbErr.IsNoSuchDataError(err) {
			return err
		}
	}
	if resBu == nil {
		resBu.BuildingID = a.Building().ID().ToString()
		resBu.R.City.CityID = a.City().ID().ToString()
		resBu.Name = a.Building().ToString()
		if err := resBu.Insert(ctx, exec, boil.Infer()); err != nil {
			return err
		}
	}

	ad := &models.Address{
		AddressID:  a.ID().ToString(),
		BuildingID: resBu.BuildingID,
	}
	if err := ad.Insert(ctx, exec, boil.Infer()); err != nil {
		return err
	}

	return nil
}
