import (
	"golang.org/x/crypto/bcrypt"
)

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
