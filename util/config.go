package util

type Config struct {
	TokenSymmetricKey string
	ServerAddress     string
}

func LoadConfig(directory string) (*Config, error) {
	return &Config{TokenSymmetricKey: "fadil", ServerAddress: "0.0.0.0:8000"}, nil
}
