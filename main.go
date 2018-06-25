package main

import (
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
	)

func executar(imagem string) {
	tesseract := fmt.Sprintf("tesseract %s texto -l eng", imagem)
	cmd := exec.Command("cmd", "/C", tesseract)
	cmd.Run()
}

func obterTexto() (texto string) {
	binario, erro := ioutil.ReadFile("texto.txt")
	if erro != nil {
        fmt.Print(erro)
    }
	texto = string(binario)
	return
}
	
func main() {
	imagem := os.Args[1]
	executar(imagem)
	texto := obterTexto()
	fmt.Println(texto)
}