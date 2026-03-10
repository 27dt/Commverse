package queries

import (
	"context"
	"log"

	"models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func create_user(user models.User, conn *pgxpool.Conn) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO users (email, username, password) VALUES ($1, $2, $3)", user.Email, user.Username, user.Password)
	if err != nil {
		log.Fatalf("Error inserting user: %v", err)
	}
	return nil
}

/*func read_user(user_id int) {

}

func update_user(user models.User) {

}

func delete_user(user_id int) {

}*/
