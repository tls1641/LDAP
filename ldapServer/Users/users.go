package Users

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

// 유저 추가
func CreateUser(sn string, cn string, pw string, uid string, hospitalCode string) bool {
	dn := "uid=" + uid + ",ou=" + hospitalCode + ",ou=hospitals,dc=int,dc=trustnhope,dc=com"
	fmt.Println("create user")

	l, err := ldap.Dial("tcp", IP)

	if err != nil {
		log.Fatal(err)
	}

	err = l.Bind(BindUsername, BindPassword)
	if err != nil {
		log.Fatal(err)
		return false
	}

	a := ldap.NewAddRequest(dn, nil)
	a.Attribute("sn", []string{sn})
	a.Attribute("cn", []string{cn})
	a.Attribute("objectClass", []string{"top", "inetOrgPerson", "organizationalPerson", "person"})
	a.Attribute("userPassword", []string{pw})

	ldapServer.Add(a, l)
	return true
}

func ReadUserDN(baseDN string, uid string) (string, string) {
	filter := "(uid=" + uid + ")"
	BaseDN := "ou=" + baseDN + ",ou=hospitals,dc=int,dc=trustnhope,dc=com"
	l, err := ldapServer.ConnectTLS()
	if err != nil {
		log.Fatal("connect", err)
	}

	result, err := ldapServer.Search(l, BaseDN, filter)
	if err != nil {
		log.Fatal(err)
		return "err", err.Error()
	}

	return result.Entries[0].DN, result.Entries[0].GetAttributeValue("cn")
}

func ReadUserMember(dn string) []interface{} {
	filter := "(member=" + dn + ")"
	BaseDN := "dc=int,dc=trustnhope,dc=com"

	l, err := ldapServer.ConnectTLS()
	if err != nil {
		log.Fatal("connect", err)
	}

	result, err := ldapServer.Search(l, BaseDN, filter)
	if err != nil {
		log.Fatal(err)
	}

	var departments = make([]any, len(result.Entries))
	for i := 0; i < len(result.Entries); i++ {
		// fmt.Println(*result.Entries[i])
		departments[i] = result.Entries[i].GetAttributeValue("cn")
	}
	fmt.Println(reflect.TypeOf(departments[1]))
	// return departments
	return departments
}

func UpdateUser(originalUid string, hospitalCode string, sn string, cn string) (bool, string) {

	l, err := ldapServer.ConnectTLS()
	if err != nil {
		log.Fatal("connect", err)
		return false, err.Error()
	}
	fmt.Println(l)

	l.Bind(BindUsername, BindPassword)

	modify := ldap.NewModifyRequest("uid="+originalUid+",ou="+hospitalCode+",ou=hospitals,dc=int,dc=trustnhope,dc=com", nil)

	// if newUid != "" {
	// 	modify.Replace("uid", []string{newUid})
	// }

	if sn != "" {
		modify.Replace("sn", []string{sn})
	}

	if cn != "" {
		modify.Replace("cn", []string{cn})
	}

	err1 := l.Modify(modify)
	if err1 != nil {
		log.Fatal(err1)
		return false, err.Error()
	}

	return true, originalUid + "modify success"
}

func DeleteUser(uid string, hospitalCode string) (bool, string) {
	l, err := ldapServer.ConnectTLS()
	if err != nil {
		log.Fatal("connect", err)
		return false, err.Error()
	}
	fmt.Println(l)

	l.Bind(BindUsername, BindPassword)

	delete := ldap.NewDelRequest("uid="+uid+",ou="+hospitalCode+",ou=hospitals,dc=int,dc=trustnhope,dc=com", nil)

	err1 := l.Del(delete)
	if err1 != nil {
		log.Fatal(err1)
		return false, err.Error()
	}
	return true, uid + "delete success"
}
