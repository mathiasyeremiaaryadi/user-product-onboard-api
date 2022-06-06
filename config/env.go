package config

import "os"

var configEnv = map[string]map[string]string{
	"development": {
		"APP_HOST":    "localhost:8080",
		"TOKEN_KEY":   "",
		"DB_HOST":     "localhost",
		"DB_PORT":     "3306",
		"DB_USERNAME": "root",
		"DB_PASSWORD": "",
		"DB_NAME":     "db_ex_rest_api",
	},
	"staging": {
		"APP_HOST":    "",
		"TOKEN_KEY":   "",
		"DB_HOST":     "",
		"DB_PORT":     "",
		"DB_USERNAME": "",
		"DB_PASSWORD": "",
		"DB_NAME":     "",
	},
	"production": {
		"APP_HOST":    "",
		"TOKEN_KEY":   "",
		"DB_HOST":     "",
		"DB_PORT":     "",
		"DB_USERNAME": "",
		"DB_PASSWORD": "",
		"DB_NAME":     "",
	},
}

const ENVIRONMENT_STATUS string = "development"

var ENV = configEnv[ENVIRONMENT_STATUS]

func GetConfig(key string, config string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return config
}

func InitConfig() {
	for key := range ENV {
		ENV[key] = GetConfig(key, ENV[key])
		os.Setenv(key, ENV[key])
	}
}
