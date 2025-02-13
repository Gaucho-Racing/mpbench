package config

import "os"

var Version = "1.0.0"
var Env = os.Getenv("ENV")
var Port = os.Getenv("PORT")

var MapacheRepo = "https://github.com/gaucho-racing/mapache"
