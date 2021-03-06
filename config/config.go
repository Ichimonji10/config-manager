package config

import (
	"fmt"
	"os"

	clowder "github.com/redhatinsights/app-common-go/pkg/api/v1"
	"github.com/spf13/viper"
)

func Get() *viper.Viper {
	options := viper.New()

	if os.Getenv("CLOWDER_ENABLED") == "true" {
		cfg := clowder.LoadedConfig

		options.SetDefault("WebPort", cfg.WebPort)
		options.SetDefault("MetricsPort", cfg.MetricsPort)
		options.SetDefault("MetricsPath", cfg.MetricsPath)

		options.SetDefault("KafkaBrokers", fmt.Sprintf(cfg.Kafka.Brokers[0].Hostname, cfg.Kafka.Brokers[0].Port))

		options.SetDefault("LogGroup", cfg.Logging.Cloudwatch.LogGroup)
		options.SetDefault("AwsRegion", cfg.Logging.Cloudwatch.Region)
		options.SetDefault("AwsAccessKeyId", cfg.Logging.Cloudwatch.AccessKeyId)
		options.SetDefault("AwsSecretAccessKey", cfg.Logging.Cloudwatch.SecretAccessKey)

		options.SetDefault("DBHost", cfg.Database.Hostname)
		options.SetDefault("DBPort", cfg.Database.Port)
		options.SetDefault("DBName", cfg.Database.Name)
		options.SetDefault("DBUser", cfg.Database.Username)
		options.SetDefault("DBPass", cfg.Database.Password)
	} else {
		options.SetDefault("WebPort", 8081)
		options.SetDefault("MetricsPort", 9000)

		options.SetDefault("KafkaBrokers", []string{"localhost:29092"})

		options.SetDefault("LogGroup", "platform-dev")
		options.SetDefault("AwsRegion", "us-east-1")
		options.SetDefault("AwsAccessKeyId", os.Getenv("CW_AWS_ACCESS_KEY_ID"))
		options.SetDefault("AwsSecretAccessKey", os.Getenv("CW_AWS_SECRET_ACCESS_KEY"))

		options.SetDefault("DBHost", "localhost")
		options.SetDefault("DBPort", 5432)
		options.SetDefault("DBName", "insights")
		options.SetDefault("DBUser", "insights")
		options.SetDefault("DBPass", "insights")
	}

	options.SetDefault("KafkaGroupID", "config-manager")
	options.SetDefault("KafkaConsumerOffset", -1)
	options.SetDefault("KafkaResultsTopic", "platform.playbook-dispatcher.results")
	options.SetDefault("KafkaConnectionsTopic", "platform.inventory.connections")

	options.SetDefault("PlaybookPath", "./playbooks/")

	options.SetDefault("ServiceConfig", `{
		"insights": "enabled",
		"compliance_openscap": "enabled",
		"resource_optimization": "enabled",
		"remediations": "enabled"
	}`)

	options.SetEnvPrefix("CM")
	options.AutomaticEnv()

	return options
}
