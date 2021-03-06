package instance

import (
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type Instance interface {
	Apply(wait bool) error
	Destroy(force bool) error
	Init(dockercfgPath string, inputOverrides map[string]interface{}) error
	Output(key string) (string, error)
	Plan() error
	Pull(s s3iface.S3API) error
	Push(s s3iface.S3API) error
	Set(inputs map[string]interface{}) error
	Upgrade(version string, force bool) error
}
