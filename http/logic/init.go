package logic

import (
	"github.com/zxdstyle/liey-admin-scaffold/http/logic/auth"
	"github.com/zxdstyle/liey-admin-scaffold/http/logic/permission"
	"github.com/zxdstyle/liey-admin-scaffold/http/logic/role"
)

var (
	Auth       = auth.NewLogic()
	Role       = role.NewLogic()
	Permission = permission.NewLogic()
)
