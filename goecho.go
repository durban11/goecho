package main

import (
	//"database/sql"
	//"fmt"
	//_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("c:\\gohttp"))))
}
