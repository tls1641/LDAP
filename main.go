package main

import (
	"fmt"
	"log"
	"net/http"
	"project/main/ldapServer"
	"project/main/ldapServer/Services"
	"project/main/ldapServer/Users"

)

// type Reqinfo struct {
// 	Request_name  string
// 	Request_email string
// }

// type User struct {
// 	Email    string
// 	Password string
// }

func main() {

	// http.HandleFunc("/", func(wr http.ResponseWriter, r *http.Request) {
	// 	wr.Write([]byte("This is main page."))
	// })

	// http.HandleFunc("/json", func(wr http.ResponseWriter, r *http.Request) {
	// 	if r.Method == http.MethodPost {
	// 		var user_info Reqinfo

	// 		json.NewDecoder(r.Body).Decode(&user_info)
	// 		fmt.Println(user_info)

	// 		json.NewEncoder(wr).Encode(user_info)
	// 	}
	// })
	l, err := ldapServer.ConnectTLS()
	if err != nil {
		log.Fatal("connect", err)
	}
	fmt.Println(l)


	// http.HandleFunc("/", userHandler)

	// Services.CreateNewService("medic-app")
	// Services.AddServiceHospitalMember("c00040","medic-app")
	// Services.RemoveServiceHostpitalMember("c00040","medic-app")

	// Users.CreateUser("Mr.", "testperson3", "1234", "uid=person3,ou=t00002,ou=hospitals,dc=int,dc=trustnhope,dc=com")
	// Users.ReadUserDN("t00002", "person3")
	// Users.ReadUserMember("uid=hiosi,ou=t00001,ou=hospitals,dc=int,dc=trustnhope,dc=com")
	// http.HandleFunc("/", userHandler)

	http.ListenAndServe("", nil)
}

// func userHandler(w http.ResponseWriter, r *http.Request) {

// 	switch r.URL.Path {
// 	case "/sign-in":
// 		signInUser(w, r)
// 	case "/sign-up":
// 		signUpUser(w, r)
// 	case "/sign-in-form":
// 		getSignInPage(w, r)
// 	case "/sign-up-form":
// 		getSignUpPage(w, r)
// 	}
// }

// func signInUser(w http.ResponseWriter, r *http.Request) {
// 	newUser := getUser(r)

// 	ok := userLog.DefaultUserService.VerifyUser(newUser)

// 	if !ok {
// 		fileName := "sign-in.html"
// 		t, _ := template.ParseFiles(fileName)
// 		t.ExecuteTemplate(w, fileName, "User Sign-in Failure")
// 		return
// 	}

// 	fmt.Println("로그인 성공")

// 	fileName := "sign-in.html"
// 	t, _ := template.ParseFiles(fileName)
// 	t.ExecuteTemplate(w, fileName, "User Sign-in Success")
// 	return
// }

// func signUpUser(w http.ResponseWriter, r *http.Request) {
// 	newUser := getUser(r)

// 	err := userLog.DefaultUserService.CreateUser(newUser)

// 	if err != nil {
// 		fileName := "sign-up.html"
// 		t, _ := template.ParseFiles(fileName)
// 		t.ExecuteTemplate(w, fileName, "New User Sign-up Failure")
// 	}

// 	fmt.Println("회원가입 성공")


// 	fileName := "sign-up.html"
// 	t, _ := template.ParseFiles(fileName)
// 	t.ExecuteTemplate(w, fileName, "New User Sign-up Success")
// }


// func getUser(r *http.Request) userLog.User {
// 	email := r.FormValue("email")
// 	password := r.FormValue("password")
// 	return userLog.User{

// 		Email:    email,
// 		Password: password,
// 	}
// }

// func getSignInPage(w http.ResponseWriter, r *http.Request) {
// 	templating(w, "sign-in.html", nil)
// }

// func getSignUpPage(w http.ResponseWriter, r *http.Request) {
// 	templating(w, "sign-up.html", nil)
// }

// func templating(w http.ResponseWriter, fileName string, data interface{}) {
// 	t, _ := template.ParseFiles(fileName)
// 	t.ExecuteTemplate(w, fileName, data)
// }
