package ldap

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

func Add(l *ldap.Conn, addRequest *ldap.AddRequest) (string, error) {
	err := l.Add(addRequest)
	if err != nil {
		fmt.Println("Entry NOT done", err)
		return "fail", fmt.Errorf("Add Error: %s", err)
	} else {
		fmt.Println("Entry DONE", err)
		return "ok", nil
	}
}