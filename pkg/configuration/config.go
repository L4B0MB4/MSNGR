package configuration

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type ConfigProvider struct {
	config map[string]string
}

func NewConfigurationProvider() *ConfigProvider {

	err := godotenv.Load()
	if err != nil {
		log.Warn().Msg("Error loading .env file. Env variables will still be read")
	}

	config := make(map[string]string)

	config["PORT"] = "6616"
	config["HOST"] = "0.0.0.0"
	config["DISCORD_BOT_TOKEN"] = os.Getenv("DISCORD_BOT_TOKEN")
	config["DISCORD_CHANNEL_ID"] = os.Getenv("DISCORD_CHANNEL_ID")

	return &ConfigProvider{
		config: config,
	}
}

func (c *ConfigProvider) GetStringValue(key string) string {

	v, ok := c.config[key]
	if !ok {
		log.Warn().Str("Key", key).Msg("Requesting value for key that is not present")
		return ""
	}
	return v

}
