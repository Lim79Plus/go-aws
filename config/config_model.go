package config

// Config struct
type Config struct {
	App App `yaml:"app"`
	Aws Aws `yaml:"aws"`
}
