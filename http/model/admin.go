package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type (
	Admin struct {
		bases.Model
		Email      *string     `json:"email,omitempty" gorm:"not null;type:varchar(64);comment:邮箱"`
		Password   *string     `json:"password,omitempty" gorm:"not null;comment:密码"`
		Avatar     *string     `json:"avatar,omitempty" gorm:"not null;default:'';comment:头像"`
		IsActive   *bool       `json:"isActive,omitempty" gorm:"not null;type:tinyint;default:1;comment:状态 0：禁用 1：启用"`
		RegisterAt *gtime.Time `json:"registerAt,omitempty" gorm:"comment:用户注册时间"`
	}

	Admins []*Admin
)

func (r Admins) GetModel(i int) bases.RepositoryModel {
	if len(r) == 0 {
		return &Admin{}
	}
	return r[i]
}

func (Admin) GuardName() string {
	return "api"
}
