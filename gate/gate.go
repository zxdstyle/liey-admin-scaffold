package gate

import (
	"fmt"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/zxdstyle/liey-admin/framework/exception"
)

type (
	Gate interface {
		Name() string
		Can(authID uint, path, action string) error
	}
)

var (
	gates = gmap.NewStrAnyMap(true)

	ErrForbidden = gerror.NewCode(exception.CodeForbidden, "无权限访问")
)

func RegisterGate(gate Gate) error {
	if gates.Contains(gate.Name()) {
		return gerror.NewCode(exception.CodeInternalError, fmt.Sprintf("拦截器 %s 已存在", gate.Name()))
	}
	gates.Set(gate.Name(), gate)
	return nil
}

func GetGate(name string) Gate {
	val, ok := gates.Search(name)
	if !ok {
		return nil
	}
	return val.(Gate)
}

func GetDefaultGate() Gate {
	val, ok := gates.Search(defaultCasbinGateName)
	if !ok {
		return nil
	}
	return val.(Gate)
}
