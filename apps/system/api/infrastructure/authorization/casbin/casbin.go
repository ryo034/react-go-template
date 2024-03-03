package casbin

import (
	_ "embed"

	"github.com/casbin/casbin/v2/model"

	"github.com/casbin/casbin/v2"
	jsonadapter "github.com/casbin/json-adapter/v2"
)

//go:embed casbin_model.conf
var Model string

//go:embed casbin_policy.json
var Policy []byte

func NewEnforcer() (*casbin.Enforcer, error) {
	m, err := model.NewModelFromString(Model)
	if err != nil {
		return nil, err
	}
	a := jsonadapter.NewAdapter(&Policy)
	return casbin.NewEnforcer(m, a)
}
