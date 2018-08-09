package main

import (
	"fmt"
	"os"
	"os/exec"
	"net/http"
	"log"
	"io"
	"encoding/json"
	"math/rand"
	)
	
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
	nomeDaImagem := downloadDaImagem(imagem)
	resultadoDoTexto := executarOCR(nomeDaImagem)
	json.NewEncoder(w).Encode(resultadoDoTexto)
	os.Remove(nomeDaImagem)
}

func executarOCR(nomeDaImagem string) (resultadoDoTexto string) {
	tesseract := fmt.Sprintf("tesseract %s stdout -l por", nomeDaImagem)
	cmd := exec.Command("sh", "-c", tesseract)
	out, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatalf("cmd.Run() failed with %s\n", err)
    }
	resultadoDoTexto = string(out)
	return
}


func downloadDaImagem(url string) (nomeDaImagem string) {
    resposta, erro := http.Get(url)
    if erro != nil {
        log.Fatal(erro)
    }
    defer resposta.Body.Close()
	println(url)
	nomeDaImagem = fmt.Sprintf("imagem%d.png", rand.Intn(10000))
    arquivo, erro := os.Create(nomeDaImagem)
    if erro != nil {
        log.Fatal(erro)
    }
    
    _, erro = io.Copy(arquivo, resposta.Body)
    if erro != nil {
        log.Fatal(erro)
	}
	
	arquivo.Close()
	return
}
