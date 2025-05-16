package config

import "os"

var Version = "1.4.9"
var Env = os.Getenv("ENV")
var Port = os.Getenv("PORT")

var MapacheRepo = "gaucho-racing/mapache"

var DatabaseHost = os.Getenv("DATABASE_HOST")
var DatabasePort = os.Getenv("DATABASE_PORT")
var DatabaseUser = os.Getenv("DATABASE_USER")
var DatabasePassword = os.Getenv("DATABASE_PASSWORD")
var DatabaseName = os.Getenv("DATABASE_NAME")

var GithubAppID = os.Getenv("GITHUB_APP_ID")
var GithubAppClientID = os.Getenv("GITHUB_APP_CLIENT_ID")
var GithubAppAccessToken = os.Getenv("GITHUB_APP_ACCESS_TOKEN")
var GithubAppInstallationID = os.Getenv("GITHUB_APP_INSTALLATION_ID")
var GithubAppPrivateKey = os.Getenv("GITHUB_APP_PRIVATE_KEY")

var MaxWorkers = os.Getenv("MAX_WORKERS")
