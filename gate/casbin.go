package gate

import (
	"github.com/casbin/casbin/v2"
	adapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var (
	rbacModel = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
`
)

type (
	CasbinGate struct {
		enforcer *casbin.Enforcer
	}

	Auth struct {
		AuthID       uint
		PermissionId uint
	}
)

func NewCasbinGate(db *gorm.DB) *CasbinGate {
	a, _ := adapter.NewAdapterByDB(db)
	e, _ := casbin.NewEnforcer(rbacModel, a)
	g := &CasbinGate{
		enforcer: e,
	}
	return g
}

func (g CasbinGate) Can(auth Auth) error {
	return nil
}
