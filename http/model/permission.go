package model

import "github.com/zxdstyle/liey-admin/framework/http/bases"

type (
	Permission struct {
		bases.Model
		Name  *string          `gorm:"not null;type:varchar(64);comment:名称" json:"name" v:"required-if:id,0"`
		Slug  *string          `gorm:"not null;type:varchar(64);unique;comment:标识" json:"slug" v:"required-if:id,0|unique:permissions,slug"`
		Rules *PermissionRules `gorm:"not null;serializer:json;comment:权限规则" json:"rules" v:"required-if:id,0"`
	}

	Permissions []*Permission

	PermissionRule struct {
	}

	PermissionRules []PermissionRule
)

func (r Permissions) GetModel(i int) bases.RepositoryModel {
	if len(r) == 0 {
		return &Permission{}
	}
	return r[i]
}
