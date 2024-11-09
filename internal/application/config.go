package application

type Config struct {
	Service Service `yaml:"service"`
}

type Service struct {
	Name        string `yaml:"name"`
	Secret      string `yaml:"secret"`
	Environment string `yaml:"environment"`
}
