package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitonamentos = 3
const delay = 4

func main() {

	exibeIntroducao()
	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs")
		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	nome := "Fylip"
	versao := 1.1
	inscrito := true

	fmt.Println("Olá, sr.", nome)
	if inscrito {
		fmt.Println("Sou inscrito")
	}
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir os Logs")
	fmt.Println("0 - Sair do Programa")
	fmt.Println("")
}

func iniciarMonitoramento() {
	fmt.Println("Tela de Monitoramento")
	sites := lerSitesDoArquivo()

	for i := 0; i < monitonamentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func lerSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	return sites
}

func testaSite(site string) {

	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", response.StatusCode)
	}

}
func exibeNomes() {
	nomes := []string{"Fylip", "Luiz", "Aline", "Ana Clara"}
	fmt.Println(nomes)
}
