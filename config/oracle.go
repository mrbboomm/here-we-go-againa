package config

type OracleConfig struct {
	Url string
	Port int
	ServiceName string
	User string
	Password string
	Options map[string]string
}

func LocalOracleConfig() OracleConfig {
	return OracleConfig{
		Url: "localhost",
		Port: 1521,
		ServiceName: "godev",
		User: "godev_user",
		Password: "godev_pass",
		Options: map[string]string{
			"CONNECTION TIMEOUT": "90",
			"SSL": "false",
		},
	}
}