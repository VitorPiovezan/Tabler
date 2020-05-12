//IT ALWAYS STARTS LIKE THIS
package main

//THE STANDARD GOLANG SHENANIGANS TO IMPORT PACKAGES
import (
	"database/sql"
	"log"      //TO USE THE log.Fatal FUNCTION
	"net/http" //TO HANDLE HTTP REQUEST

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux" //TO CREATE ROUTES. PEOPLE TEND TO USE THIS PACKAGE
)

//User teste
type User struct {
	Name string
}

//getBook FUNCTION
//EVERY FUNCTION THAT IS A ROUT HANDLER NEED TO HAVES THESE PROPERTIES. SO, BASICALY JUST COPY PASTE IT WHEN NEEDED.
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func dbConnect() (db *sql.DB) {

	db, err := sql.Open("mysql", "user_tester:123456@tcp(127.0.0.1:3000)/tabler")

	if err != nil {
		panic(err.Error())
	}

	return db
}

//LIKE IN C, IT ALWAYS NEED A MAIN FUNCTION TO START THINGS OFF
func main() {

	insert, err := db.Query("INSERT INTO cadastro_usuario(NOME_USUAR) VALUES ('Rafael')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	log.Println("Connected to database!")

	//MUX IS USED TO CREATE THE BACKEND ROUTES
	r := mux.NewRouter()

	//CREATING THE ROUTES
	//THE FIRST PARAMETER IS THE NAME OF THE ROUT "/something/whatever", THE SECOND, IS THE FUNCTION IT CALLS.
	//THEN WE USE .Methods(GET OR POST OR DELETE OR WHATEVER) TO SPECIFY THE METHOD WE ARE USING IN THIS ROUT
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Println("Server Online!")
	//SETS THE SERVER UP AT PORT 3035. YOU CAN USE WHATEVER PORT YOU WANT
	//THE log.Fatal FUNCTION, IS USED JUST TO LOG THE ERRORS IF THEY HAPPEN
	log.Fatal(http.ListenAndServe(":3035", r))

}
