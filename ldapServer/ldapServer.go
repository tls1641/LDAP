package ldapServer

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap/v3"
)

const (
	BindUsername = "CN=admin,DC=int,DC=trustnhope,DC=com"
	BindPassword = "admin"
	FQDN         = "118.67.131.11:3000" //"192.168.163.129:389" //"20.196.153.228:3389"

)

func Add(addRequest *ldap.AddRequest, l *ldap.Conn) {
	err := l.Add(addRequest)
	if err != nil {
		fmt.Println("Entry NOT done", err)
	} else {
		fmt.Println("Entry DONE", err)
	}
}

func Modify(modifyRequest * ldap.ModifyRequest, l *ldap.Conn) {
	err := l.Modify(modifyRequest)
	if err != nil {
		log.Fatal(err)
	}
}

func Delete(deleteRequest *ldap.DelRequest, l *ldap.Conn) {
	err := l.Del(deleteRequest)
		if err != nil {
		log.Fatal(err)
	}
}

// Normal Bind and Search
func Search(l *ldap.Conn, BaseDN string, filter string) (*ldap.SearchResult, error) {
	searchReq := ldap.NewSearchRequest(
		BaseDN,
		ldap.ScopeWholeSubtree, // you can also use ldap.ScopeWholeSubtree
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		filter,
		[]string{},
		nil,
	)
	result, err := l.Search(searchReq)
	if err != nil {
		return nil, fmt.Errorf("Search Error: %s", err)
	}

	if len(result.Entries) > 0 {
		return result, nil
	} else {
		return nil, fmt.Errorf("Couldn't fetch search entries")
	}
}

func ConnectTLS() (*ldap.Conn, error) {
	// You can also use IP instead of FQDN
	l, err := ldap.Dial("tcp", FQDN)
	if err != nil {
		return nil, err
	}

	return l, nil
}

// Ldap Connection without TLS
func Connect() (*ldap.Conn, error) {
	// You can also use IP instead of FQDN
	l, err := ldap.Dial("tcp", FQDN)
	if err != nil {
		return nil, err
	}

	return l, nil
}

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