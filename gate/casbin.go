package gate

import (
	"context"
	"github.com/casbin/casbin/v2"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	defaultCasbinGateName  = "casbinGate"
	defaultCasbinModelPath = "config/rbac_model.conf"
)

type (
	CasbinGate struct {
		name     string
		enforcer *casbin.SyncedEnforcer
	}

	Auth struct {
		AuthID       uint
		PermissionId uint
	}
)

func NewCasbinGate(db *gorm.DB) (Gate, error) {
	enforcer, err := initAdapter(db)
	if err != nil {
		return nil, err
	}
	g := &CasbinGate{
		name:     defaultCasbinGateName,
		enforcer: enforcer,
	}
	return g, nil
}

func NewCasbinGateWithName(db *gorm.DB, name string) (Gate, error) {
	enforcer, err := initAdapter(db)
	if err != nil {
		return nil, err
	}
	g := &CasbinGate{
		name:     name,
		enforcer: enforcer,
	}
	return g, nil
}

func initAdapter(db *gorm.DB) (*casbin.SyncedEnforcer, error) {
	db = db.Session(&gorm.Session{
		Logger: db.Logger.LogMode(logger.Error),
	})
	a, err := adapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}
	val, _ := g.Cfg("scaffold").Get(context.Background(), "model_path", defaultCasbinModelPath)
	e, er := casbin.NewSyncedEnforcer(val.String(), a)
	if er != nil {
		return nil, er
	}
	if err := e.LoadPolicy(); err != nil {
		return nil, err
	}
	return e, nil
}

func (g CasbinGate) Name() string {
	return g.name
}

func (g CasbinGate) Can(authID uint, path, action string) error {
	ok, err := g.enforcer.Enforce(authID, path, action)
	if err != nil {
		return err
	}
	if !ok {
		return ErrForbidden
	}
	return nil
}
