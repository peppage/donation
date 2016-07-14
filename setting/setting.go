package setting

import (
	"github.com/pelletier/go-toml"
)

var (
	SecretKey      string
	PublishableKey string
	config         *toml.TomlTree
)

func init() {
	var err error

	config, err = toml.LoadFile("conf.toml")
	if err != nil {
		panic("Failed to open conf.toml")
	}
}

func Initialize() {
	if config.Has("stripe.secret_key") {
		SecretKey = config.Get("stripe.secret_key").(string)
	}
	if config.Has("stripe.publishable_key") {
		PublishableKey = config.Get("stripe.publishable_key").(string)
	}
}
