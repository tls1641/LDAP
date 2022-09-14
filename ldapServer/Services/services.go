package Services

import (
	"fmt"
	"log"
	"project/main/ldapServer"

	"github.com/go-ldap/ldap/v3"
)

const (
	BindUsername = "CN=admin,DC=int,DC=trustnhope,DC=com"
	BindPassword = "admin"
	FQDN         = "118.67.131.11:3000" //"192.168.163.129:389" //"20.196.153.228:3389"
	BaseDN       = "dc=int,dc=trustnhope,dc=com"
	Filter       = "(&(objectclass=inetOrgPerson)(gidNumber=1000))"
)

// 서비스 구독중인 병원 목록 조회(필요가 있을까??)
// https://golangdocs.com/multiple-return-values-in-golang-functions
func ReadHospitalList(servicename string) ([]string, error) {
	l, err := ldapServer.DialAndBind(BindUsername,BindPassword)

	if err != nil {
		log.Fatal(err)
		return []string{}, err
	}

	s, searchError := ldapServer.Search(
		l,
		"cn="+servicename+",ou=services,dc=int,dc=trustnhope,dc=com",
		"(objectclass=groupOfNames)",
	)

	if searchError != nil {
		return []string{}, searchError
	}
	fmt.Println(s.Entries[0].GetAttributeValues("member"))
	return s.Entries[0].GetAttributeValues("member"), nil
}

// 신규 서비스 등록
func CreateNewService(servicename string) {
	l, err := ldapServer.DialAndBind(BindUsername,BindPassword)

	if err != nil {
		log.Fatal(err)
		return
	}

	//Create new Add request object to be added to LDAP server.
	a := ldap.NewAddRequest("cn="+ servicename +",ou=services,dc=int,dc=trustnhope,dc=com", nil)
	a.Attribute("cn", []string{servicename})
	a.Attribute("member", []string{""})
	a.Attribute("objectClass", []string{"top", "groupOfNames"})

	fmt.Println("Testing.",l)

	ldapServer.Add(a, l)
}

// 서비스 구독중인 병원 추가
func AddServiceHospitalMember(hospitalCode string, servicename string) {
	l, err := ldapServer.DialAndBind(BindUsername,BindPassword)

	if err != nil {
		log.Fatal(err)
		return
	}

	m := ldap.NewModifyRequest("cn="+servicename+",ou=services,dc=int,dc=trustnhope,dc=com", nil)
	s, searchError := ldapServer.Search(
		l,
		BaseDN,
		"(&(objectclass=groupOfNames)(cn="+servicename+"))",
	)

	if searchError != nil {
		log.Fatal(searchError)
		return
	}

	vals := s.Entries[0].GetAttributeValues("member")
	if vals[0] == "" {
		m.Replace("member", []string{"ou=" + hospitalCode + ",ou=hospitals,dc=int,dc=trustnhope,dc=com"})
	} else {
		m.Add("member", []string{"ou=" + hospitalCode + ",ou=hospitals,dc=int,dc=trustnhope,dc=com"})
	}

	ldapServer.Modify(m,l)
}

// 서비스 구독 해지:: 서비스 목록에서 병원 member 삭제
func RemoveServiceHostpitalMember(hospitalCode string, servicename string) {
	l, err := ldapServer.DialAndBind(BindUsername,BindPassword)

	if err != nil {
		log.Fatal(err)
		return
	}

	m := ldap.NewModifyRequest("cn="+servicename+",ou=services,dc=int,dc=trustnhope,dc=com", nil)
	s, searchError := ldapServer.Search(
		l,
		BaseDN,
		"(&(objectclass=groupOfNames)(member=ou="+hospitalCode+",ou=hospitals,dc=int,dc=trustnhope,dc=com))",
	)

	if searchError != nil {
		return
	}
	vals := s.Entries[0].GetAttributeValues("member")
	if len(vals) == 1 {
		m.Replace("member", []string{""})
	} else {
		m.Delete("member", []string{"ou=" + hospitalCode + ",ou=hospitals,dc=int,dc=trustnhope,dc=com"})
	}

	ldapServer.Modify(m,l)
}

// 서비스 삭제(엔트리 자체를 삭제)
func DeleteService(servicename string) {
	l, err := ldapServer.DialAndBind(BindUsername,BindPassword)
	if err != nil {
		log.Fatal(err)
		return
	}

	d := ldap.NewDelRequest("cn="+servicename+",ou=services,dc=int,dc=trustnhope,dc=com", nil)
	ldapServer.Delete(d, l)
}


