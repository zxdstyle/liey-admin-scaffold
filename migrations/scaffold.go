package migrations

import (
	"github.com/zxdstyle/liey-admin-scaffold/http/model"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type ScaffoldMigration struct {
}

func (r ScaffoldMigration) Models() []bases.RepositoryModel {
	return []bases.RepositoryModel{
		&model.Role{},
		&model.Permission{},
		&model.RoleHasPermission{},
	}
}
