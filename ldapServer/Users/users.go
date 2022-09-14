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
)

// 유저 추가
func CreateUser(sn string, cn string, pw string, uid string, hospitalCode string) string {
	dn := "uid=" + uid + ",ou=" + hospitalCode + ",ou=hospitals,dc=int,dc=trustnhope,dc=com"
	fmt.Println("create user")

	l, err := ldapServer.DialAndBind(BindUsername, BindPassword)

	if err != nil {
		log.Fatal(err)
		return err.Error()
	}

	a := ldap.NewAddRequest(dn, nil)
	a.Attribute("sn", []string{sn})
	a.Attribute("cn", []string{cn})
	a.Attribute("objectClass", []string{"top", "inetOrgPerson", "organizationalPerson", "person"})
	a.Attribute("userPassword", []string{pw})

	ldapServer.Add(a, l)
	return "success" + uid
}

func ReadUserDN(baseDN string, uid string) (string, string) {
	filter := "(uid=" + uid + ")"
	BaseDN := "ou=" + baseDN + ",ou=hospitals,dc=int,dc=trustnhope,dc=com"
	l, err := ldapServer.DialAndBind(BindUsername, BindPassword)
	if err != nil {
		log.Fatal(err)
		return "err", err.Error()
	}

	result, err := ldapServer.Search(l, BaseDN, filter)
	if err != nil {
		log.Fatal(err)
		return "err", err.Error()
	}
	fmt.Println(result.Entries[0].DN, result.Entries[0].GetAttributeValue("cn"))
	return result.Entries[0].DN, result.Entries[0].GetAttributeValue("cn")
}

func ReadUserMember(uid string, ou string) ([]string, string) {
	filter := "(member=uid=" + uid + ",ou=" + ou + ",ou=hospitals,dc=int,dc=trustnhope,dc=com)"
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

	var departments = make([]string, len(result.Entries))
	for i := 0; i < len(result.Entries); i++ {
		// fmt.Println(*result.Entries[i])
		departments[i] = result.Entries[i].DN
	}
	fmt.Println(departments)
	// return departments
	return departments, ""
}

func UpdateUser(originalUid string, hospitalCode string, sn string, cn string) (bool, string) {

	l, err := ldapServer.DialAndBind(BindUsername, BindPassword)
	if err != nil {
		log.Fatal(err)
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
	l, err := ldapServer.DialAndBind(BindUsername, BindPassword)
	if err != nil {
		log.Fatal(err)
		return false, err.Error()
	}

	delete := ldap.NewDelRequest("uid="+uid+",ou="+hospitalCode+",ou=hospitals,dc=int,dc=trustnhope,dc=com", nil)

	err1 := l.Del(delete)
	if err1 != nil {
		log.Fatal(err1)
		return false, err.Error()
	}
	return true, uid + "delete success"
}
