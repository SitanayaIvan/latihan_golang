package domains

type ConfigBody struct {
	App AppAccount
	Db  DBAccount
}

type AppAccount struct {
	Port string
	Env  string
}

type DBAccount struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}
