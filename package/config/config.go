package config

type Config struct {
	PostgresDsn string
	Port        string
}

func NewConfig() (config Config, err error) {
	config = Config{
		PostgresDsn: "host=localhost port=5432 user=robbo password=robbo_pwd dbname=robbo_db",
		Port:        ":8000",
	}
	return
}
