package routes

import (
	"fmt"
	"net/http"
	//"models"
	//"github.com/jackc/pgx/v5/pgxpool"
)

/*
func get_user(user models.User, conn *pgxpool.Conn) {
	testUser := models.User{User_id: 3, Email: "danno@balls.com", Username: "dannn0", Password: "passwerd"}

}
*/

func get_user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "got path\n")
}
