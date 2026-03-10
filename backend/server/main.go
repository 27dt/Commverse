package main

// ----------------------------------------------------| IMPORTS |---------------------------------------------------- //
import (
	// To pass/receive request bodies on routes
	"context"  // For sql requests
	"fmt"      // For printing
	"log"      // For logging errors
	"net/http" // For routes

	// Access db models
	// Access route functions
	"github.com/jackc/pgx/v5/pgxpool" // pgx for PostgreSQL db driver
)

func main() {
	// ----------------------------------------------| DATABASE CONNECTION |---------------------------------------------- //

	// Sets up connection string equal to database url: postgres://<username>:<password>@localhost:<port>/<db_name>
	connstr := "postgres://postgres:1mS00oooo0tir3d@localhost:5432/commversedb"

	// Initializes a database connection pool
	/* pgx.pool.New() either creates a pointer to a pgxpool.Pool, or an error
	- pool := pgxpool.Pool if the function call is successful
	- err := error if the function call is unsuccessful. null otherwise
	  (The "error" type is an interface type and represents a string) */
	pool, err := pgxpool.New(context.Background(), connstr)

	// In the event err != nil, an error occurs. Display log message and print error
	if err != nil {
		log.Fatal("Error: unable to connect to database:", err)
	}

	// If the error is specifically from pinging
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal("Error: unable to ping database: ", err)
	}

	// If successful, print message
	fmt.Println("Success: connected to PostgreSQL database")

	// -----------------------------------------------------| ROUTES |---------------------------------------------------- //

	// Creates an HTTP request multiplexer
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", routes.get_user)

	// Creates listening server for routes on localhost:9090
	http.ListenAndServe("localhost:9090", mux)
}
