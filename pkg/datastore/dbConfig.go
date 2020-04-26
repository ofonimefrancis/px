package datastore

type Config struct {
	AppName  string `mapstructure:app_name"`
	Port     string `mapstructure:"app_port"`
	DBConfig DBConfig
}

type DBConfig struct {
	DatabaseName    string `mapstructure:"database_name"`
	DatabaseURI     string `mapstructure:"database_uri"`
	DatabaseTimeout int    `mapstructure:"database_timeout"`
}
