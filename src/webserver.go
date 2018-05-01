package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func createNewRouter() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	staticFileDir := http.Dir("/home/abc/go/src/github.com/webserver/assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDir))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")	
	
	return r
	
}

func main(){
	r := createNewRouter()
//	r.HandleFunc("/hello",handler).Methods("GET")
//	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", r)
	fmt.Println("Started server")
	connString:="root:root@/bird_encyclopedia"
	db, err := sql.Open("mysql", connString)
	if err!=nil{
		panic(err)
	}
	InitStore(&dbStore{db: db})
	err := db.Ping()
	if err!=nil{
		fmt.Println("error in pinging db")
	}
	defer db.Close()
	
	
}

func handler(w http.ResponseWriter,req *http.Request){
	fmt.Fprint(w, "hello world")
}