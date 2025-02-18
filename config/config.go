package config

import "os"

var Version = "1.0.0"
var Env = os.Getenv("ENV")
var Port = os.Getenv("PORT")

var MapacheRepo = "gaucho-racing/mapache"

var GithubAppID = os.Getenv("GITHUB_APP_ID")
var GithubAppClientID = os.Getenv("GITHUB_APP_CLIENT_ID")
var GithubAppAccessToken = os.Getenv("GITHUB_APP_ACCESS_TOKEN")
var GithubAppInstallationID = os.Getenv("GITHUB_APP_INSTALLATION_ID")
var GithubAppPrivateKey = os.Getenv("GITHUB_APP_PRIVATE_KEY")
