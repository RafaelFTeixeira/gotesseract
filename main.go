package main

import (
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
	"net/http"
	"log"
	"io"
	"strings"
	)
	
const NomeDoArquivoDaImagem = "imagem.png"
const NomeDoArquivoDeTexto = "texto.txt"

func main() {
	if (len(os.Args) > 1) {
		imagem := os.Args[1]
		downloadDaImagem(imagem)
		executarOCR()
		texto := obterTexto()
		fmt.Println(texto)
		removerArquivos()
	} else {
		fmt.Println("Informe o link da imagem") 
	}
}

func executarOCR() {
	nomeDoArquivoDeTexto := strings.Split(NomeDoArquivoDeTexto, ".")[0]
	tesseract := fmt.Sprintf("tesseract %s %s -l por", NomeDoArquivoDaImagem, nomeDoArquivoDeTexto)
    exec.Command("sh", "-c", tesseract).Output()
}

func obterTexto() (texto string) {
	binario, erro := ioutil.ReadFile(NomeDoArquivoDeTexto)
	if erro != nil {
        log.Fatal(erro)
    }
	texto = string(binario)
	return
}

func downloadDaImagem(url string)  {
    resposta, erro := http.Get(url)
    if erro != nil {
        log.Fatal(erro)
    }
    defer resposta.Body.Close()
    
    arquivo, erro := os.Create(NomeDoArquivoDaImagem)
    if erro != nil {
        log.Fatal(erro)
    }
    
    _, erro = io.Copy(arquivo, resposta.Body)
    if erro != nil {
        log.Fatal(erro)
	}
	
    arquivo.Close()
}


func removerArquivos() {
	os.Remove(NomeDoArquivoDaImagem)
	os.Remove(NomeDoArquivoDeTexto)
}
