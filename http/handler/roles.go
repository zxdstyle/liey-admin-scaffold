package handler

import (
	"context"
	"github.com/zxdstyle/liey-admin-scaffold/http/logic"
	"github.com/zxdstyle/liey-admin-scaffold/http/model"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
)

type Role struct {
	
}

func (r Role) Index(ctx context.Context, req requests.Request) (*responses.Response, error) {
	panic("implement me")
}

func (r Role) Show(ctx context.Context, req requests.Request) (*responses.Response, error) {
	role := &model.Role{}
	if err := logic.Role.Show(ctx, []string{}, role); err != nil {
		return nil, err
	}
	return responses.Success(role), nil
}

func (r Role) Update(ctx context.Context, req requests.Request) (*responses.Response, error) {
	panic("implement me")
}

func (r Role) Create(ctx context.Context, req requests.Request) (*responses.Response, error) {
	panic("implement me")
}

func (r Role) Destroy(ctx context.Context, req requests.Request) (*responses.Response, error) {
	panic("implement me")
}



