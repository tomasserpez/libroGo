// main.go
package main

import (
	"fmt"
	"os"
	"os/user"
	"tomprueba/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Bienvenido %s! Este es un lenguaje que estoy programando con fines didacticos para aprender el funcionamiento interno de un interprete.", user.Username)
	fmt.Printf("\nEjecute sus comandos:%s")
	repl.Start(os.Stdin, os.Stdout)
}
