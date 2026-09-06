package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

type vkError struct {
	ErrorMsg string `json:"error_msg"`
}

func vkCall(method string, params url.Values, result interface{}) error {

	baseURL := "https://api.vk.com/method/"
	reqURL := baseURL + method
	params.Set("v", "5.199")

	resp, err := http.PostForm(reqURL, params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP status %d", resp.StatusCode)
	}

	var wrapper struct {
		Response json.RawMessage `json:"response"`
		Error    vkError         `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return err
	}

	if wrapper.Error.ErrorMsg != "" {
		return fmt.Errorf("VK error: %s", wrapper.Error.ErrorMsg)
	}

	if len(wrapper.Response) == 0 {
		return fmt.Errorf("empty response")
	}

	// Распаковываем response сразу в переданный result
	return json.Unmarshal(wrapper.Response, result)
}

// getUploadServer получает URL для загрузки фото в сообщения
func getUploadServer(token string, peerID int64) (string, error) {
	var resp struct {
		UploadURL string `json:"upload_url"`
	}
	params := url.Values{}
	params.Set("access_token", token)
	params.Set("peer_id", fmt.Sprintf("%d", peerID))

	err := vkCall("photos.getMessagesUploadServer", params, &resp)
	if err != nil {
		return "", err
	}
	return resp.UploadURL, nil
}

// uploadPhoto загружает фото на сервер ВК и возвращает server, photo, hash
func uploadPhoto(uploadURL string, filePath string) (server, photo, hash string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	bodyBuf := &bytes.Buffer{}
	writer := multipart.NewWriter(bodyBuf)

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return
	}

	err = writer.Close()
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", uploadURL, bodyBuf)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var upResp struct {
		Server int    `json:"server"`
		Photo  string `json:"photo"`
		Hash   string `json:"hash"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&upResp); err != nil {
		return
	}

	return fmt.Sprintf("%d", upResp.Server), upResp.Photo, upResp.Hash, nil
}

// saveMessagesPhoto сохраняет фото после загрузки и возвращает attachment
func saveMessagesPhoto(token string, server, photo, hash string) (int, error) {
	var resp []struct {
		ID int `json:"id"`
	}
	params := url.Values{}
	params.Set("access_token", token)
	params.Set("server", server)
	params.Set("photo", photo)
	params.Set("hash", hash)
	params.Set("random_id", strconv.Itoa(rand.Intn(1e9)))

	err := vkCall("photos.saveMessagesPhoto", params, &resp)
	if err != nil {
		return 0, err
	}
	if len(resp) == 0 {
		return 0, fmt.Errorf("no photo id returned")
	}

	// attachment формат: photo<owner_id>_<media_id>
	// owner_id можно взять из ответа или использовать свой (для бота это важно)
	// В saveMessagesPhoto owner_id — это ID пользователя/бота, который загрузил фото.
	// Для простоты берем ID из самого ID фото (VK возвращает <owner_id>_<id>)
	return resp[0].ID, nil
}

func sendMessageWithPhoto(token, message string, attachment int, peerID int64) error {
	params := url.Values{}
	params.Set("access_token", token)
	params.Set("peer_id", fmt.Sprintf("%d", peerID))
	params.Set("message", message)
	params.Set("attachment", fmt.Sprintf("photo%d_%d", peerID, attachment))
	params.Set("random_id", "0")

	var resp int64

	return vkCall("messages.send", params, &resp)
}


func sendPhoto(peerID int64, msg, photoPath string) {
	fmt.Println("func sendphoto")
	// 1. Получаем URL загрузки
	uploadURL, err := getUploadServer(token, int64(peerID))
	if err != nil {
		msg += "Ошибка получения upload URL"
		fmt.Println("Ошибка получения upload URL:", err)
		return
	}
	fmt.Println("upload url   ", uploadURL)
	// 2. Загружаем фото
	server, photo, hash, err := uploadPhoto(uploadURL, photoPath)
	if err != nil {
		msg += "Ошибка загрузки фото"
		fmt.Println("Ошибка загрузки фото:", err, photoPath)
		return
	}

	fmt.Println("server", server, "\n", "photo", photo, "\n", "hash:", hash, "\n", "err", err)

	// 3. Сохраняем фото
	attachment, err := saveMessagesPhoto(token, server, photo, hash)
	if err != nil {
		msg+="Ошибка сохранения фото"
		fmt.Println("Ошибка сохранения фото:", err)
		return
	}

	fmt.Println("attachment", attachment)

	// 4. Отправляем сообщение с фото

	err = sendMessageWithPhoto(token, msg, attachment, peerID)
	if err != nil {
		fmt.Println("Ошибка отправки сообщения:", err)
		return
	}

	fmt.Println("Фото успешно отправлено! Attachment:", attachment)
}
