package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth "github.com/abbot/go-http-auth"
)

func Secret(user, realm string) string {
	if user == "Cleverton" {
		stringValue := "ADMIN"
		return stringValue
	}
	return ""
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go main.go <diretorio> <porta>")
		os.Exit(1)
	}
	httpDir := os.Args[1]
	port := os.Args[2]

	authenticator := auth.NewBasicAuthenticator("", Secret)
	http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, ar *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &ar.Request)
	}))

	fmt.Printf("Executando o servidor na porta %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
