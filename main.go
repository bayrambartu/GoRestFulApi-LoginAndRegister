package main

import (
	"fmt"
	"module/helpers"
	"net/http"
)

func main() {

	uName, email, pwd, pwdConfirm := "", "", "", ""

	mux := http.NewServeMux()
	// SignUp
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		uName = r.FormValue("username")
		email = r.FormValue("email")
		pwd = r.FormValue("password")
		pwdConfirm = r.FormValue("confirm")

		uNameCheck := helpers.IsEmpty(uName)
		emailCheck := helpers.IsEmpty(email)
		pwdCheck := helpers.IsEmpty(pwd)
		pwdConfirmCheck := helpers.IsEmpty(pwdConfirm)

		if uNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
			fmt.Fprintf(w, "ErrorCode is -10:There is empty data. ")
			return
		}
		if pwd == pwdConfirm {
			// --> veritabanına kullanıcıyı kaydet..
			fmt.Fprintf(w, "Registration succesful!")
		} else {
			fmt.Fprintln(w, "password information must be the same")
		}
		/**/
	})

	// Login

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		email = r.FormValue("email")
		pwd = r.FormValue("password")

		emailCheck := helpers.IsEmpty(email)
		pwdCheck := helpers.IsEmpty(pwd)

		if emailCheck || pwdCheck {
			fmt.Fprintf(w, "There is empty data.")
			return
		}
		dbPwd := "12345!*."
		dbEmail := "bb@gmail.com"

		if email == dbEmail && pwd == dbPwd {
			fmt.Fprintln(w, "Login succesful.")
		} else {
			fmt.Fprintln(w, "Login failed")
		}
	})
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Sunucu başlatılamadı:", err)
	}

} // for key, value := range r.Form {
// 	fmt.Printf("%s = %s \n", key, value)
// }
