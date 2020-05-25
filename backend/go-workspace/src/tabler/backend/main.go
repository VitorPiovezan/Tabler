//IT ALWAYS STARTS LIKE THIS
package main

//THE STANDARD GOLANG SHENANIGANS TO IMPORT PACKAGES
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http" //TO HANDLE HTTP REQUEST

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//User Struct
type User struct {
	Name string `json:"name"`
}

//Room Struct
type Room struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

var db *sql.DB
var err error

//LIKE IN C, IT ALWAYS NEED A MAIN FUNCTION TO START THINGS OFF
func main() {

	db, err = sql.Open("mysql", "user_tester:123456@tcp(127.0.0.1:3000)/tabler_db")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Connected to DB!")

	//MUX IS USED TO CREATE THE BACKEND ROUTES
	r := mux.NewRouter()

	//CREATING THE ROUTES
	//THE FIRST PARAMETER IS THE NAME OF THE ROUT "/something/whatever", THE SECOND, IS THE FUNCTION IT CALLS.
	//THEN WE USE .Methods(GET OR POST OR DELETE OR WHATEVER) TO SPECIFY THE METHOD WE ARE USING IN THIS ROUT
	r.HandleFunc("/api/homePage", homePage).Methods("GET")
	/*r.HandleFunc("/api/searchRooms/{info}", searchRooms).Methods("GET")
	r.HandleFunc("/api/createRoom", createRoom).Methods("POST")
	r.HandleFunc("/api/updateRoom/{id}", updateRoom).Methods("PUT")
	r.HandleFunc("/api/deleteRoom/{id}", deleteRoom).Methods("DELETE")*/

	log.Println("Server Online!")
	//SETS THE SERVER UP AT PORT 3035. YOU CAN USE WHATEVER PORT YOU WANT
	//THE log.Fatal FUNCTION, IS USED JUST TO LOG THE ERRORS IF THEY HAPPEN
	log.Fatal(http.ListenAndServe(":8000", r))

	//NEED TO CREATE SEPARATE FUNCTION TO INSERT (DUNNO HOW :^[ )
	/*insert, err := db.Query("INSERT INTO usuario(NOME_USUAR, APELIDO_USUAR, EMAIL_USUAR, AVATAR_USUAR) VALUES ('Rafael Carvalho', 'Offar', 'email@email.com','C:/Users/Offar/Pictures/osmar')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()*/

}

//EVERY FUNCTION THAT IS A ROUT HANDLER NEED TO HAVES THESE PROPERTIES. SO, BASICALY JUST COPY PASTE IT WHEN NEEDED.
//homePage FUNCTION
func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var rooms []Room

	result, err := db.Query("SELECT ID_MESA, TITULO_MESA, DESC_MESA FROM mesa")

	if err != nil {

		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var room Room

		err := result.Scan(&room.ID, &room.Title, &room.Desc)
		if err != nil {
			panic(err.Error())
		}
		rooms = append(rooms, room)
	}

	json.NewEncoder(w).Encode(rooms)
}

//--------------ROOM FUNCTIONS---------------------------

//createRoom FUNCTION
func createRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

//searchRooms FUNCTION

func searchRooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

//joinRoom FUNCTION
func joinRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

//updateRoom FUNCTION
func updateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

//deleteRoom FUNCTION
func deleteRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

//-------------------------------------------------------
