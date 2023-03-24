package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

/*
	go build main.go - Cria o executavel compilado do programa

	./main - Executa o programa

	o comando go run main.go no terminal dentro da pasta que contêm nosso arquivo com o código fonte do programa que o executável
	será automaticamente criado e executado

	Módulo 1 - O que aprendemos
		Instalação do Go
		Go Workspace
			A pasta bin para os arquivos binários, compilados
			A pasta src para o código fonte
			A pasta pkg para os pacotes compartilhados entre as aplicações
		Instalação da extensão do Go no Visual Studio Code
			Com isso temos autocomplete, detecção de erros, etc
		Convenções da linguagem
		Implementação do Hello World
		Compilando e executando um programa em Go

*/

/*
	Declaração de variáveis
		Valor padrão das variáveis sem valor inicial

	Inferência de tipos de variáveis

	Descobrir o tipo da variável
		Através da função TypeOf, do pacote reflect

	Declaração curta de variáveis

	Ler dados digitados do usuário
		Através das funções Scanf e Scan, do pacote fmt

	Mais convenções do Go
		Variáveis e imports não utilizados são deletados

*/

func main() {
	// Tipos de Variaveis
	// var nome string = "Alef"
	// var idade int
	// var versao float32 = 1.1
	// fmt.Println("Olá sr.", nome, "sua idade é", idade)
	// fmt.Println("Este programa está na versão", versao)

	// Inferencia de tipos
	// Por padrão o go sempre irá utilizar o maior tipo daquela variável por exemplo versao_two será float64
	// var versao_two = 1.1
	// fmt.Println("O tipo da variavel nome é", reflect.TypeOf(versao_two))

	// Declaração curta de variáveis
	// sobrenome := "White"
	// fmt.Println("Meu sobrenome é", sobrenome)

	// fmt.Println("1- Iniciar Monitoramento")
	// fmt.Println("2- Exibir Logs")
	// fmt.Println("0- Sair do Programa")

	// var comando int
	//fmt.Scan(&comando) // fmt.Scanf("%d", &comando)

	//fmt.Println("O comando escolhido foi", comando)

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo Logs...")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo do programa...")
	// } else {
	// 	fmt.Println("Não conheço este comando")
	// }

	// switch comando {
	// case 1:
	// 	fmt.Println("Monitorando...")
	// case 2:
	// 	fmt.Println("Exibindo Logs...")
	// case 0:
	// 	fmt.Println("Saindo do programa...")
	// 	os.Exit(0)
	// default:
	// 	fmt.Println("Não conheço este comando")
	// 	os.Exit(-1)
	// }

	//registraLog("SITE", false)

	menu()
}

func menu() {
	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
			iniciarMonitoramentoArrays()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func devolveNomeEIdade() (string, int) {
	nome := "Alef"
	idade := 26

	return nome, idade
}

func exibeIntroducao() {
	nome := "Douglas"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	var sites [4]string // Arrays tem tamanhos fixos sempre devemos informa o seu tamanho
	sites[0] = "https://random-status-code.herokuapp.com/"

	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

func exibeNome() {
	nomes := []string{"Alef", "Théo", "Samara"}           // Slice não precisamos informa seu tamanho
	fmt.Println("O meu slice tem capacidade", cap(nomes)) // Capacidade dele será 3 pois inserimos 3 nomes

	nomes = append(nomes, "White")
	fmt.Println("O meu slice tem capacidade", cap(nomes)) // Por padrão o slice irá dobra a capacidade quando precisarmos inserir um novo valor

	fmt.Println(nomes)
	fmt.Println("O meu slice tem", len(nomes))
}

// restante do código omitido

func iniciarMonitoramentoArrays() {
	fmt.Println("Monitorando...")

	sites := letSitesDoArquivo()
	// sites := []string{"https://random-status-code.herokuapp.com/",
	// 	"https://www.alura.com.br", "https://www.caelum.com.br"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)

			testaSite(site)
		}

		time.Sleep(delay * time.Second)
		// for i := 0; i < len(sites); i++ {
		// 	fmt.Println("\n Testando site", i, ":", sites[i])

		// 	testaSite(sites[i])
		// }
	}

}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func letSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	defer arquivo.Close()

	// arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	// fmt.Println(string(arquivo))

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')

		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		fmt.Println(linha)

		if err == io.EOF {
			// fmt.Println("Ocorreu um erro: ", err)
			break
		}

	}

	fmt.Println("sites", sites)

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(arquivo)

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " " + site + " - Online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}
