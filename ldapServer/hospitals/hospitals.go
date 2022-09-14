package Hospitals

import (
	"fmt"
	"log"
	"project/main/ldapServer"
	"reflect"

	"github.com/go-ldap/ldap/v3"
)

const (
	BindUsername = "CN=admin,DC=int,DC=trustnhope,DC=com"
	BindPassword = "admin"
	IP           = "118.67.131.11:3000" //"192.168.163.129:389" //"20.196.153.228:3389"
)

func CreateHospital(ou string) {
	dn := "ou=" + ou + ",ou=hospitals,dc=int,dc=trustnhope,dc=com"
	fmt.Println("create hospital")

	l, err := ldap.Dial("tcp", IP)

	if err != nil {
		log.Fatal(err)
	}

	err = l.Bind(BindUsername, BindPassword)
	if err != nil {
		log.Fatal(err)
	}

	a := ldap.NewAddRequest(dn, nil)
	a.Attribute("objectClass", []string{"top", "organizationalUnit"})

	ldapServer.Add(a, l)
}

func ReadHospitalMember(ou string) {
	filter := "(member=ou=" + ou + ",ou=hospitals,dc=int,dc=trustnhope,dc=com)"
	BaseDN := "dc=int,dc=trustnhope,dc=com"

	l, err := ldapServer.ConnectTLS()
	if err != nil {
		log.Fatal("connect", err)
	}

	result, err := ldapServer.Search(l, BaseDN, filter)
	if err != nil {
		log.Fatal(err)
	}

	var services = make([]any, len(result.Entries))
	for i := 0; i < len(result.Entries); i++ {
		// fmt.Println(*result.Entries[i])
		services[i] = result.Entries[i].GetAttributeValue("cn")
	}
	fmt.Println(reflect.TypeOf(services[1]))
	// return departments

}

func UpdateHospitalMember() {

}

func DeleteHospitalMember(ou string) (bool, string) {
	l, err := ldapServer.ConnectTLS()
	if err != nil {
		log.Fatal("connect", err)
		return false, err.Error()
	}
	fmt.Println(l)

	l.Bind(BindUsername, BindPassword)

	delete := ldap.NewDelRequest("ou="+ou+",ou=hospitals,dc=int,dc=trustnhope,dc=com", nil)

	err1 := l.Del(delete)
	if err1 != nil {
		log.Fatal(err1)
		return false, err.Error()
	}
	return true, ou + "delete success"
}
