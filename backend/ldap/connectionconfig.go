package ldap

import (
	"log"

	"github.com/go-ldap/ldap/v3"
)

const (
	BindUsername = "CN=admin,DC=int,DC=trustnhope,DC=com"
	BindPassword = "admin"
	FQDN         = "118.67.131.11:3000" 
)

func DialAndBind(bindUsername string, bindPassword string) (l *ldap.Conn, err error) {
	l, dialError := ldap.Dial("tcp", FQDN)
	if dialError != nil {
		log.Fatal(err)
		return nil, err
	}

	err = l.Bind(bindUsername, bindPassword)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return l, nil
}

func Connect() (l *ldap.Conn, err error) {
	lc, err := DialAndBind(BindUsername, BindPassword)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return lc, nil
}