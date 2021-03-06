//Optimizing HTTP server responses with GZIP compression

package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

const(
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func sayHello(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "Hello World!")
}

func main(){
	
	mux := http.NewServeMux()
	mux.HandleFunc("/", sayHello)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, handlers.CompressHandler(mux))
	if err != nil{
		log.Fatal("Erro ao iniciar http server : ", err)
		return
	}
}