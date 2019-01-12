package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// HandlePicture ... Tries to save picture at SNE-Life blog using API
func HandlePicture(msg *tgbotapi.Message) {
	photo := *msg.Photo

	r1, err := http.Get("https://api.telegram.org/bot" + Token +
		"/getFile?file_id=" + photo[2].FileID)
	if err != nil {
		log.Print("[HandlePicture], ", err)
		return
	} else if r1.StatusCode != http.StatusOK {
		log.Print("[HandlePicture], Not 200 response")
		return
	}
	defer r1.Body.Close()

	// Read response body, convert to json
	bodyBytes, _ := ioutil.ReadAll(r1.Body)
	var raw map[string]interface{}
	json.Unmarshal(bodyBytes, &raw)
	raw = raw["result"].(map[string]interface{})

	filePath := raw["file_path"].(string)

	// Load image from Telegram server
	r2, err := http.Get("https://api.telegram.org/file/bot" + Token + "/" + filePath)
	if err != nil {
		log.Print("[HandlePicture], ", err)
		return
	}
	defer r2.Body.Close()

	// Create temp file (unnecessary)
	file, err := os.Create("/tmp/aaaaa.jpg")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	// Copy data from resopnse to file
	_, err = io.Copy(file, r2.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Try upload image to Blog
	UploadImage()
}
