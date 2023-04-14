package main

import (
	"fmt"

	"poo/contas"
)

func Somando(numeros ...int) int {
	resultadoDaSoma := 0
	for _, numero := range numeros {
		resultadoDaSoma += numero
	}
	return resultadoDaSoma
}

func main() {
	var contaDoGuilhermeDois contas.ContaCorrente = contas.ContaCorrente{}
	contaDoGuilhermeDois.Titular = "Guilherme Dois"
	fmt.Println(contaDoGuilhermeDois)

	contaDoGuilherme := contas.ContaCorrente{Titular: "Guilherme",
		NumeroAgencia: 589, NumeroConta: 123456, Saldo: 125.5}

	// Usado quando iremos informar todos os valores para a nossa struct
	// contaDaBruna := contas.ContaCorrente{"Bruna", 222, 111222, 200}

	contaDaBruna := contas.ContaCorrente{Titular: "Bruna", NumeroAgencia: 222, NumeroConta: 111222, Saldo: 200}

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)

	/*
		Nil com um tipo definido

		var s *string = nil
		fmt.Println(s)
	*/

	// Declarando um ponteiro do tipo ContaCorrente
	// var contaDaCris *contas.ContaCorrente
	// contaDaCris = new(contas.ContaCorrente)

	// new Retorna um ponteiro para o tipo que ele recebe por parametro
	// new é uma maneira de obter ponteiros para novos tipos
	contaDaCris := new(contas.ContaCorrente)

	contaDaCris.Titular = "Cris"
	contaDaCris.Saldo = 500.00

	fmt.Println(contaDaCris)  // Devolve o endereço da variavel
	fmt.Println(&contaDaCris) // Devolve o local na memoria onde está a variavel
	fmt.Println(*contaDaCris) // Devolve o conteudo da variavel

	fmt.Println(Somando(1))
	fmt.Println(Somando(1, 1, 2, 4))

	msg := contaDoGuilherme.Sacar(60.00)

	fmt.Println(msg)

	status, valor := contaDoGuilherme.Depositar(2000)

	fmt.Println(status, valor)

	// Como a conta da cris foi definida como ponteiro na sua declaração e usa o new por padrão a variavel sempre retorna o endereço de onde ela está
	success := contaDoGuilherme.Transferir(300, contaDaCris)

	// Nesse caso precisamos incluir o & comercial para passar o endereço de onde está a variavel na memoria
	success2 := contaDoGuilherme.Transferir(300, &contaDaBruna)

	fmt.Println(success, success2)
	fmt.Println(contaDaCris.Saldo, contaDaBruna.Saldo)
}
