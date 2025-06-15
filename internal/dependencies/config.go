package dependencies

import (
	"github.com/joho/godotenv"
	"hermes/internal/pkg/logger"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name    string `mapstructure:"name"`
		Address string `mapstructure:"address"`
		Key     string `mapstructure:"key"`
		Port    int    `mapstructure:"port"`
		Debug   bool   `mapstructure:"debug"`
	} `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
}

type DatabaseConfig struct {
	Host        string `mapstructure:"host"`
	Database    string `mapstructure:"database"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	Dialect     string `mapstructure:"dialect"`
	AutoMigrate bool   `mapstructure:"autoMigrate"`
	Pool        struct {
		Max     int `mapstructure:"max"`
		Min     int `mapstructure:"min"`
		Acquire int `mapstructure:"acquire"`
		Idle    int `mapstructure:"idle"`
	} `mapstructure:"pool"`
	Port int `mapstructure:"port"`
}

func LoadConfig(path string) (*Config, error) {
	// Load .env file into environment variables
	if err := godotenv.Load(); err != nil {
		logger.Warn("No .env file found, using environment variables only: ", err)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")

	// Set environment variable prefix for nested configs
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	config := &Config{}
	bindEnvs(config, "")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if err = viper.Unmarshal(config); err != nil {
		panic("failed to unmarshal config: " + err.Error())
	}

	return config, nil
}

func bindEnvs(iface interface{}, parentKey string) {
	t := reflect.TypeOf(iface)
	v := reflect.ValueOf(iface)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := v.Field(i)

		tag := field.Tag.Get("mapstructure")
		if tag == "" {
			continue
		}

		fullKey := tag
		if parentKey != "" {
			fullKey = parentKey + "." + tag
		}

		// Handle nested structs
		if fieldVal.Kind() == reflect.Struct {
			bindEnvs(fieldVal.Addr().Interface(), fullKey)
			continue
		}

		// Bind environment variable
		viper.BindEnv(fullKey)
	}
}
