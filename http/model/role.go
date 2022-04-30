package model

import "github.com/zxdstyle/liey-admin/framework/http/bases"

type (
	Role struct {
		bases.Model
		Name    *string `json:"name"`
		Status  *int    `json:"status"`
		SortNum *int    `json:"sort_num"`
	}

	Roles []*Role
)

func (r Roles) GetModel(i int) bases.RepositoryModel {
	if len(r) == 0 {
		return Role{}
	}
	return r[i]
}
