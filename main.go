package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/JamiuJimoh/hngorg/api"
	"github.com/JamiuJimoh/hngorg/db/sqlc"
	"github.com/JamiuJimoh/hngorg/utils"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("unable to connect to database: %v\n", err)
	}
	defer conn.Close(ctx)

	db := sqlc.New(conn)
	secretKey := os.Getenv("TOKEN_SYMMETRIC_KEY")
	tokenCfg, err := utils.NewAccessToken([]byte(secretKey))
	apiCfg, err := api.NewApiConfig(db, tokenCfg)
	if err != nil {
		log.Fatalf("error while setting up token config")
	}
	server := &http.Server{
		Addr: ":" + port,
	}

	http.HandleFunc("POST /auth/register", apiCfg.Register)
	http.HandleFunc("POST /auth/login", apiCfg.Login)

	http.HandleFunc("GET /api/users/{id}", apiCfg.AuthMiddleware(apiCfg.GetUserInSameOrg))
	http.HandleFunc("GET /api/organisations", apiCfg.AuthMiddleware(apiCfg.GetOrganistions))
	http.HandleFunc("GET /api/organisations/{orgId}", apiCfg.AuthMiddleware(apiCfg.GetOrganistion))
	http.HandleFunc("POST /api/organisations", apiCfg.AuthMiddleware(apiCfg.CreateOrganistion))
	http.HandleFunc("POST /api/organisations/{orgId}/users", apiCfg.PatchOrganistionWithUser)

	log.Fatal(server.ListenAndServe())
}
