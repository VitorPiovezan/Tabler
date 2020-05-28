//IT ALWAYS STARTS LIKE THIS
package main

//THE STANDARD GOLANG SHENANIGANS TO IMPORT PACKAGES
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http" //TO HANDLE HTTP REQUESTS
	"os"

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
	r.HandleFunc("/api/joinRoom", joinRoom).Methods("POST")
	r.HandleFunc("/api/updateRoom", updateRoom).Methods("PUT")
	r.HandleFunc("/api/deleteRoom", deleteRoom).Methods("DELETE")
	//-----------------------------------------------------------------------

	//----------------------USER ROUTES--------------------------------------
	r.HandleFunc("/api/createUser", createUser).Methods("POST")
	r.HandleFunc("/api/updateProfile", updateProfile).Methods("PUT")
	r.HandleFunc("/api/uploadAvatar", uploadAvatar).Methods("POST")
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
	w.WriteHeader(http.StatusOK)

}

func updateProfile(w http.ResponseWriter, r *http.Request) {

	stmtIns, err := db.Prepare("UPDATE usuario SET NOME_USUAR = ? , APELIDO_USUAR = ? , SENHA_USUAR = ? , EMAIL_USUAR = ? , AVATAR_USUAR = ? WHERE ID_USUAR = ?")
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
	idUsuar := keyVal["ID_USUAR"]

	_, err = stmtIns.Exec(nomeUsuar, apelidoUsuar, senhaUsuar, emailUsuar, avatarUsuar, idUsuar)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Dados do usuario alterados com sucesso!")
	w.WriteHeader(http.StatusOK)
}

func uploadAvatar(w http.ResponseWriter, r *http.Request) {

	var picPath string
	var userName string

	userName = "Offar"

	picPath = "ProfilePics/" + userName

	//---------------------CREATING FOLDER--------------------
	//Sets folder path and folder name
	_, err := os.Stat(picPath)

	//Checks if that folder path already exists
	//If the nested folder exists, do nothing
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(picPath, os.ModePerm)
		stmtUpdt, err := db.Prepare("UPDATE usuario SET AVATAR_USUAR = ? WHERE APELIDO_USUAR = ?")
		_, err = stmtUpdt.Exec(picPath, userName)
		if err != nil {
			panic(err.Error())
		}
		if errDir != nil {
			log.Fatal(err)
		}
	}
	//--------------------------------------------------------

	//----------------------SAVING FILE----------------------
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile(picPath, userName+"-*.png")

	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)
	w.WriteHeader(http.StatusOK)
	//------------------------------------------------------------------
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
	w.WriteHeader(http.StatusOK)
}

//--------------ROOM FUNCTIONS---------------------------

//createRoom FUNCTION
func createRoom(w http.ResponseWriter, r *http.Request) {

	stmtIns, err := db.Prepare("INSERT INTO mesa(ADM_MESA, TITULO_MESA, DESC_MESA, QTDEJOG_MESA, FORMA_MESA, STATUS_MESA) VALUES (?,?,?,?,?,?) ")

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
	w.WriteHeader(http.StatusOK)
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
	w.WriteHeader(http.StatusOK)
}

//joinRoom FUNCTION
func joinRoom(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	idMesa := keyVal["ID_MESA"]
	idUsuar := keyVal["ID_USUAR"]
	mestreJoga := keyVal["MESTRE_JOGA"]
	charName := keyVal["NOMECHAR_JOGA"]
	userName := keyVal["UserName"]

	sheetPath := "CharSheet/" + userName + "/" + charName

	//---------------------CREATING FOLDER--------------------
	_, errFolder := os.Stat(sheetPath)

	if os.IsNotExist(errFolder) {
		errDir := os.MkdirAll(sheetPath, os.ModePerm)
		if errDir != nil {
			log.Fatal(errFolder)
		}
	}
	//--------------------------------------------------------

	//----------------------SAVING FILE----------------------
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile(sheetPath, charName+"-*.pdf")

	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
	//------------------------------------------------------------------

	stmtIns, err := db.Prepare("INSERT INTO mesa_jogadores(ID_MESA, ID_USUAR, MESTRE_JOGA, FICHA_JOGA) VALUES (?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	//CHECK IF THERE'S ALREADY A DM AT THE TABLE
	var isThereDm int
	_ = db.QueryRow("SELECT COUNT(*) FROM mesa_jogadores WHERE ID_MESA = ? AND MESTRE_JOGA = 1", idMesa).Scan(&isThereDm)

	if err != nil {

		panic(err.Error())

	}

	if isThereDm != 0 && mestreJoga == "1" {

		fmt.Fprintf(w, "Já existe mestre nesta mesa!")

	} else { //IF THERE'S NO DM, INSERT THE PLAYER IN THE ROOM

		_, err = stmtIns.Exec(idMesa, idUsuar, mestreJoga, sheetPath)
		if err != nil {
			panic(err.Error())
		}

		fmt.Fprintf(w, "Jogador inserido na mesa!")
		w.WriteHeader(http.StatusOK)
	}

}

//updateRoom FUNCTION
func updateRoom(w http.ResponseWriter, r *http.Request) {

	stmtIns, err := db.Prepare("UPDATE mesa SET TITULO_MESA = ? , DESC_MESA = ? , QTDEJOG_MESA = ? , FORMA_MESA = ? WHERE ID_MESA = ?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	idMesa := keyVal["ID_MESA"]
	tituloMesa := keyVal["TITULO_MESA"]
	descMesa := keyVal["DESC_MESA"]
	qtdejogMesa := keyVal["QTDEJOG_MESA"]
	formaMesa := keyVal["FORMA_MESA"]

	_, err = stmtIns.Exec(tituloMesa, descMesa, qtdejogMesa, formaMesa, idMesa)
	if err != nil {
		panic(err.Error())
	}

	w.WriteHeader(http.StatusOK)
}

//deleteRoom FUNCTION
func deleteRoom(w http.ResponseWriter, r *http.Request) {

	stmt, err := db.Prepare("DELETE FROM mesa WHERE ID_MESA = ?")

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	idMesa := keyVal["ID_MESA"]

	_, err = stmt.Exec(idMesa)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Mesa com o ID = %s foi deletada", idMesa)
	w.WriteHeader(http.StatusOK)
}

//-------------------------------------------------------
