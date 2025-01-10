package main

import (
	"demo/handlers"
	"flag"
	"log"
	"net/http"
	"os"
	"runtime"
)

var PORT string

//var DB_CONN string

func init() {
	// args := os.Args
	// fmt.Println(args)
	// --port=8084 -port= 8081
	PORT = os.Getenv("PORT")
	if PORT == "" {
		flag.StringVar(&PORT, "port", "8083", "--port=8083 or -port=8083 or --port 8083")
	}
	//flag.StringVar(&DB_CONN, "db", "some db", "something")
	flag.Parse()
}

func main() {

	log.Println("application is running on port", PORT)
	http.HandleFunc("/", handlers.Root)

	// http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("pong"))
	// 	w.WriteHeader(http.StatusOK)
	// })

	http.HandleFunc("/ping", handlers.Ping)

	http.HandleFunc("/health", handlers.Health)

	http.HandleFunc("/user", handlers.CreateUser)

	if err := http.ListenAndServe(":"+PORT, nil); err != nil { // 0.0.0.0:8083{
		log.Println(err.Error())
		runtime.Goexit() // what it does?
	}
	log.Println("End of main")
}
