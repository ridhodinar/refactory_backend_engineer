package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"log"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	googleClientId = "763372287989-psn6mk9mj16su28jsb3qce0ii4g6ehg2.apps.googleusercontent.com"
	googleClientSecret = "CmTFaixJP0bHlZI7ZFbuqRmg"
	googleCallback = "http://localhost:3000/auth/google/callback"
)

var db *gorm.DB

type User struct {
  gorm.Model
  Email  		string
	Name 			string
  Provider 	string
	AvatarURL 	string
}

func dbInit(){
	dsn := "root:@tcp(127.0.0.1:3306)/golang-oauth?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
  db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
  db.AutoMigrate(&User{})
}

func home(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "type http://localhost:3000/auth/google to login with google")
}

func authHandler(res http.ResponseWriter, req *http.Request) {
	gothic.BeginAuthHandler(res, req)
}

func googleCallbackHandler(res http.ResponseWriter, req *http.Request) {

	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Fprintln(res, err)
		
		return
	}
	//fmt.Fprintln(res, user)
	createUser := User{Email: user.Email, Name: user.Name, Provider: user.Provider, AvatarURL: user.AvatarURL}
	db.Create(&createUser)
	fmt.Fprintf(res, "http://localhost:3000/user/%d", createUser.ID)
}

func getUser(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId := params["id"]
	var user User
	db.First(&user, userId)

	userJson, err := json.Marshal(user)

	if err != nil {
		fmt.Println("Unable to encode JSON")
	}

	//fmt.Println(string(userJson))

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(userJson)
}

func main() {
	dbInit()
  
  key := "Secret-session-key"  // Replace with your SESSION_SECRET or similar
  maxAge := 86400 * 30  // 30 days
  isProd := false       // Set to true when serving over https

  store := sessions.NewCookieStore([]byte(key))
  store.MaxAge(maxAge)
  store.Options.Path = "/"
  store.Options.HttpOnly = true   // HttpOnly should always be enabled
  store.Options.Secure = isProd

  gothic.Store = store

  goth.UseProviders(
    google.New(googleClientId, googleClientSecret, googleCallback, "email", "profile"),
  )

  r := mux.NewRouter()
  r.HandleFunc("/auth/{provider}/callback", googleCallbackHandler)
	r.HandleFunc("/", home)
	r.HandleFunc("/auth/{provider}", authHandler)
	r.HandleFunc("/user/{id}", getUser)
  
  log.Println("listening on localhost:3000")
  log.Fatal(http.ListenAndServe(":3000", r))
}