package main

import (
	"fmt"
	"net/http"
    "github.com/bayrambartu/GoRestFulApi-LoginAndRegister/helpers"
)

func main() {

	uName,email,pwd,pwdConfirm := "","","","",

	mux := http.NewServeMux()
	// SignUp
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		uName = r.FormValue("username")
		email = r.FormValue("email")
		pwd = r.FormValue("pwd")
		pwdconfirm = r.FormValue("confirm")	

		uNameCheck := helper.IsEmpty(uName)
		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pwd)
		pwdConfirmCheck := helper.IsEmpty(pwdConfirm)

		if uNameCheck 	|| emailCheck || pwdCheck ||pwdConfirmCheck{
			fmt.Fprint(w,"ErrorCode is -10:There is empty data. ")
			return
		}
		if pwd == pwdConfirm{
			// --> veritabanına kullanıcıyı kaydet.. 
			fmt.Fprintf(w,"Registration succesful!")
		}else{
			fmt.Fprintln(w,"password information must be the same")
		}

		


		
	})

	// Login

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

	})
	http.ListenAndServe(":8080", mux)

}

// for key, value := range r.Form {
		// 	fmt.Printf("%s = %s \n", key, value)
		// }
