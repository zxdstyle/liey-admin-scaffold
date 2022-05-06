package scaffold

import (
	"github.com/zxdstyle/liey-admin-scaffold/http/handler"
	"github.com/zxdstyle/liey-admin-scaffold/migrations"
	"github.com/zxdstyle/liey-admin/framework/adm"
	migrator "github.com/zxdstyle/liey-admin/framework/database/migrations"
	"github.com/zxdstyle/liey-admin/framework/http/server"
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
	return nil
}

func (p Plugin) Migrations() []migrator.Migration {
	return []migrator.Migration{
		migrations.ScaffoldMigration{},
	}
}
