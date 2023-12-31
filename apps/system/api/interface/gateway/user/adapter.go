package user

import (
	"fmt"
	"github.com/ryo034/react-go-template/apps/system/api/domain/user"
	"github.com/ryo034/react-go-template/packages/go/domain/shared/account"
	"github.com/ryo034/react-go-template/packages/go/domain/shared/phone"
	models "github.com/ryo034/react-go-template/packages/go/infrastructure/database/sqlboiler/api"
)

type Adapter interface {
	Adapt(s *models.User) (*user.User, error)
	AdaptGender(g string) (account.Gender, error)
	AdaptProfession(p string) (account.Profession, error)
}

type adapter struct {
}

func NewAdapter() Adapter {
	return &adapter{}
}

func (a *adapter) Adapt(u *models.User) (*user.User, error) {
	aID, err := account.NewID(u.UserID)
	if err != nil {
		return nil, err
	}
	e, err := account.NewEmail(u.Email)
	if err != nil {
		return nil, err
	}
	var phoneNumber *phone.Number = nil
	if u.R != nil {
		if u.R.UserPhoneNumber != nil {
			tmpPh, err := phone.NewPhoneNumber(u.R.UserPhoneNumber.PhoneNumber)
			if err != nil {
				return nil, err
			}
			phoneNumber = &tmpPh
		}
	}
	fn, err := account.NewFirstName(u.FirstName)
	if err != nil {
		return nil, err
	}
	ln, err := account.NewLastName(u.LastName)
	if err != nil {
		return nil, err
	}
	return user.NewUser(aID, e, phoneNumber, fn, ln), nil
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
