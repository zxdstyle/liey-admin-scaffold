package role

import (
	"context"
	"github.com/zxdstyle/liey-admin-scaffold/http/model"
	"github.com/zxdstyle/liey-admin/framework/adm"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
)

type dbRepository struct {
	*bases.GormRepository
}

func NewDbRepository() Repository {
	return &dbRepository{
		GormRepository: bases.NewGormRepository(adm.DB().Model(model.Role{})),
	}
}

// AttachPermissions 添加权限
func (repo *dbRepository) AttachPermissions(ctx context.Context, role *model.Role, permissions *model.Permissions) error {
	if role == nil || permissions == nil {
		return nil
	}

	if err := repo.Orm.WithContext(ctx).Model(role).Omit("Permissions.*").Association("Permissions").Append(permissions); err != nil {
		return err
	}

	if err := repo.Show(ctx, requests.Resources{"Permissions"}, role); err != nil {
		return err
	}

	return nil
}

// SyncPermissions 同步权限
func (repo *dbRepository) SyncPermissions(ctx context.Context, role *model.Role, permissions *model.Permissions) error {
	if role == nil || permissions == nil {
		return nil
	}

	tx := repo.Orm.WithContext(ctx).Begin()

	if err := tx.Model(role).Association("Permissions").Clear(); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(role).Omit("Permissions.*").Association("Permissions").Append(permissions); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	if err := repo.Show(ctx, requests.Resources{"Permissions"}, role); err != nil {
		return err
	}

	return nil
}
