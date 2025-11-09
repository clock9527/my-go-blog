package config

type DBMysql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *DBMysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
