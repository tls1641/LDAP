package Services

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

// //서비스 구독중인 병원 목록 조회(필요가 있을까??)
// func readHospitalList(l *ldap.Conn, filter string) (*ldap.SearchRequest, error) {
// 	l.Bind(BindUsername, BindPassword)

// 	searchReq := ldap.SearchRequest(
// 		BaseDN,
// 		ldap.ScopeWholeSubtree, // you can also use ldap.ScopeWholeSubtree
// 		ldap.NeverDerefAliases,
// 		0,
// 		0,
// 		false,
// 		filter,
// 		[]string{},
// 		nil,
// 	)

// 	result, err := l.Search(searchReq)
// 	if err != nil {
// 		return nil, fmt.Errorf("Search Error: %s", err)
// 	}

// 	if len(result.Entries) > 0 {
// 		return result, nil
// 	} else {
// 		return nil, fmt.Errorf("Couldn't fetch search entries")
// 	}
// }

//신규 서비스 등록
func CreateNewService(servicename string) {
	fmt.Println("Adding started")

	//Initialize connection
	l, err := ldap.Dial("tcp", FQDN)

	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	//Bind to the LDAP server
	bindusername := "cn=admin,dc=int,dc=trustnhope,dc=com"
	bindpassword := "admin"

	err = l.Bind(bindusername, bindpassword)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Testing.")

	//Create new Add request object to be added to LDAP server.
	a := ldap.NewAddRequest("cn="+ servicename + ",ou=services,dc=int,dc=trustnhope,dc=com", nil)
	a.Attribute("cn", []string{servicename})
	a.Attribute("member", []string{""})
	a.Attribute("objectClass", []string{"top", "groupOfNames"})

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

//서비스 구독중인 병원 추가
func AddServiceHospitalMember(hospitalCode string, servicename string){
	//Initialize connection
	l, err := ldap.Dial("tcp", FQDN)

	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	//Bind to the LDAP server
	bindusername := "cn=admin,dc=int,dc=trustnhope,dc=com"
	bindpassword := "admin"

	err = l.Bind(bindusername, bindpassword)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("connection ok")

	fmt.Println("adding new Hospital to ServiceMember...")
	m := ldap.NewModifyRequest("cn="+ servicename +",ou=services,dc=int,dc=trustnhope,dc=com",nil)

	searchReq := ldap.NewSearchRequest(
		BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(&(objectclass=groupOfNames)(cn="+ servicename +"))",
		[]string{},
		nil,
	)

	sr, err := l.Search(searchReq)
	if err != nil {
		return
	}

	vals := sr.Entries[0].GetAttributeValues("member")
	//members := make([]string, len(vals))
	for i, dn := range vals {
		fmt.Println(i,dn,"사용자 목록")
	}

	fmt.Println(vals)

	if vals[0] == "" {
		m.Replace("member", []string{"ou="+ hospitalCode +",ou=hospitals,dc=int,dc=trustnhope,dc=com"})
	}else{
		//그렇지 않다면 아래의 구문을 실행한다
		//member attribute를 추가한다
		m.Add("member", []string{"ou="+ hospitalCode +",ou=hospitals,dc=int,dc=trustnhope,dc=com"})
	}

	err1 := l.Modify(m)
	if err1 != nil{
		log.Fatal(err1)
	}
}

//서비스 구독 해지
func RemoveServiceHostpitalMember(hospitalCode string, servicename string){
	//Initialize connection
	l, err := ldap.Dial("tcp", FQDN)

	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	//Bind to the LDAP server
	bindusername := "cn=admin,dc=int,dc=trustnhope,dc=com"
	bindpassword := "admin"

	err = l.Bind(bindusername, bindpassword)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("connection ok")
//================================================================================================================
	m := ldap.NewModifyRequest("cn="+ servicename +",ou=services,dc=int,dc=trustnhope,dc=com",nil)

	// searchReq := ldap.NewSearchRequest(
	// 	BaseDN,
	// 	ldap.ScopeWholeSubtree,
	// 	ldap.NeverDerefAliases,
	// 	0,
	// 	0,
	// 	false,
	// 	"(&(objectclass=groupOfNames)(member=ou="+ hospitalCode +",ou=hospitals,dc=int,dc=trustnhope,dc=com))",
	// 	[]string{},
	// 	nil,
	// )

	// sr, err := l.Search(searchReq)
	// if err != nil {
	// 	return
	// }

	// vals := sr.Entries[0].GetAttributeValues("member")
	// //members := make([]string, len(vals))
	// for i, dn := range vals {
	// 	fmt.Println(i,dn,"사용자 목록")
	// }

	// fmt.Println(vals)

	m.Delete("member", []string{"ou="+ hospitalCode +",ou=hospitals,dc=int,dc=trustnhope,dc=com"})

	err1 := l.Modify(m)
	if err1 != nil{
		log.Fatal(err1)
	}
}

//서비스 삭제


//공통 연결
func initializeConnection() {
	fmt.Println("binding...")

	l, err := ldap.Dial("tcp", FQDN)

	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	//Bind to the LDAP server
	bindusername := "cn=admin,dc=int,dc=trustnhope,dc=com"
	bindpassword := "admin"

	err = l.Bind(bindusername, bindpassword)
	if err != nil {
		log.Fatal(err)
		return
	}
}