package controllers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//User Struct
type User struct {
	Apelido string `json:"apelido"`
}

//CheckLogin exported
func CheckLogin(w http.ResponseWriter, r *http.Request) {
	var userLogged User

	db, err = sql.Open("mysql", "user_tester:123456@tcp(127.0.0.1:3000)/tabler_db")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	apelidoUsuar := keyVal["APELIDO_USUAR"]
	senhaUsuar := keyVal["SENHA_USUAR"]

	result, err := db.Query("SELECT APELIDO_USUAR FROM usuario WHERE APELIDO_USUAR = ? AND SENHA_USUAR = ?", apelidoUsuar, senhaUsuar)
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		err := result.Scan(userLogged.Apelido)
		if err != nil {
			panic(err.Error())
		}

	}
	json.NewEncoder(w).Encode(userLogged)
	w.WriteHeader(http.StatusOK)

}
