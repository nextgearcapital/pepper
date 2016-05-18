package device42

import (
	"os"

	"github.com/nextgearcapital/pepper/pkg/log"

	"github.com/spf13/viper"
)

// Config :
var Config = viper.New()

var (
	// Username :
	Username = Config.GetString("username")
	// Password :
	Password = Config.GetString("password")
	address  = Config.GetString("address")
	// PrdRange :
	PrdRange = Config.GetString("ip_ranges.production")
	// DevRange :
	DevRange = Config.GetString("ip_ranges.development")
	// TestRange :
	TestRange = Config.GetString("ip_ranges.testing")
	// StageRange :
	StageRange = Config.GetString("ip_ranges.staging")
	// UATRange :
	UATRange = Config.GetString("ip_ranges.uat")
	// BaseURL :
	BaseURL = address + "/api/1.0/"
)

// ReadConfig :
func ReadConfig() error {
	config := "/etc/pepper/provider.d/device42"

	if err := os.MkdirAll(config, 0755); err != nil {
		log.Err("%s", err)
	}

	Config.SetConfigName("device42")
	Config.SetConfigType("yaml")
	Config.AddConfigPath("/etc/pepper/provider.d/device42")
	if err := Config.ReadInConfig(); err != nil {
		log.Die("%s", err)
	}
	return nil
}
