package controllers

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/pedrohmachado/brejaria-backend/models"
	u "github.com/pedrohmachado/brejaria-backend/utils"

	"github.com/gorilla/mux"
)

// UploadImagemProduto sobe imagem produto
var UploadImagemProduto = func(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("image")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	defer file.Close()

	tempFile, err := ioutil.TempFile("static/temp-images", "upload-*.png")

	if err != nil {
		fmt.Println(err)
		//return
	}

	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
		//return
	}

	tempFile.Write(fileBytes)

	fmt.Println(tempFile.Name())

	params := mux.Vars(r)
	idProduto, err := strconv.Atoi(params["id"])

	if err != nil {
		u.Respond(w, u.Message(false, "Erro enquanto decodificava corpo da requisição"))
	}

	imagemProduto := &models.ImagemProduto{}
	imagemProduto.CaminhoImagem = tempFile.Name()
	imagemProduto.IDProduto = uint(idProduto)

	resp := imagemProduto.Cria()
	u.Respond(w, resp)
}

// GetImagemProduto recupera imagem do produto
var GetImagemProduto = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	IDProduto, err := strconv.Atoi(params["id"])

	if err != nil {
		u.Respond(w, u.Message(false, "Erro enquanto decodificava corpo da requisição"))
	}

	caminhoImagem := models.GetImagemProduto(uint(IDProduto))

	file, err := os.Open("C:/Users/pedro.machado/Projects/src/github.com/pedrohmachado/brejaria-backend/" + caminhoImagem)

	if err != nil {
		log.Fatal(err)
		u.Respond(w, u.Message(false, "Erro enquanto recuperava imagem"))
	}

	w.Header().Set("Content-Type", "image/png")
	io.Copy(w, file)

}
