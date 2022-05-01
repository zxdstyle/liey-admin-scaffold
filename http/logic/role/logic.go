package role

import (
	"context"
	"github.com/zxdstyle/liey-admin-scaffold/http/model"
	"github.com/zxdstyle/liey-admin-scaffold/http/repository"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type Logic struct {
	*bases.BaseLogic
}

func NewLogic() *Logic {
	return &Logic{
		BaseLogic: bases.NewBaseLogic(repository.Role),
	}
}

func (Logic) Create(ctx context.Context, mo bases.RepositoryModel) error {
	if err := repository.Role.Create(ctx, mo); err != nil {
		return err
	}

	role := mo.(*model.Role)
	if role.Permissions != nil {
		return repository.Role.AttachPermissions(ctx, role, role.Permissions)
	}

	return nil
}

func (Logic) Update(ctx context.Context, mo bases.RepositoryModel) error {
	if err := repository.Role.Update(ctx, mo); err != nil {
		return err
	}

	role := mo.(*model.Role)
	if role.Permissions != nil {
		return repository.Role.SyncPermissions(ctx, role, role.Permissions)
	}

	return nil
}
