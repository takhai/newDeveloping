package main

import (
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      //user ID user
var dbSessions = map[string]string{} //session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	bs, _ := bcrypt.GenerateFromPassword([]byte("mypass"), bcrypt.MinCost)
	dbUsers["test@test.com"] = user{"test@test.com", bs, "Tak", "Mo"}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favi.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var u user

	//process form submission
	if req.Method == http.MethodPost {

		//get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")

		//username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "username taken", http.StatusForbidden)
			return
		}

		//create session
		sID := uuid.Must(uuid.NewV4())
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
			//Secure: true,
			HttpOnly: true,
		}

		http.SetCookie(w, c)
		dbSessions[c.Value] = un

		//store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		u := user{un, bs, f, l, r}
		dbUsers[un] = u

		//redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", u)

}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	//var u user

	//process form submission
	if req.Method == http.MethodPost {
		//get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		//is there username
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "username or password does not match", http.StatusForbidden)
			return
		}
		//does entered password match the stored password
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "username or password does not match", http.StatusForbidden)
			return
		}

		//create session
		sID := uuid.Must(uuid.NewV4())
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
			//Secure: true,
			HttpOnly: true,
		}

		http.SetCookie(w, c)
		dbSessions[c.Value] = un
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.gohtml", nil)

}

func logout(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	//delete session
	delete(dbSessions, c.Value)
	//remove cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
		//Secure: true,
		HttpOnly: true,
	}

	http.SetCookie(w, c)

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
