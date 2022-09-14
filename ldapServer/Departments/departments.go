package Departments

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
)


//특정 병원의 특정 부서에 소속된 모든 사람들 목록 조회하기
// func ReadAllMembers(hospitalCode string, departmentname string) ([]string, error) {
// 	l, err := ldapServer.DialAndBind(BindUsername,BindPassword)

// 	if err != nil {
// 		log.Fatal(err)
// 		return []string{}, err
// 	}

// 	adminSearch, adminSearchError := ldapServer.Search(
// 		l,
// 		"cn=admin,ou="+ departmentname +"ou=department,dc=int,dc=trustnhope,dc=com",
// 		"(&(objectclass=groupOfNames)(member=))",
// 	)
// 	if adminSearchError != nil {
// 		return []string{}, adminSearchError
// 	}

// 	// return adminSearch
// }

//특정 병원의 특정 부서의 관리자 조회하기
//특정 병원, 특정 부서의 관리자 제외한 팀원 조회하기

//특정 부서에 새로운 사람 추가
func AddMember(hospitalCode string, departmentname string, position string, uid string) (string, error) {
	fmt.Println("추가중",hospitalCode,departmentname,position,uid)
	l, err := ldapServer.DialAndBind(BindUsername,BindPassword)

	if err != nil {
		fmt.Println("1",err)
		log.Fatal(err)
		return "fail", err
	}
	m := ldap.NewModifyRequest("cn="+ position +",ou="+ departmentname +",ou=department,dc=int,dc=trustnhope,dc=com", nil)
	s, searchError := ldapServer.Search(
		l,
		BaseDN,
		"(&(objectClass=groupOfNames)(dn=cn="+position+",ou="+departmentname+",ou=department,dc=int,dc=trustnhope,dc=com))",
	)
	if searchError != nil {
		fmt.Println("2",searchError) //원하는 부서, 직책 찾지 못함
		return "fail", searchError
	}

	vals := s.Entries[0].GetAttributeValues("member")
	fmt.Println(vals,"값 확인&&")
	if vals[0] == "" {
		m.Replace("member", []string{"uid="+uid+",ou="+hospitalCode+",ou=hospitals,dc=int,dc=trustnhope,dc=com"})
	} else {
		m.Add("member", []string{"uid="+uid+",ou="+hospitalCode+",ou=hospitals,dc=int,dc=trustnhope,dc=com"})
	}
	modifyResult, modifyError := ldapServer.Modify(m,l)
	if modifyError != nil {
		fmt.Println("3",modifyError)
		return modifyResult, modifyError
	}else{
		return modifyResult, nil
	}
}
//특정 병원, 특정 부서에 있는 기존 구성원의 정보 변경
//특정 병원, 특정 부서의 특정인 제거
func DeleteMember(hospitalCode string, departmentname string, position string, uid string) (string, error) {
	fmt.Println("제거중",hospitalCode,departmentname,position,uid)
	l, err := ldapServer.DialAndBind(BindUsername,BindPassword)

	if err != nil {
		fmt.Println("1",err)
		log.Fatal(err)
		return "fail", err
	}

	m := ldap.NewModifyRequest("cn="+ position +",ou="+ departmentname +",ou=department,dc=int,dc=trustnhope,dc=com", nil)
	s, searchError := ldapServer.Search(
		l,
		BaseDN,
		"(&(objectclass=groupOfNames)(member=uid="+uid+",ou="+hospitalCode+",ou=hospitals,dc=int,dc=trustnhope,dc=com))",
	)
	if searchError != nil {
		fmt.Println("2",searchError)
		return "fail", searchError
	}
	vals := s.Entries[0].GetAttributeValues("member")
	fmt.Println(vals,"값 확인&&")
	if len(vals) == 1 {
		m.Replace("member", []string{""})
	} else {
		m.Delete("member", []string{"uid="+uid+",ou="+ hospitalCode +",ou=hospitals,dc=int,dc=trustnhope,dc=com"})
	}
	modifyResult, modifyError := ldapServer.Modify(m,l)
	if modifyError != nil {
		fmt.Println("3",modifyError)
		return modifyResult, modifyError
	}else{
		return modifyResult, nil
	}
}


