package config

import (
	"os"
)

const (
	LOCAL       = "local"
	DEVELOPMENT = "development"
)

const ENVIRONMENT string = DEVELOPMENT

var env = map[string]map[string]string{
	"local": {

		"MYSQL_HOST":   "localhost",
		"MYSQL_PORT":   "3307",
		"MYSQL_USER":   "root",
		"MYSQL_PASS":   "root",
		"MYSQL_SCHEMA": "service_management_mysql",

		"POSTGRES_HOST":   "localhost",
		"POSTGRES_PORT":   "5433",
		"POSTGRES_USER":   "root",
		"POSTGRES_PASS":   "root",
		"POSTGRES_SCHEMA": "service_management_postgres",

		"AUDIT_TRAILS": "true",

		"KAFKA_SERVERS":             "localhost:9092",
		"KAFKA_SCHEMA_REGISTRY_URL": "https://localhost:8181",
		"KAFKA_SASL_USERNAME":       "",
		"KAFKA_SASL_PASSWORD":       "",
	},
	"development": {
		"MYSQL_HOST":   "172.18.136.12",
		"MYSQL_PORT":   "3306",
		"MYSQL_USER":   "bigdata",
		"MYSQL_PASS":   "P@ssw0rd",
		"MYSQL_SCHEMA": "man_db",

		"POSTGRES_HOST":   "localhost",
		"POSTGRES_PORT":   "5433",
		"POSTGRES_USER":   "root",
		"POSTGRES_PASS":   "root",
		"POSTGRES_SCHEMA": "service_management_postgres",
	},
}

var CONFIG = env[ENVIRONMENT]

func Getenv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func InitConfig() {
	for key := range CONFIG {
		CONFIG[key] = Getenv(key, CONFIG[key])
		os.Setenv(key, CONFIG[key])
	}
}
