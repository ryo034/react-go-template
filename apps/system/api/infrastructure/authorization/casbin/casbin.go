package casbin

import (
	_ "embed"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	stringadapter "github.com/casbin/casbin/v2/persist/string-adapter"
)

//go:embed casbin_model.conf
var Model string

//go:embed casbin_policy.csv
var Policy string

func NewEnforcer() (*casbin.Enforcer, error) {
	m, err := model.NewModelFromString(Model)
	if err != nil {
		return nil, err
	}
	a := stringadapter.NewAdapter(Policy)
	return casbin.NewEnforcer(m, a)
}
