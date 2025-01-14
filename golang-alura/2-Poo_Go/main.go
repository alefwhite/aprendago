package main

import (
	"fmt"

	"poo/clientes"
	"poo/contas"
)

func Somando(numeros ...int) int {
	resultadoDaSoma := 0
	for _, numero := range numeros {
		resultadoDaSoma += numero
	}
	return resultadoDaSoma
}

type verificarConta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func main() {
	var contaDoGuilhermeDois contas.ContaCorrente = contas.ContaCorrente{}
	contaDoGuilhermeDois.Titular = clientes.Titular{
		Nome: "Guilherme Dois", CPF: "456", Profissao: "DEV JR",
	}
	fmt.Println(contaDoGuilhermeDois)

	contaDoGuilherme := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Guilherme", CPF: "274", Profissao: "DEV PLENO"},
		NumeroAgencia: 589, NumeroConta: 123456}

	contaDoGuilherme.Depositar(500)

	PagarBoleto(
		&contaDoGuilherme, 499,
	)

	fmt.Println(contaDoGuilherme.ObterSaldo())

	// Usado quando iremos informar todos os valores para a nossa struct
	// contaDaBruna := contas.ContaCorrente{"Bruna", 222, 111222, 200}

	contaDaBruna := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Bruna", CPF: "555", Profissao: "DEVOPS"},
		NumeroAgencia: 222, NumeroConta: 111222,
	}

	contaDaBruna.Depositar(200)

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

	contaDaCris.Titular = clientes.Titular{Nome: "Cris", CPF: "4545", Profissao: "QA"}
	// contaDaCris.saldo = 500.00

	fmt.Println(contaDaCris)  // Devolve o endereço da variavel | Nome do ponteiro a ser utilizado
	fmt.Println(&contaDaCris) // Devolve o local na memoria onde está a variavel | Nesse trecho é feito a referência de memoria da variavel ao ponteiro através do
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
	fmt.Println(contaDaCris.ObterSaldo(), contaDaBruna.ObterSaldo())

	clienteAlef := clientes.Titular{Nome: "Alef", CPF: "456.274.898-20", Profissao: "Dev Pleno"}
	contaDoAlef := contas.ContaCorrente{Titular: clienteAlef, NumeroAgencia: 249, NumeroConta: 95399}
	contaDoAlef.Depositar(4000)
	fmt.Println(contaDoAlef)

	contaDoDenis := contas.ContaPoupanca{Titular: clientes.Titular{Nome: "Denis", CPF: "123", Profissao: "SRE"}, NumeroAgencia: 123, NumeroConta: 321, Operacao: 1}
	fmt.Println(contaDoDenis)
}
