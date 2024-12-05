package main

import (
	"database/sql"
	"go-htmx-tailwind-app/internal/handlers"
	"html/template"
	"log"
	"mime"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var tmpl *template.Template
var db *sql.DB

var Store = sessions.NewCookieStore([]byte("usermanagementsecret"))

type Task struct {
	Id   int
	Task string
	Done bool
}

func init() {
	tmpl, _ = template.ParseGlob("web/templates/*.html")

		//Set up Sessions
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 3,
		HttpOnly: true,
	}
}

func initDB() {
	var err error
	// Initialize the db variable
	db, err = sql.Open("mysql", "root:CxjTikmIOw@(mysql.default.svc.cluster.local:3306)/usermanagement?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	// Check the database connection
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func main() {
  // Set MIME type for CSS files
  mime.AddExtensionType(".css", "text/css")

  gRouter := mux.NewRouter()

	//Setup MySQL
	initDB()
	defer db.Close()

  // Serve static files
  fs := http.FileServer(http.Dir("web/static"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))

	fileServer := http.FileServer(http.Dir("uploads"))
	gRouter.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads", fileServer))

	gRouter.HandleFunc("/", handlers.Homepage(db, tmpl, Store))

	gRouter.HandleFunc("/register", handlers.RegisterPage(db, tmpl)).Methods("GET")

	gRouter.HandleFunc("/register", handlers.RegisterHandler(db, tmpl)).Methods("POST")

	gRouter.HandleFunc("/login", handlers.LoginPage(db, tmpl)).Methods("GET")

	gRouter.HandleFunc("/login", handlers.LoginHandler(db, tmpl, Store)).Methods("POST")

	gRouter.HandleFunc("/edit", handlers.Editpage(db, tmpl, Store)).Methods("GET")

	gRouter.HandleFunc("/edit", handlers.UpdateProfileHandler(db, tmpl, Store)).Methods("POST")

	gRouter.HandleFunc("/upload-avatar", handlers.AvatarPage(db, tmpl, Store)).Methods("GET")

	gRouter.HandleFunc("/upload-avatar", handlers.UploadAvatarHandler(db, tmpl, Store)).Methods("POST")

	gRouter.HandleFunc("/logout", handlers.LogoutHandler(Store)).Methods("GET")

	http.ListenAndServe(":8000", gRouter)
}