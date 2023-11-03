package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 10

func main() {
	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço seu comando!")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	nome := "Luan Diniz"
	versao := 1.1

	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido) // & indica o ponteiro da minha variavel
	fmt.Println("O camando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{
		"https://random-status-code.herokuapp.com/",
		"https://alura.com.br/",
		"https://poupacred.com.br/",
		"https://google.com/",
	}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
