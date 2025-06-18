package webserver

import (
	"fmt"
	"net/http"
)

// funciones
func MiWebServer() {
	http.HandleFunc("/", Home)
	fmt.Println("http://localhost:3000/index.html")
	http.ListenAndServe(":3000",nil)
}

func Home(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"../index.html")
}
