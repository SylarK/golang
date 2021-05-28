// github.com/asaskevich/govalidator

package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/schema"
)

const (
	CONST_HOST         = "localhost"
	CONST_PORT         = "8080"
	USERNAME_ERROR_MSG = "Please enter a valid Username"
	PASSWORD_ERROR_MSG = "Please enter a valid Password"
	GENERIC_ERROR_MSG  = "Validation Error"
)

type User struct {
	Username string `valid:"alpha, required"`
	Password string `valid:"alpha, required"`
}

func readForm(r *http.Request) *User {
	r.ParseForm()
	user := new(User)
	decoder := schema.NewDecoder()
	decodeErr := decoder.Decode(user, r.PostForm)
	if decodeErr != nil {
		log.Printf("error mapping parsed form data to struct : ", decodeErr)
	}
	return user
}

func validationUser(w http.ResponseWriter, r *http.Request, user *User) (bool, string) {
	valid, validationError := govalidator.ValidateStruct(user)
	if !valid {
		usernameError := govalidator.ErrorByField(validationError, "Username")
		passwordError := govalidator.ErrorByField(validationError, "Password")
		if usernameError != "" {
			log.Printf("username validation error : ", usernameError)
			return valid, USERNAME_ERROR_MSG
		}
		if passwordError != "" {
			log.Printf("password validation error : ", passwordError)
			return valid, PASSWORD_ERROR_MSG
		}
	}
	return valid, GENERIC_ERROR_MSG
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parsedTemplate, _ := template.ParseFiles("templates/04.html")
		parsedTemplate.Execute(w, nil)
	} else {
		user := readForm(r)
		valid, validationErrorMessage := validationUser(w, r, user)
		if !valid {
			fmt.Fprintf(w, validationErrorMessage)
			return
		}
		fmt.Fprintf(w, "Hello "+user.Username+"!")
	}
}

func main() {
	http.HandleFunc("/", login)
	err := http.ListenAndServe(CONST_HOST+":"+CONST_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
