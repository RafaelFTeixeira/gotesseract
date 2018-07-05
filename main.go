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
	"encoding/json"
	)
	
const NomeDoArquivoDaImagem = "imagem.png"
const NomeDoArquivoDeTexto = "texto.txt"

func main() {
	http.HandleFunc("/", obterImagemDaUrl)
    http.ListenAndServe(":3000", nil)
}

func obterImagemDaUrl(w http.ResponseWriter, r *http.Request) {
	imagens, ok := r.URL.Query()["image"]
    
    if !ok || len(imagens[0]) < 1 {
        json.NewEncoder(w).Encode("Informe o link da imagem. Examplo: http://localhost:3000/?image=hello.png")
        return
	}
	
	imagem := string(imagens[0])
	downloadDaImagem(imagem)
	executarOCR()
	texto := obterTexto()
	json.NewEncoder(w).Encode(texto)
	removerArquivos()
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
