package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"encoding/hex"
	"crypto/aes"
)

type plainTextStruct struct{
	PlainText string 
}

type encryptedTextStruct struct{
	EncrypetedText string 
}

const (
	key = "myverystrongpasswordo32bitlength"
)

func encrypt(plainText string) (string){
	cipherBlock, err := aes.NewCipher([]byte(key))
	if err != nil {  
		log.Fatal("NewCipher(%d bytes) = %s", len(key), err)  
		panic(err)  
	 }
	 out := make([]byte, len(plainText))
	 cipherBlock.Encrypt(out, []byte(plainText))
	 return hex.EncodeToString(out)  
}

type encryptionHandler struct {
	
}

func (encHandler encryptionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var inputDataStruct plainTextStruct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&inputDataStruct)
	defer r.Body.Close()
	if(err != nil) {
		log.Println("Error occurred while converting to JSON")
		w.Write([]byte("Error occurred while converting to JSON"))
	} else {
		encryptedData := encrypt(inputDataStruct.PlainText)
		outputDataStruct := encryptedTextStruct{encryptedData}
		jsonData, err := json.Marshal(outputDataStruct)
		if(err != nil) {
			log.Fatal("Error occurred while converting to JSON")
			w.Write([]byte("Error occurred while converting to JSON"))
		} else {
			w.Write(jsonData)
		}
	}
}

func main() {

	router := mux.NewRouter()

	encHandler := encryptionHandler{}
	router.Handle("/encrypt/", encHandler).Methods("POST")

	http.ListenAndServe(":8081", router)
}