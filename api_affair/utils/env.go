package utils

import "os"

const (
	defaultFilmServiceHost = "localhost"
	defaultFilmServicePort = "8080"

	defaultAdServiceHost = "localhost"
	defaultAdServicePort = "8083"

	defaultCDNAddress = "http://localhost"
)

func EnvFilmServiceAddress() string {
	var host string
	var port string
	if host = os.Getenv("FILM_SERVICE_HOST"); len(host) == 0 {
		host = defaultFilmServiceHost
	}
	if port = os.Getenv("FILM_SERVICE_PORT"); len(port) == 0 {
		port = defaultFilmServicePort
	}
	return host + ":" + port
}

func EnvAdServiceAddress() string {
	var host string
	var port string
	if host = os.Getenv("AD_SERVICE_HOST"); len(host) == 0 {
		host = defaultAdServiceHost
	}
	if port = os.Getenv("AD_SERVICE_PORT"); len(port) == 0 {
		port = defaultAdServicePort
	}
	return host + ":" + port
}

func EnvCDNAddress() string {
	var address string
	if address = os.Getenv("CDN_ADDRESS"); len(address) == 0 {
		address = defaultCDNAddress
	}
	return address
}
