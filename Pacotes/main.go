package main

import (
	"fmt"
	"modulo/Pacotes/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Olá, mundo!")

	auxiliar.Escrever()
	

	erro := checkmail.ValidateFormat("ruthfernandes583@gmail.com")
fmt.Println(erro)

}
