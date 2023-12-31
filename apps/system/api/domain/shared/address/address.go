package address

import "fmt"

type Address struct {
	id         ID
	zipCode    ZipCode
	country    Country
	prefecture Prefecture
	city       City
	street     *Street
	building   Building
}

func NewAddress(
	id ID,
	zipCode ZipCode,
	country Country,
	prefecture Prefecture,
	city City,
	street *Street,
	building Building,
) *Address {
	return &Address{
		id:         id,
		zipCode:    zipCode,
		country:    country,
		prefecture: prefecture,
		city:       city,
		street:     street,
		building:   building,
	}
}

func (a *Address) ToString() string {
	return fmt.Sprintf("%s %s %s %s %s %s", a.zipCode.ToString(), a.country.ToString(), a.prefecture.ToString(), a.city.ToString(), a.street.ToString(), a.building.ToString())
}

func (a *Address) ID() ID {
	return a.id
}

func (a *Address) ZipCode() ZipCode {
	return a.zipCode
}

func (a *Address) Country() Country {
	return a.country
}

func (a *Address) Prefecture() Prefecture {
	return a.prefecture
}

func (a *Address) City() City {
	return a.city
}

func (a *Address) Street() *Street {
	return a.street
}

func (a *Address) HasStreet() bool {
	return a.street != nil
}

func (a *Address) HasNotStreet() bool {
	return !a.HasStreet()
}

func (a *Address) Building() Building {
	return a.building
}
