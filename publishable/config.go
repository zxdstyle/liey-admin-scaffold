package publishable

import (
	_ "embed"
	"github.com/zxdstyle/liey-admin/framework/support/publish/publisher"
)

var (
	//go:embed scaffold.yaml
	config []byte
	//go:embed rbac_model.conf
	rbacModel []byte

	ConfigPublisher = publisher.NewConfigPublisher([]publisher.Publishable{
		publisher.NewPublishable("config/scaffold.yaml", config),
		publisher.NewPublishable("config/rbac_model.conf", rbacModel),
	})
)
