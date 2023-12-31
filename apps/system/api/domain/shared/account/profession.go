package account

type Profession int

const (
	ProfessionUnknown Profession = iota + 1
	ProfessionCompanyAndOrganizationOfficers
	ProfessionCompanyAndOrganizationEmployees
	ProfessionNationalPublicOfficer
	ProfessionLocalPublicEmployees
	ProfessionTemporaryStaffContractEmployees
	ProfessionPartTimeJob
	ProfessionHousewifeHousehusband
	ProfessionStudent
	ProfessionPreschooler
	ProfessionForeignStudent
	ProfessionPensionRecipients
	ProfessionRetiredOrUnemployed
)

func (p Profession) EnumIndex() int {
	return int(p)
}
