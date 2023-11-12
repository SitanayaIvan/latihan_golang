package domains

type ConfigBody struct {
	App AppAccount
	Db  DbAccount
}

type AppAccount struct {
	Env  string
	Port string
}

type DbAccount struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}
