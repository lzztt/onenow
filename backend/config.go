package main

type config struct {
	Server struct {
		Port     string `yaml:"port"`
		KeyFile  string `yaml:"keyFile"`
		CertFile string `yaml:"certFile"`
	} `yaml:"server"`

	Data struct {
		NoteDir string `yaml:"noteDir"`
	} `yaml:"data"`

	Secret struct {
		AllowedEmail string `env:"ALLOWED_EMAIL"`
	}
}
