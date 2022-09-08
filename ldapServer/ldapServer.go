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
	BaseDN       = "dc=int,dc=trustnhope,dc=com"
	Filter       = "(&(objectclass=inetOrgPerson)(gidNumber=1000))"
)

func addEntries() {
	fmt.Println("Adding started")

	//Initialize connection
	l, err := ldap.Dial("tcp", FQDN)

	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	//Bind to the LDAP server
	bindusername := "cn=admin,dc=trustnhope,dc=com"
	bindpassword := "admin"

	err = l.Bind(bindusername, bindpassword)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Testing.")

	//Create new Add request object to be added to LDAP server.
	a := ldap.NewAddRequest("cn=test2,ou=groups,dc=trustnhope,dc=com", nil)
	a.Attribute("cn", []string{"gotest"})
	a.Attribute("objectClass", []string{"top", "inetOrgPerson"})
	a.Attribute("description", []string{"this is a test to add an entry using golang"})
	a.Attribute("sn", []string{"Google"})

	fmt.Println("Testing.")
	add(a, l)

}
func add(addRequest *ldap.AddRequest, l *ldap.Conn) {
	err := l.Add(addRequest)
	if err != nil {
		fmt.Println("Entry NOT done", err)
	} else {
		fmt.Println("Entry DONE", err)
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

// Anonymous Bind and Search
func AnonymousBindAndSearch(l *ldap.Conn) (*ldap.SearchResult, error) {
	l.UnauthenticatedBind("")

	anonReq := ldap.NewSearchRequest(
		"",
		ldap.ScopeBaseObject, // you can also use ldap.ScopeWholeSubtree
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		Filter,
		[]string{},
		nil,
	)
	result, err := l.Search(anonReq)
	if err != nil {
		return nil, fmt.Errorf("Anonymous Bind Search Error: %s", err)
	}

	if len(result.Entries) > 0 {
		result.Entries[0].Print()
		return result, nil
	} else {
		return nil, fmt.Errorf("Couldn't fetch anonymous bind search entries")
	}
}

// Normal Bind and Search
func BindAndSearch(l *ldap.Conn, filter string) (*ldap.SearchResult, error) {
	l.Bind(BindUsername, BindPassword)

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
