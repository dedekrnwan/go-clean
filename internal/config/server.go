package config

type Server struct {
	AppName string `json:"app_name"`
	AppKey  string `json:"app_key"`
	Port    string `json:"port"`
	Version string `json:"version"`
}
