package user

import (
	"fmt"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/phone"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
)

type Adapter interface {
	Adapt(u *models.SystemAccount) (*user.User, error)
	AdaptGender(g string) (account.Gender, error)
	AdaptProfession(p string) (account.Profession, error)
}

type adapter struct {
}

func NewAdapter() Adapter {
	return &adapter{}
}

func (a *adapter) Adapt(u *models.SystemAccount) (*user.User, error) {
	aID, err := account.NewID(u.SystemAccountID.String())
	if err != nil {
		return nil, err
	}
	email, err := account.NewEmail(u.Profile.Email)
	if err != nil {
		return nil, err
	}
	name, err := account.NewName(u.Profile.Name)
	if err != nil {
		return nil, err
	}
	var pn *phone.Number = nil
	if u.PhoneNumber != nil {
		tmpPn, err := phone.NewPhoneNumber(u.PhoneNumber.PhoneNumber)
		if err != nil {
			return nil, err
		}
		pn = &tmpPn
	}
	return user.NewUser(aID, email, name, pn), nil
}

func (a *adapter) AdaptGender(g string) (account.Gender, error) {
	switch g {
	case "man":
		return account.GenderMan, nil
	case "wom":
		return account.GenderWoman, nil
	case "unk":
		return account.GenderUnknown, nil
	}
	return account.Gender(0), fmt.Errorf("invalid gender code: %s", g)
}

func (a *adapter) AdaptProfession(p string) (account.Profession, error) {
	switch p {
	case "coo":
		return account.ProfessionCompanyAndOrganizationOfficers, nil
	case "coe":
		return account.ProfessionCompanyAndOrganizationEmployees, nil
	case "nat":
		return account.ProfessionNationalPublicOfficer, nil
	case "loc":
		return account.ProfessionLocalPublicEmployees, nil
	case "tem":
		return account.ProfessionTemporaryStaffContractEmployees, nil
	case "par":
		return account.ProfessionPartTimeJob, nil
	case "hou":
		return account.ProfessionHousewifeHousehusband, nil
	case "stu":
		return account.ProfessionStudent, nil
	case "pre":
		return account.ProfessionPreschooler, nil
	case "for":
		return account.ProfessionForeignStudent, nil
	case "pen":
		return account.ProfessionPensionRecipients, nil
	case "ret":
		return account.ProfessionRetiredOrUnemployed, nil
	case "unk":
		return account.ProfessionUnknown, nil
	}
	return account.Profession(0), fmt.Errorf("invalid profession code: %s", p)
}
