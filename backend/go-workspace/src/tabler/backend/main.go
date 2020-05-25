//IT ALWAYS STARTS LIKE THIS
package main

//THE STANDARD GOLANG SHENANIGANS TO IMPORT PACKAGES
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	//----------------------ROOM ROUTES--------------------------------------
	r.HandleFunc("/api/homePage", homePage).Methods("GET")
	r.HandleFunc("/api/searchRooms/{tituloMesa}", searchRooms).Methods("GET")
	r.HandleFunc("/api/createRoom", createRoom).Methods("POST")
	//r.HandleFunc("/api/updateRoom/{id}", updateRoom).Methods("PUT")
	r.HandleFunc("/api/deleteRoom/{id}", deleteRoom).Methods("DELETE")
	//-----------------------------------------------------------------------

	//----------------------USER ROUTES--------------------------------------
	r.HandleFunc("/api/createUser", createUser).Methods("POST")
	//r.HandleFunc("/api/updateProfile/{id}", updateProfile).Methods("DELETE")
	//-----------------------------------------------------------------------

	log.Println("Server Online!")

	log.Fatal(http.ListenAndServe(":8000", r))

}

//----------------------USER FUNCTIONS--------------------------------------
func createUser(w http.ResponseWriter, r *http.Request) {
	stmtIns, err := db.Prepare("INSERT INTO usuario(NOME_USUAR, APELIDO_USUAR, SENHA_USUAR, EMAIL_USUAR, AVATAR_USUAR) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	nomeUsuar := keyVal["NOME_USUAR"]
	apelidoUsuar := keyVal["APELIDO_USUAR"]
	senhaUsuar := keyVal["SENHA_USUAR"]
	emailUsuar := keyVal["EMAIL_USUAR"]
	avatarUsuar := keyVal["AVATAR_USUAR"]

	_, err = stmtIns.Exec(nomeUsuar, apelidoUsuar, senhaUsuar, emailUsuar, avatarUsuar)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Usuario Criado!")

}

func updateProfile(w http.ResponseWriter, r *http.Request) {
}

//--------------------------------------------------------------------------

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
	stmtIns, err := db.Prepare("INSERT INTO mesa(ADM_MESA, TITULO_MESA, DESC_MESA, QTDEJOG_MESA, FORMA_MESA, STATUS_MESA) VALUES (?,?,?,?,?,?) ")
	//stmtIns, err := db.Prepare("INSERT INTO mesa(TITULO_MESA, DESC_MESA) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	admMesa := keyVal["ADM_MESA"]
	tituloMesa := keyVal["TITULO_MESA"]
	descMesa := keyVal["DESC_MESA"]
	qtdejogMesa := keyVal["QTDEJOG_MESA"]
	formaMesa := keyVal["FORMA_MESA"]
	statusMesa := keyVal["STATUS_MESA"]

	_, err = stmtIns.Exec(admMesa, tituloMesa, descMesa, qtdejogMesa, formaMesa, statusMesa)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Mesa criada com sucesso!")
}

//searchRooms FUNCTION
func searchRooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var rooms []Room
	params := mux.Vars(r)

	searchKey := "%" + params["tituloMesa"] + "%"

	result, err := db.Query("SELECT ID_MESA, TITULO_MESA, DESC_MESA FROM mesa WHERE TITULO_MESA LIKE ?", searchKey)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var room Room

	for result.Next() {
		err := result.Scan(&room.ID, &room.Title, &room.Desc)
		if err != nil {
			panic(err.Error())
		}

		rooms = append(rooms, room)
	}

	json.NewEncoder(w).Encode(rooms)
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
	params := mux.Vars(r)

	stmt, err := db.Prepare("DELETE FROM mesa WHERE ID_MESA = ?")

	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Mesa com o ID = %s foi deletada", params["id"])
}

//-------------------------------------------------------
