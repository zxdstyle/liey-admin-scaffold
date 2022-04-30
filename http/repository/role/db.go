package role

import (
	"github.com/zxdstyle/liey-admin-scaffold/http/model"
	"github.com/zxdstyle/liey-admin/framework/adm"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type DbRepository struct {
	*bases.GormRepository
}

func NewDbRepository() Repository  {
	return &DbRepository{
		GormRepository: bases.NewGormRepository(adm.DB().Model(model.Role{})),
	}
}
