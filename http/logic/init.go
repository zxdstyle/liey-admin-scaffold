package logic

import (
	"github.com/zxdstyle/liey-admin-scaffold/http/logic/permission"
	"github.com/zxdstyle/liey-admin-scaffold/http/logic/role"
)

var (
	Role       = role.NewLogic()
	Permission = permission.NewLogic()
)
