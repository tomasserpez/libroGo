// main.go
package main

import (
	"fmt"
	"os"
	"os/user"
	"tomlang/repl"
)

func main() {
	usuario, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Bienvenido %s! Este es un lenguaje que estoy programando con fines didacticos para aprender el funcionamiento interno de un interprete.", usuario.Username)
	fmt.Printf("\nEjecute sus comandos:\n")
	repl.Start(os.Stdin, os.Stdout)
}
