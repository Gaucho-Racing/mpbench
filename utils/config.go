package utils

import "mpbench/config"

func VerifyConfig() {
	if config.Port == "" {
		config.Port = "9999"
		SugarLogger.Infof("PORT is not set, defaulting to %s", config.Port)
	}
	if config.GithubAppID == "" {
		SugarLogger.Fatal("GITHUB_APP_ID is not set")
	}
	if config.GithubAppClientID == "" {
		SugarLogger.Fatal("GITHUB_APP_CLIENT_ID is not set")
	}
	if config.GithubAppAccessToken == "" {
		SugarLogger.Fatal("GITHUB_APP_ACCESS_TOKEN is not set")
	}
	if config.GithubAppInstallationID == "" {
		SugarLogger.Fatal("GITHUB_APP_INSTALLATION_ID is not set")
	}
	// if config.DatabaseHost == "" {
	// 	config.DatabaseHost = "localhost"
	// 	SugarLogger.Infof("DATABASE_HOST is not set, defaulting to %s", config.DatabaseHost)
	// }
	// if config.DatabasePort == "" {
	// 	config.DatabasePort = "3306"
	// 	SugarLogger.Infof("DATABASE_PORT is not set, defaulting to %s", config.DatabasePort)
	// }
	// if config.DatabaseUser == "" {
	// 	config.DatabaseUser = "root"
	// 	SugarLogger.Infof("DATABASE_USER is not set, defaulting to %s", config.DatabaseUser)
	// }
	// if config.DatabasePassword == "" {
	// 	config.DatabasePassword = "password"
	// 	SugarLogger.Infof("DATABASE_PASSWORD is not set, defaulting to %s", config.DatabasePassword)
	// }
}
