package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project/main/ldapServer"

	"project/main/ldapServer/Departments"
	"project/main/ldapServer/Services"
	"project/main/ldapServer/Users"

	"github.com/gorilla/mux"
	"time"

)

// type Reqinfo struct {
// 	Request_name  string
// 	Request_email string
// }

// type User struct {
// 	Email    string
// 	Password string
// }

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request:", err)
		return
	}
	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)

}

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
	Departments.AddMember("t00001","departmentB","members","hiosi")
	//Departments.DeleteMember("t00001","departmentA","members","hiosi")
	Services.ReadHospitalList("hiosk")
	//Services.CreateNewService("medic-app")
	//Services.AddServiceHospitalMember("c00047","medic")
	//Services.RemoveServiceHostpitalMember("c00047","medic")


	// Users.CreateUser("Mr.", "testperson3", "1234", "tls16411", "t00002")
	// Users.ReadUserDN("t00002", "person3")
	// Users.ReadUserMember("hiosi", "t00001")


	//Users.UpdateUser("person5", "t00002", "test5", "changedtest5")

	// Users.DeleteUser("person10", "t00002")
	// http.HandleFunc("/", userHandler)
	// Hospitals.CreateHospital("t00003")
	// Hospitals.DeleteHospitalMember("t00003")

	//절대 경로
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})
	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", mux)

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
