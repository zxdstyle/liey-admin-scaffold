package scaffold

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/zxdstyle/liey-admin-scaffold/gate"
	"github.com/zxdstyle/liey-admin-scaffold/http/handler"
	"github.com/zxdstyle/liey-admin-scaffold/migrations"
	"github.com/zxdstyle/liey-admin-scaffold/publishable"
	"github.com/zxdstyle/liey-admin/framework/adm"
	migrator "github.com/zxdstyle/liey-admin/framework/database/migrations"
	"github.com/zxdstyle/liey-admin/framework/http/server"
	"github.com/zxdstyle/liey-admin/framework/support/publish/publisher"
)

func RegisterRoutes(group *server.RouterGroup) {
	group.Resource("roles", &handler.Role{})
	group.Resource("permissions", &handler.Permission{})

	if adm.Debug() {
		group.GET("routes", handler.Routes)
	}
}

type Plugin struct {
}

func (p Plugin) Name() string {
	return "scaffold"
}

func (p Plugin) Boot() error {
	ctx := context.Background()
	ga, err := gate.NewCasbinGate(adm.DB())
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	if err := gate.RegisterGate(ga); err != nil {
		g.Log().Fatal(ctx, err)
	}

	return nil
}

func (p Plugin) Migrations() []migrator.Migration {
	return []migrator.Migration{
		migrations.ScaffoldMigration{},
	}
}

func (p Plugin) Resources() []publisher.Publisher {
	return []publisher.Publisher{
		publishable.ConfigPublisher,
	}
}
