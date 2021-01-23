package main

import (
	"crypto/tls"
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/ldap.v2"
)

type LdapConfig struct {
	Url        string
	ServerName string
	Tls        bool
}

type LoginConfig struct {
	UserName string
	Password string
}

func loadConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}

func main() {
	err := loadConfig()
	if err != nil {
		fmt.Println("Unable to load config")
		panic(err)
	}

	config := LdapConfig{
		Url:        viper.GetString("ldap.url"),
		ServerName: viper.GetString("ldap.servername"),
		Tls:        viper.GetBool("ldap.tls"),
	}

	loginConfig := LoginConfig{
		UserName: viper.GetString("ldap.user"),
		Password: viper.GetString("ldap.password"),
	}

	l, err := ldap.DialTLS("tcp",
		config.Url,
		&tls.Config{ServerName: config.ServerName, InsecureSkipVerify: !config.Tls})
	if err != nil {
		fmt.Println("Failed to dial LDAP server")
		panic(err)
	}
	defer l.Close()

	l.Bind(loginConfig.UserName, loginConfig.Password)
	if err != nil {
		fmt.Println("Failed to login as admin")
		panic(err)
	}

	fmt.Println("Application started")
}
