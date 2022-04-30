package scaffold

import (
	"github.com/zxdstyle/liey-admin-scaffold/http/handler"
	"github.com/zxdstyle/liey-admin-scaffold/migrations"
	"github.com/zxdstyle/liey-admin/framework/database"
	"github.com/zxdstyle/liey-admin/framework/http/server"
)

func RegisterRoutes(group *server.RouterGroup) {
	group.Resource("roles", &handler.Role{})
}

type Plugin struct {
}

func (p Plugin) Name() string {
	return "scaffold"
}

func (p Plugin) Boot() error {
	return nil
}

func (p Plugin) Migrations() []database.Migration {
	return []database.Migration{
		migrations.RoleMigration{},
	}
}
