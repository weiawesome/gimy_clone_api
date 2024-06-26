package utils

import "os"

const (
	defaultMongoDbHost     = "localhost"
	defaultMongoDbPort     = "27017"
	defaultMongoDbUser     = "defaultUser"
	defaultMongoDbPassword = "defaultPassword"
	defaultMongoDb         = "DefaultAdvertisementDb"

	defaultServerProtocol = "tcp"
	defaultServerPort     = "8080"

	defaultRedisHost     = "localhost"
	defaultRedisPort     = "6379"
	defaultRedisPassword = "DefaultPassword"
	defaultRedisDb       = "0"
)

var protocolTypes = map[string]bool{"tcp": true, "udp": true}

func EnvMongoDbAddress() string {
	var host string
	var port string
	if host = os.Getenv("MONGODB_HOST"); len(host) == 0 {
		host = defaultMongoDbHost
	}
	if port = os.Getenv("MONGODB_PORT"); len(port) == 0 {
		port = defaultMongoDbPort
	}
	return host + ":" + port
}
func EnvMongoDbUser() string {
	var user string
	if user = os.Getenv("MONGODB_USER"); len(user) == 0 {
		user = defaultMongoDbUser
	}
	return user
}
func EnvMongoDbPassword() string {
	var password string
	if password = os.Getenv("MONGODB_PASSWORD"); len(password) == 0 {
		password = defaultMongoDbPassword
	}
	return password
}
func EnvMongoDb() string {
	var database string
	if database = os.Getenv("MONGODB_DATABASE"); len(database) == 0 {
		database = defaultMongoDb
	}
	return database
}
func EnvServerProtocol() string {
	var protocol string
	if protocol = os.Getenv("SERVER_PROTOCOL"); len(protocol) == 0 || !protocolTypes[protocol] {
		protocol = defaultServerProtocol
	}
	return protocol
}
func EnvServerPort() string {
	var port string
	if port = os.Getenv("SERVER_PORT"); len(port) == 0 {
		port = defaultServerPort
	}
	return port
}

func EnvRedisAddress() string {
	var ip string
	var port string
	if ip = os.Getenv("REDIS_HOST"); len(ip) == 0 {
		ip = defaultRedisHost
	}
	if port = os.Getenv("REDIS_PORT"); len(port) == 0 {
		port = defaultRedisPort
	}
	return ip + ":" + port
}

func EnvRedisPassword() string {
	var password string
	if password = os.Getenv("REDIS_PASSWORD"); len(password) == 0 {
		password = defaultRedisPassword
	}
	return password
}

func EnvRedisDb() string {
	var db string
	if db = os.Getenv("REDIS_DB"); len(db) == 0 {
		db = defaultRedisDb
	}
	return db
}
