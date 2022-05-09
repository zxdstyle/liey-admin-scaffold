package publishable

import (
	_ "embed"
	"github.com/zxdstyle/liey-admin/framework/support/publish/publisher"
)

var (
	//go:embed scaffold.yaml
	config []byte

	ConfigPublisher = publisher.NewConfigPublisher(config)
)
