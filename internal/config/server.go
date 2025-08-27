package config

type Server struct {
	Port string
	Mode string
}

func LoadServerConfig() Server {
	return Server{
		Port: getEnv("PORT", "8080"),
		Mode: getEnv("GIN_MODE", "debug"), // debug, release, test
	}
}
