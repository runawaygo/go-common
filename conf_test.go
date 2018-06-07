package go_common

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
)

type RemoteConfig struct {
	Eureka struct {
		Address string `yaml:address`
	}
	Server struct {
		Port int `yaml:"port"`
	}
}

func TestGetConfFromConfigserver(t *testing.T) {
	os.Setenv("GO-ENV", "k8s")
	tests := []struct {
		name string
	}{
		{name: "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := RemoteConfig{}
			GetAppConfig("simulator-go", &c)
			assert.Equal(t, 8031, c.Server.Port, "server port should be equal")
			assert.Equal(t, "http://eureka-server-primary-svc/eureka/,http://eureka-server-secondary-svc/eureka/,http://eureka-server-tertiary-svc/eureka/",
				c.Eureka.Address, "eureka address should be equal")

		})
	}
}
