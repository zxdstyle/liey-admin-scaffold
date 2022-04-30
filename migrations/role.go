package migrations

import (
	"github.com/zxdstyle/liey-admin-scaffold/http/model"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type RoleMigration struct {
}

func (r RoleMigration) Models() []bases.RepositoryModel {
	return []bases.RepositoryModel{
		model.Role{},
	}
}
