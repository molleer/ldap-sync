package services

type LdapConfig struct {
	Url        string
	ServerName string
	Tls        bool
}

type LoginConfig struct {
	UserName string
	Password string
}
