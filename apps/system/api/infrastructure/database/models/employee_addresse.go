package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type EmployeeAddress struct {
	bun.BaseModel `bun:"table:employee_addresses,alias:ea"`

	EmployeeID               uuid.UUID `bun:"employee_id,pk"`
	PostalCode               string    `bun:"postal_code"`
	BuildingComponentID      uuid.UUID `bun:"building_component_id"`
	StreetAddressComponentID uuid.UUID `bun:"street_address_component_id,notnull"`
	CityComponentID          uuid.UUID `bun:"city_component_id,notnull"`
	StateComponentID         uuid.UUID `bun:"state_component_id,notnull"`
	CountryComponentID       uuid.UUID `bun:"country_component_id,notnull"`
	CreatedAt                time.Time `bun:"created_at,notnull,default:current_timestamp"`

	Employee *Employee         `bun:"rel:belongs-to"`
	Country  *AddressComponent `bun:"rel:belongs-to,join:country_component_id=component_id"`
	State    *AddressComponent `bun:"rel:belongs-to,join:state_component_id=component_id"`
	City     *AddressComponent `bun:"rel:belongs-to,join:city_component_id=component_id"`
	Street   *AddressComponent `bun:"rel:belongs-to,join:street_address_component_id=component_id"`
	Building *AddressComponent `bun:"rel:belongs-to,join:building_component_id=component_id"`
}
