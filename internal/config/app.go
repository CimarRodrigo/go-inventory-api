package config

type App struct {
	Name        string
	Version     string
	Environment string
}

func LoadAppConfig() App {
	return App{
		Name:        getEnv("APP_NAME", "My API"),
		Version:     getEnv("APP_VERSION", "1.0.0"),
		Environment: getEnv("APP_ENV", "development"), // development, staging, production
	}
}
