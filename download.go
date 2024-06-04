package main

import (
    "fmt"
    "image/png"
    "golang.org/x/image/bmp"
    "net/http"
    "os"
    "io"
    "log"
)

func downloadImage(url, filepath string) error {
    // Faz o request para a URL
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Cria um arquivo vazio para salvar a imagem
    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    // Copia o conteúdo do response para o arquivo
    _, err = io.Copy(out, resp.Body)
    return err
}

func convertPngToBmp(inputPath, outputPath string) error {
    // Abre o arquivo PNG
    inputFile, err := os.Open(inputPath)
    if err != nil {
        return err
    }
    defer inputFile.Close()

    // Decodifica a imagem PNG
    img, err := png.Decode(inputFile)
    if err != nil {
        return err
    }

    // Cria um novo arquivo BMP
    outputFile, err := os.Create(outputPath)
    if err != nil {
        return err
    }
    defer outputFile.Close()

    // Codifica a imagem como BMP
    err = bmp.Encode(outputFile, img)
    if err != nil {
        return err
    }

    return nil
}

func main() {
    url := os.Args[1]
	nomeImagem := os.Args[2]
	nomeImagemConvertida := os.Args[3]

    // Baixa a imagem
    err := downloadImage(url, nomeImagem)
    if err != nil {
        log.Fatalf("Erro ao baixar a imagem: %v", err)
    }
    fmt.Println("Imagem baixada com sucesso:", nomeImagem)

    // Converte a imagem PNG para BMP
    err = convertPngToBmp(nomeImagem, nomeImagemConvertida)
    if err != nil {
        log.Fatalf("Erro ao converter a imagem: %v", err)
    }
    fmt.Println("Imagem convertida com sucesso para BMP:", nomeImagemConvertida)
}
