package main

import (
	"html/template"
	"net/http"
)

// type Reqinfo struct {
// 	Request_name  string
// 	Request_email string
// }

type User struct {
	Email    string
	Password string
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
	http.HandleFunc("/", userHandler)
	http.ListenAndServe(":8880", nil)
}

func userHandler(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/sign-in":
		signInUser(w, r)
	case "/sign-up":
		signUpUser(w, r)
	case "/sign-in-form":
		getSignInPage(w, r)
	case "/sign-up-form":
		getSignUpPage(w, r)
	}
}

func signInUser(w http.ResponseWriter, r *http.Request) {

}

func signUpUser(w http.ResponseWriter, r *http.Request) {

}

func getUser(r *http.Request) User {
	email := r.FormValue("email")
	password := r.FormValue("password")
	return User{
		Email:    email,
		Password: password,
	}
}

func getSignInPage(w http.ResponseWriter, r *http.Request) {
	templating(w, "sign-in.html", nil)
}

func getSignUpPage(w http.ResponseWriter, r *http.Request) {
	templating(w, "sign-up.html", nil)
}

func templating(w http.ResponseWriter, fileName string, data interface{}) {
	t, _ := template.ParseFiles(fileName)
	t.ExecuteTemplate(w, fileName, data)
}
