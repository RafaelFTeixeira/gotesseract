package main

import (
	"fmt"
	"os"
	"os/exec"
	"net/http"
	"io"
	"encoding/json"
	"strings"
	"io/ioutil"
	"time"
	)
	
func main() {
	http.HandleFunc("/", obterImagemDaUrl)
	http.HandleFunc("/erros", obterErros)
	http.ListenAndServe(":3000", nil)
}

func catalogarErro(mensagem string) {
	errosAntigos := obterMensagensDosErros()
    bytes := []byte(time.Now().Format("2006-01-02 15:04:05")+ " - " + mensagem+ "\n" + errosAntigos)
    ioutil.WriteFile("erros.txt", bytes, 0)
}

func obterMensagensDosErros() (texto string) {
	binario, erro := ioutil.ReadFile("erros.txt")
	if erro != nil {
		texto = ""
        return
    }
	texto = string(binario)
	return
}

func obterErros(w http.ResponseWriter, r *http.Request) {
	erros := obterMensagensDosErros()
 	w.Write([]byte(erros))
}

func obterImagemDaUrl(w http.ResponseWriter, r *http.Request) {
	imagens, ok := r.URL.Query()["image"]
    if !ok || len(imagens[0]) < 1 {
        json.NewEncoder(w).Encode("Informe o link da imagem. Exemplo: http://localhost:3000/?image=hello.png")
        return
	}

	imagem := string(imagens[0])
	nomeDaImagem := downloadDaImagem(imagem)
	if "" != nomeDaImagem {
		resultadoDoTexto := executarOCR(nomeDaImagem)
		json.NewEncoder(w).Encode(resultadoDoTexto)
		os.Remove(nomeDaImagem)
	} else {
		w.Write([] byte("Erro a fazer download da imagem"))
	}
}

func executarOCR(nomeDaImagem string) (resultadoDoTexto string) {
	tesseract := fmt.Sprintf("tesseract %s stdout -l por", nomeDaImagem)
	cmd := exec.Command("sh", "-c", tesseract)
	out, err := cmd.CombinedOutput()
    if err != nil {
		catalogarErro("erro na imagem: "+nomeDaImagem)
		resultadoDoTexto = ""
		return
    }
	resultadoDoTexto = string(out)
	return
}

func obterNomeDaImagem(url string) (nomeDaImagem string) {
	url = strings.Replace(url, "/", "", -1)
	url = strings.Replace(url, ":", "", -1)
	url = strings.Replace(url, ".", "", -1)
	url = strings.Replace(url, "?", "", -1)
	url = strings.Replace(url, "=", "", -1)
	url = strings.Replace(url, "#", "", -1)
	url = strings.Replace(url, "&", "", -1)
	nomeDaImagem = fmt.Sprintf("%s.png", url)
	return
}

func downloadDaImagem(url string) (nomeDaImagem string) {
	resposta, erro := http.Get(url)

    if erro != nil {
		nomeDaImagem = ""
		return
    }
	defer resposta.Body.Close()
	nomeDaImagem = obterNomeDaImagem(url)
	arquivo, erro := os.Create(nomeDaImagem)
	
    if erro != nil {
		nomeDaImagem = ""
		return
    }
    
    _, erro = io.Copy(arquivo, resposta.Body)
    if erro != nil {
		nomeDaImagem = ""
		return
	}
	
	arquivo.Close()
	return
}
