package Hospitals

import (
	"fmt"
	"log"
	"project/main/ldapServer"

	"github.com/go-ldap/ldap/v3"
)

const (
	BindUsername = "CN=admin,DC=int,DC=trustnhope,DC=com"
	BindPassword = "admin"
)

func CreateHospital(ou string) string {
	dn := "ou=" + ou + ",ou=hospitals,dc=int,dc=trustnhope,dc=com"

	l, err := ldapServer.DialAndBind(BindUsername, BindPassword)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}

	a := ldap.NewAddRequest(dn, nil)
	a.Attribute("objectClass", []string{"top", "organizationalUnit"})

	ldapServer.Add(a, l)
	return "success" + ou

}

func ReadHospitalMember(ou string) ([]string, string) {
	filter := "(member=ou=" + ou + ",ou=hospitals,dc=int,dc=trustnhope,dc=com)"
	BaseDN := "dc=int,dc=trustnhope,dc=com"

	l, err := ldapServer.DialAndBind(BindUsername, BindPassword)
	if err != nil {
		log.Fatal(err)
		return nil, err.Error()
	}

	result, err := ldapServer.Search(l, BaseDN, filter)
	if err != nil {
		log.Fatal(err)
		return nil, err.Error()
	}

	var services = make([]string, len(result.Entries))
	for i := 0; i < len(result.Entries); i++ {
		// fmt.Println(*result.Entries[i])
		services[i] = result.Entries[i].DN
	}
	fmt.Println(services)
	// return departments
	return services, ""

}

func UpdateHospitalMember() {

}

func DeleteHospitalMember(ou string) (bool, string) {
	l, err := ldapServer.DialAndBind(BindUsername, BindPassword)
	if err != nil {
		log.Fatal(err)
		return false, err.Error()
	}

	delete := ldap.NewDelRequest("ou="+ou+",ou=hospitals,dc=int,dc=trustnhope,dc=com", nil)

	err1 := l.Del(delete)
	if err1 != nil {
		log.Fatal(err1)
		return false, err.Error()
	}
	return true, ou + "delete success"
}
