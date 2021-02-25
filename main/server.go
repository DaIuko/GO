package main

import(
	"database/sql"
	"log"
	"net/http"
	"text/template"
	_ "github.com/go-sql-driver/mysql"
)

type Employee struct{
	Id int
	Name string
	City string

}
// connect sql Database
// using username : root / password : dbpass / Name : dbName
func dbconn()(db *sql.DB){
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "naveen"
	dbName := "employeedb"
	db, err := sql.Open(dbDriver, dbUser +":"+ dbPass +"@/"+dbName)

// error Handling for database conecting issues
	if err != nil{
		log.Println("Database Not connected ?")
		panic(err.Error())
	}
	log.Println("Database connect Success")
	return db
}

//Get template file from 'form' folder
var tmpl = template.Must(template.ParseGlob("D:/Company/GitLab project/crud_operation_golang/src/main/template/*"))

//Display data homepage.
//Display data ascending order.
func Index(w http.ResponseWriter, r *http.Request) {
	db := dbconn()
	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id ASC")   // Display Data Ascending order
	if err != nil {
		log.Println("Can't find , Error? ")
		panic(err.Error())
	}
	emp := Employee{}
	res := []Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
		res = append(res, emp)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}
// View Data from the database
// View selected data using ID
func Show(w http.ResponseWriter, r *http.Request) {
	db := dbconn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
	if err != nil {
		log.Println("Can't find , Error? ")
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
	}
	tmpl.ExecuteTemplate(w, "Show", emp)
	defer db.Close()
}

//create new template
// for enter the new data set.
func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

// Edit enterd data
// Employee Data edit by using ID
// can edit selected data
func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbconn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
	if err != nil {
		log.Println("Can't find , Error? ")
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
	}
	tmpl.ExecuteTemplate(w, "Edit", emp)
	defer db.Close()
}

// Insert new Employee Data
// Store data into Database.
func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbconn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		insForm, err := db.Prepare("INSERT INTO Employee(name, city) VALUES(?,?)")
		if err != nil {
			log.Println("Can't find , Error? ")
			panic(err.Error())
		}
		insForm.Exec(name, city)
		log.Println("INSERT: Name: " + name + " | City: " + city)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbconn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE Employee SET name=?, city=? WHERE id=?")
		if err != nil {
			log.Println("Can't find , Error? ")
			panic(err.Error())
		}
		insForm.Exec(name, city, id)
		log.Println("UPDATE: Name: " + name + " | City: " + city)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

// Delete Data from the database
func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbconn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
	if err != nil {
		log.Println("Can't find , Error? ")
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

// main controller
// handle all above crud function
// handle frontend request
func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}

