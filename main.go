package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load(".env")

	portString:= os.Getenv("PORT");
	if portString==""{
		log.Fatal("PORT not found")
	}
	
	router:=chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: false,
		MaxAge: 300,
	}))
	 
	server:=http.Server{
		Handler: router,
		Addr: ":"+portString,
	}

	v1Router:=chi.NewRouter()

	v1Router.Get("/health-check",handlerReadiness)
	v1Router.Get("/error-check",handleErr)

	router.Mount("/v1",v1Router)
	fmt.Printf("%v is the Port \n",portString)
	err:=server.ListenAndServe()
	
	if err!=nil {
		log.Fatal(err)
	}
}