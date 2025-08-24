package config

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func LoadDatabaseConfig() Database {
	return Database{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", "root"),
		Name:     getEnv("DB_NAME", "project_inventory"),
	}
}
