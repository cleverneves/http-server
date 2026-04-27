package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go main.go <diretorio> <porta>")
		os.Exit(1)
	}
	httpDir := os.Args[1]
	port := os.Args[2]

	filesystem := http.FileServer(http.Dir(httpDir))

	fmt.Printf("Executando o servidor na porta %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, filesystem))
}
