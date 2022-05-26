package types

type Env struct {
	Host         string `yaml:"pg_host"`
	User         string `yaml:"pg_user"`
	password     string `yaml:"pg_password"`
	DatabaseName string `yaml:"pg_database_name"`
	Port         string `yaml:"pg_port"`
}
