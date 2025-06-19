package main

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"net/http"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"database/sql"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var jwtKey = []byte("your_secret_key")
var db *sql.DB

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	var id int
	var hashedPwd string
	err := db.QueryRow("SELECT id, password_hash FROM users WHERE username=$1", req.Username).Scan(&id, &hashedPwd)
	if err != nil {
		http.Error(w, "Invalid user", 401)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid password", 401)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString(jwtKey)

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	_, err := db.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", req.Username, string(hashedPwd))
	if err != nil {
		http.Error(w, "Username taken", 400)
		return
	}

	w.WriteHeader(201)
}
