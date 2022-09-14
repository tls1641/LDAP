package Users

import (
	"fmt"
	"log"
	"project/main/ldapServer"

	"github.com/go-ldap/ldap/v3"
)

const (
	BindUsername = "CN=admin,DC=int,DC=trustnhope,DC=com"
	BindPassword = "admin"
	IP           = "118.67.131.11:3000" //"192.168.163.129:389" //"20.196.153.228:3389"
	BaseDN       = "dc=int,dc=trustnhope,dc=com"
	Filter       = "(&(objectclass=inetOrgPerson)(gidNumber=1000))"
)

// 유저 추가
func CreateUser(sn string, cn string, pw string, dn string) {
	fmt.Println("create user")

	l, err := ldap.Dial("tcp", IP)

	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	err = l.Bind(BindUsername, BindPassword)
	if err != nil {
		log.Fatal(err)
		return
	}

	a := ldap.NewAddRequest(dn, nil)
	a.Attribute("sn", []string{sn})
	a.Attribute("cn", []string{cn})
	a.Attribute("objectClass", []string{"top", "inetOrgPerson", "organizationalPerson", "person"})
	a.Attribute("userPassword", []string{pw})

	ldapServer.Add(a, l)

}

func ReadUserDN(baseDN string, uid string) {
	filter := "(uid=" + uid + ")"
	BaseDN := "ou=" + baseDN + ",ou=hospitals,dc=int,dc=trustnhope,dc=com"
	l, err := ldapServer.ConnectTLS()
	if err != nil {
		log.Fatal("connect", err)
	}

	result, err := ldapServer.BindAndSearch(l, BaseDN, filter)
	if err != nil {
		log.Fatal(err)
	}
	result.Entries[0].Print()

}

func ReadUserMember(dn string) {
	filter := "(member=" + dn + ")"
	BaseDN := "dc=int,dc=trustnhope,dc=com"

	l, err := ldapServer.ConnectTLS()
	if err != nil {
		log.Fatal("connect", err)
	}

	result, err := ldapServer.BindAndSearch(l, BaseDN, filter)
	if err != nil {
		log.Fatal(err)
	}

	var departments = make([]any, len(result.Entries))
	for i := 0; i < len(result.Entries); i++ {
		// fmt.Println(*result.Entries[i])
		departments[i] = result.Entries[i].GetAttributeValue("cn")
	}
	fmt.Println(departments)
}

func UpdateUser() {
	l, err := ldapServer.ConnectTLS()
	if err != nil {
		log.Fatal("connect", err)
	}
	fmt.Println(l)

}

func DeleteUser() {

}
