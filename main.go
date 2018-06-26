package main

import (
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
	"net/http"
	"log"
	"io"
	)

func executar() {
		tesseract := fmt.Sprintf("tesseract imagem.png texto -l eng")
    exec.Command("sh", "-c", tesseract).Output()
}

func obterTexto() (texto string) {
	binario, erro := ioutil.ReadFile("texto.txt")
	if erro != nil {
        fmt.Print(erro)
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

    
    arquivo, erro := os.Create("imagem.png")
    if erro != nil {
        log.Fatal(erro)
    }
    
    _, erro = io.Copy(arquivo, resposta.Body)
    if erro != nil {
        log.Fatal(erro)
	}
	
    arquivo.Close()
}
	
func main() {
	imagem := os.Args[1]
	downloadDaImagem(imagem)
	executar()
	texto := obterTexto()
	fmt.Println(texto)
}