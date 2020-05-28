package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var db *sql.DB
var err error

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
