package main

import (
	"fmt"
	"os"
	"os/exec"
	"net/http"
	"log"
	"io"
	"encoding/json"
	)
	
const NomeDoArquivoDaImagem = "imagem.png"

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
	resultadoDoTexto := executarOCR()
	json.NewEncoder(w).Encode(resultadoDoTexto)
	os.Remove(NomeDoArquivoDaImagem)
}

func executarOCR() (resultadoDoTexto string) {
	tesseract := fmt.Sprintf("tesseract %s stdout -l por", NomeDoArquivoDaImagem)
	cmd := exec.Command("sh", "-c", tesseract)
	out, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatalf("cmd.Run() failed with %s\n", err)
    }
	resultadoDoTexto = string(out)
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
