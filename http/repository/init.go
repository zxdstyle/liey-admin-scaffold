package repository

import (
	"github.com/zxdstyle/liey-admin-scaffold/http/repository/permission"
	"github.com/zxdstyle/liey-admin-scaffold/http/repository/role"
)

var (
	Role       = role.NewDbRepository()
	Permission = permission.NewDbRepository()
)
