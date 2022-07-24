package provider

import (
	"errors"
	"os"

	"github.com/zhangdapeng520/zdpgo_nacos/sdk/auth/credentials"

	"github.com/zhangdapeng520/zdpgo_nacos/sdk/auth"
)

type EnvProvider struct{}

var ProviderEnv = new(EnvProvider)

func NewEnvProvider() Provider {
	return &EnvProvider{}
}

func (p *EnvProvider) Resolve() (auth.Credential, error) {
	accessKeyID, ok1 := os.LookupEnv(ENVAccessKeyID)
	accessKeySecret, ok2 := os.LookupEnv(ENVAccessKeySecret)
	if !ok1 || !ok2 {
		return nil, nil
	}
	if accessKeyID == "" || accessKeySecret == "" {
		return nil, errors.New("Environmental variable (ALIBABACLOUD_ACCESS_KEY_ID or ALIBABACLOUD_ACCESS_KEY_SECRET) is empty")
	}
	return credentials.NewAccessKeyCredential(accessKeyID, accessKeySecret), nil
}
