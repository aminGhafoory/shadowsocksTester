package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aminghafoory/shadowTester/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func init() {

}
func main() {

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is found in the .env file")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL not found in .env file: ")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("can't connect to Database: ", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	go startScraping(apiCfg.DB, 3, time.Second*15)
	go startTesting(apiCfg.DB, 50, time.Minute*1)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/sub", apiCfg.addSub)
	r.Get("/sub", apiCfg.showSubs)
	r.Get("/ss", apiCfg.ShowAllSSs)
	r.Get("/best", apiCfg.BestConfigs)
	http.ListenAndServe(":"+portString, r)

}
