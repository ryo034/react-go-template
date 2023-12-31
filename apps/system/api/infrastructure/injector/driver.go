package injector

import (
	firebaseDriver "github.com/ryo034/react-go-template/apps/system/api/driver/firebase"
	beDriver "github.com/ryo034/react-go-template/apps/system/api/driver/sqlboiler/business_entity"
	"github.com/ryo034/react-go-template/apps/system/api/driver/sqlboiler/business_entity/employee"
	employeeDriver "github.com/ryo034/react-go-template/apps/system/api/driver/sqlboiler/business_entity/employee"
	meDriver "github.com/ryo034/react-go-template/apps/system/api/driver/sqlboiler/me"
	storeDriver "github.com/ryo034/react-go-template/apps/system/api/driver/sqlboiler/store"
	itemDriver "github.com/ryo034/react-go-template/apps/system/api/driver/sqlboiler/store/item"
	staffDriver "github.com/ryo034/react-go-template/apps/system/api/driver/sqlboiler/store/staff"
	transactionDriver "github.com/ryo034/react-go-template/apps/system/api/driver/sqlboiler/store/transaction"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
)

type driverInjector struct {
	Firebase       firebaseDriver.Driver
	Me             meDriver.Driver
	Store          storeDriver.Driver
	Item           itemDriver.Driver
	BusinessEntity beDriver.Driver
	Employee       employeeDriver.Driver
	Staff          staffDriver.Driver
	Transaction    transactionDriver.Driver
}

func newDriverInjector(f *firebase.Firebase) driverInjector {
	ed := employee.NewDriver()
	return driverInjector{
		Firebase:       firebaseDriver.NewDriver(f),
		Me:             meDriver.NewDriver(),
		Store:          storeDriver.NewDriver(ed),
		Item:           itemDriver.NewDriver(),
		BusinessEntity: beDriver.NewDriver(ed),
		Employee:       ed,
		Staff:          staffDriver.NewDriver(),
		Transaction:    transactionDriver.NewDriver(),
	}
}
