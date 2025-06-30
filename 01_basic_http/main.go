package main

/*
	Etapa 1:
	Objetivo: Criar uma API mínima que responde apenas à rota raiz /.
	O servidor escuta na porta 8080.
	O handler responde com uma simples mensagem de texto.
*/

import (
	"fmt"
	"net/http" // Pacote padrão para criar servidores HTTP em Go
)

func main() {
	// Registramos a função `home` como handler da rota "/"
	http.HandleFunc("/", home)

	// Iniciamos o servidor na porta 8080.
	// O segundo argumento `nil` significa que vamos usar o DefaultServeMux
	http.ListenAndServe(":8080", nil)
}

// Função que lida com a requisição da rota "/"
// `w` permite escrever a resposta, e `r` contém informações da requisição
func home(w http.ResponseWriter, r *http.Request) {
	// Escreve uma mensagem de boas-vindas no corpo da resposta HTTP
	fmt.Fprintf(w, "Welcome to the home page!")
}
