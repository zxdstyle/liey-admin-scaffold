package repository

import (
	"github.com/zxdstyle/liey-admin-scaffold/http/repository/admin"
	"github.com/zxdstyle/liey-admin-scaffold/http/repository/permission"
	"github.com/zxdstyle/liey-admin-scaffold/http/repository/role"
)

var (
	Admin      = admin.NewDbRepository()
	Role       = role.NewDbRepository()
	Permission = permission.NewDbRepository()
)
