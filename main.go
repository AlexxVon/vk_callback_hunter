package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type CallbackRequest struct {
	Type    string          `json:"type"`
	GroupID int             `json:"group_id"`
	Secret  string          `json:"secret"`
	Object  json.RawMessage `json:"object,omitempty"`
	EventID string          `json:"event_id,omitempty"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	log.Printf("[HIT] callback received at %v", time.Now())

	// Сразу пишем 200 OK: VK требует быстрый ответ
	w.WriteHeader(http.StatusOK)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("confirm: read body error: %v", err)
		return
	}
	// log.Printf("[DEBUG] raw body: %s", string(body))

	var req CallbackRequest
	if err := json.Unmarshal(body, &req); err != nil {
		log.Printf("confirm: json error: %v, body=%s", err, string(body))
		return
	}

	expectedSecret := os.Getenv("VK_CALLBACK_SECRET")
	if expectedSecret != "" && req.Secret != expectedSecret {
		log.Println("confirm: invalid secret")
		// Для неверного секрета можно вообще ничего не делать: 200 уже отправлен
		return
	}

	switch req.Type {
	case "confirmation":
		confirmString := os.Getenv("VK_CONFIRM_STRING")
		if confirmString == "" {
			log.Println("confirm: VK_CONFIRM_STRING is not set")
			// Не отправляем ничего: VK получит пустой ответ и пометит адрес как неподтверждённый
			return
		}
		// Только строка подтверждения, без JSON, без переносов
		fmt.Fprint(w, confirmString)
		return

	case "message_read":
		fmt.Println("Type: message_read send to VK")
	confirmString := os.Getenv("VK_CONFIRM_STRING")
		if confirmString == "" {
			log.Println("confirm: VK_CONFIRM_STRING is not set")
			// Не отправляем ничего: VK получит пустой ответ и пометит адрес как неподтверждённый
			return
		}
		// Только строка подтверждения, без JSON, без переносов
		return

		

	case "message_reply":

		fmt.Println("Type: message_reply", req.Type, "Object :", string(req.Object) )

		confirmString := os.Getenv("VK_CONFIRM_STRING")
		if confirmString == "" {
			log.Println("confirm: VK_CONFIRM_STRING is not set")
			// Не отправляем ничего: VK получит пустой ответ и пометит адрес как неподтверждённый
			return
		}
		// Только строка подтверждения, без JSON, без переносов
		fmt.Fprint(w, confirmString)
		return

	case "message_new":
		confirmString := os.Getenv("VK_CONFIRM_STRING")
		if confirmString == "" {
			log.Println("confirm: VK_CONFIRM_STRING is not set")
			// Не отправляем ничего: VK получит пустой ответ и пометит адрес как неподтверждённый
			return
		}
		// Только строка подтверждения, без JSON, без переносов
		fmt.Fprint(w, confirmString)
		
		select {
		case taskChan <- TaskVK{Type: req.Type, Data: req.Object}:
		default:
			log.Println("queue full")
			// Здесь можно добавить метрику или алерт, но не меняем HTTP-ответ: он уже 200
		}
		log.Printf("hendler for message_new processed in %v", time.Since(start))

		return

	default:
		log.Printf("confirm: unknown callback type: %s", req.Type)
		// Ничего не пишем в w: 200 уже отправлен, тело не обязательно
	}
}

		// 	var client struct {
		// 			Message json.RawMessage `json:"message"`
		// 		}

		// 			var msg struct {
		//         FromID  int    `json:"from_id"`
		//         Text    string `json:"text"`
		//         PeerID  int    `json:"peer_id"`
		//         RandomID int   `json:"random_id"`
		//     }
		//     if err := json.Unmarshal(req.Object, &client); err != nil {
		//         log.Printf("parse message_new.object: %v", err)
		//       return
		// 	}
		// 	// log.Printf("[DEBUG] raw object bytes: %q", string(req.Object))

		// 	    if err := json.Unmarshal(client.Message, &msg); err != nil {
		//         log.Printf("parse message_new.object: %v", err)
		//       return
		// 	}

		// // log.Printf("[DEBUG] raw client bytes: %q", string(req.Object))

		// 	peerID := msg.PeerID
		// 	fromID  := msg.FromID
		// 	textRaw := msg.Text

		// 		peerIDInt := int64(peerID)
		// 		// fromIDInt := int64(fromID)
		// 		text := strings.ToLower(strings.TrimSpace(textRaw))

		// 		log.Printf("Новое сообщение: fromID=%d, peerID=%d, text=%q", fromID, peerID, text)
		// 		log.Printf("New message from %d (peer %d): %q", msg.FromID, msg.PeerID, strings.TrimSpace(msg.Text))
		// 		// Тут можно вызвать sendMessage(msg.PeerID, "Привет!")

		// 		message(float64(peerIDInt), text)



func main() {
initValues()

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

http.HandleFunc("/callback/confirm", func(w http.ResponseWriter, r *http.Request) {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("PANIC: %v", r)
            // Даже при панике отдаём 200, чтобы VK не повторял запрос
            w.WriteHeader(http.StatusOK)
        }
    }()
    Handler(w, r)
})

	go worker()
	// http.HandleFunc("/callback/confirm", Handler)

	log.Printf("Server listening on :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}

func message(peerID float64, text string) {
   start := time.Now()
	//Проверка авторизации. Если пользователь не авторизован

	//отправляется предолжение авторизации

	// fmt.Println("\n\n\n", usersCheckIn)
	// fmt.Println("peerid float ", peerID, "\npeerid int ", peerIDInt)
	// sendPhoto(peerIDInt, "1.jpg")

	peerIDInt := int64(peerID)

	if peerID == 12934363 {
		switch text {
		case "обнулить":
			initRoutes()
			initCodes()
			setPassForCommand()
			dbInsertRoutsInDB()

		}
	}

	if dbAuthOk(int(peerID), usersCheckIn[peerID]) {

		//проверяет закончена ли авторизация пользователя
		//(ищет в таблице базы данных CheckList - совпадения из userchecklist по peerID и названию команды)
		//если все ок запускается функционал "ИГРЫ"
		numberTask, _ := strconv.Atoi(dbGetActualTask(usersCheckIn[peerID]))

		if numberTask > countTask {

			str := fmt.Sprint("Команда " + usersCheckIn[peerID] + " закончила игру")
			fmt.Println(str)
			dbLog(str, int(peerID))
			switch text {

			case "выход":
				dbExit(int(peerID))
				delete(usersCheckIn, peerID)
				_ = sendMessageWithKeyboard(token, peerIDInt, "Вы покинули игру", auth())

			case "отчет":

				_ = sendMessageWithKeyboard(token, peerIDInt, report(usersCheckIn[peerID]), gameOver())

			default:
				_ = sendMessageWithKeyboard(token, peerIDInt, "Game over", gameOver())

			}

		} else {

			if time.Now().In(location).Before(TIMESTARTGAME) {
				fmt.Println("текущее время после анонсированного начала игры\nклавиатура с кнопкой начать")

				_ = sendMessageWithKeyboard(token, peerIDInt, fmt.Sprintf("Старт игры: %s", TIMESTARTGAME.Format(time.DateTime)), game())

			} else if !getEnteringCodeStatus(peerID) {

				fmt.Println("getEnteringCodeStatus false")

				switch text {
				case "попросить подсказку":
					dbGetHelp(peerIDInt)

				case "выход":

					dbExit(int(peerID))
					delete(usersCheckIn, peerID)
					_ = sendMessageWithKeyboard(token, peerIDInt, "Вы покинули игру", auth())

				case "начать игру":
					dbGetTask(peerIDInt, usersCheckIn[peerID])
					_ = sendMessageWithKeyboard(token, peerIDInt, "Задание выдано!\nВперед!", gameAfterStart())

				case "повторить вопрос":

					dbGetTask(peerIDInt, usersCheckIn[peerID])
					_ = sendMessageWithKeyboard(token, peerIDInt, "Вперед!", gameAfterStart())

				case "статус":
					str := dbGetStatus(usersCheckIn[peerID])
					_ = sendMessageWithKeyboard(token, peerIDInt, str, gameAfterStart())

				case "ввести код":
					fmt.Printf("пользователь %d из команды %s пытается ввести код", peerIDInt, usersCheckIn[peerID])
					_ = sendMessageWithKeyboard(token, peerIDInt, "Отправьте мне код", gameEnteringCode())
					userEnterCode[peerID] = dbGetCommandNameFromPeerId(peerIDInt)

				case "правила":
					if gameIsActiv(usersCheckIn[peerID]) {
						_ = sendMessageWithKeyboard(token, peerIDInt, rules, gameAfterStart())
					} else {
						_ = sendMessageWithKeyboard(token, peerIDInt, rules, game())
					}

				default:
					if gameIsActiv(usersCheckIn[peerID]) {
						fmt.Println("задание уже было выдано/ команда ", usersCheckIn[peerID])
						_ = sendMessageWithKeyboard(token, peerIDInt, "Выбирайте команды из списка", gameAfterStart())

					} else {
						fmt.Println("задание еще не было выдано / команда", usersCheckIn[peerID])
						_ = sendMessageWithKeyboard(token, peerIDInt, "Выбирайте команды из списка", game())
					}
				}
			} else if getEnteringCodeStatus(peerID) {
				fmt.Println("getEnteringCodeStatus true")

				switch text {
				case "отмена":
					delete(userEnterCode, peerID)
					fmt.Println("id пользователя удален из userEnterCode")
					_ = sendMessageWithKeyboard(token, peerIDInt, "Выбирайте команды из списка", gameAfterStart())

				default:

					if gbCheckCode(peerIDInt, text) {
						fmt.Println("code is correct")
						dbLog(fmt.Sprintf("пользователь ввел верный код -- %s", text), int(peerIDInt))
						delete(userEnterCode, peerID)
						nextTaskForCommand(usersCheckIn[peerID])
					} else {
						fmt.Println("code is not correct")

						_ = sendMessageWithKeyboard(token, peerIDInt, "Неверный код", gameEnteringCode())
						dbLog(fmt.Sprintf("пользователь пытался ввести неверный код -- %s", text), int(peerIDInt))

					}
				}
			}
		}

	} else {

		if dbAuthCheck(text) {
			//проверка на наличие команды в списке
			// если в списке то записывает ip_peer в мапу userCheckIn
			//
			usersCheckIn[peerID] = text

		}

		var b bool

		for key, _ := range usersCheckIn {
			if key == peerID {
				b = true
				continue
			} else {
				b = false
			}
		}

		if b {

			if dbCheckPassword(text, usersCheckIn[peerID]) {
				// если пароль верный - записывает команду и id_peer в
				// базу данных таблица checkgame
				// и удаляет id_peer из мапы userChekIn
				// если пароль неверный - просит ввести еще раз
				//

				dbSetCheckGame(peerID, usersCheckIn[peerID])

				if gameIsActiv(usersCheckIn[peerID]) {

					fmt.Println("gameIsActiv true  line 215")
					_ = sendMessageWithKeyboard(token, peerIDInt, "Выбирайте команды из списка", gameAfterStart())

				} else {
					_ = sendMessageWithKeyboard(token, peerIDInt, "Выбирайте команды из списка", game())

				}

			} else {

				if text == "отмена" {
					_ = sendMessageWithKeyboard(token, peerIDInt, "Привет! Чем могу помочь?", auth())
					delete(usersCheckIn, peerID)

					return
				}

				_ = sendMessageWithKeyboard(token, peerIDInt, "Введите пароль", enterPass())

			}

		} else {

			switch text {
			case "привет", "хай", "hello", "отмена":
				_ = sendMessageWithKeyboard(token, peerIDInt, "Привет! Чем могу помочь?", auth())
			case "авторизация":
				_ = sendMessageWithKeyboard(token, peerIDInt, "Выбери команду из списка", commandList(dbGetCommandList()))
			case "назад":
				_ = sendMessageWithKeyboard(token, peerIDInt, "Чем могу помочь?", auth())
			case "Введите пароль":
				fmt.Println("case введите пароль")
			case "куда я попал и что делать?":
				_ = sendMessageWithKeyboard(token, peerIDInt, "Тут будет описание игры", auth())
			default:
				_ = sendMessageWithKeyboard(token, peerIDInt, "Чем могу помочь?", auth())

			}
		}
	}

	log.Printf("func message processed in %v", time.Since(start))

}

// func longPollLoop() {

// 	server, err := getLongPollServer(token, groupID)
// 	if err != nil {
// 		log.Fatalf("Не удалось получить сервер Long Poll: %v", err)
// 	}

// 	ts := server.Ts
// 	key := server.Key
// 	wait := server.Wait

// 	for {
// 		lpURL := fmt.Sprintf("%s?act=a_check&key=%s&ts=%s&wait=%d", server.Server, key, ts, wait)

// 		resp, err := http.Get(lpURL)
// 		if err != nil {
// 			log.Printf("Ошибка запроса Long Poll: %v, ждём и пробуем снова", err)
// 			time.Sleep(5 * time.Second)
// 			continue
// 		}
// 		defer resp.Body.Close()

// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			log.Printf("Чтение ответа Long Poll: %v", err)
// 			time.Sleep(5 * time.Second)
// 			continue
// 		}

// 		var lpResp LongPollResponse
// 		if err := json.Unmarshal(body, &lpResp); err != nil {
// 			log.Printf("Не удалось распарсить Long Poll ответ: %v", err)
// 			log.Printf("Тело: %s", string(body))
// 			time.Sleep(5 * time.Second)
// 			continue
// 		}

// 		ts = lpResp.Ts // обновляем ts для следующего запроса

// 		for _, update := range lpResp.Updates {
// 			if update.Type == "message_new" {
// 				handleMessage(update)
// 			}
// 		}
// 	}

// }
func sendMessageWithKeyboard(token string, peerID int64, text string, keyboard *Keyboard) error {
	token = strings.TrimSpace(token)
	if token == "" {
		return fmt.Errorf("токен пустой")
	}

	vkURL := "https://api.vk.com/method/messages.send"

	params := url.Values{}
	params.Set("access_token", token)
	params.Set("v", "5.199")
	params.Set("peer_id", strconv.FormatInt(peerID, 10))

	params.Set("message", text)

	randomID := time.Now().UnixNano()
	params.Set("random_id", strconv.FormatInt(randomID, 10))

	if keyboard != nil {
		kbJSON, err := json.Marshal(keyboard)
		if err != nil {
			return fmt.Errorf("ошибка сериализации клавиатуры: %w", err)
		}
		params.Set("keyboard", string(kbJSON))
	}

	resp, err := http.PostForm(vkURL, params)
	if err != nil {
		return fmt.Errorf("HTTP запрос: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("чтение ответа: %w", err)
	}

	fmt.Println("Ответ VK на send:", string(body))

	var vkResp map[string]interface{}
	if err := json.Unmarshal(body, &vkResp); err != nil {
		return fmt.Errorf("не удалось распарсить ответ VK: %w", err)
	}

	if _, ok := vkResp["error"]; ok {
		return fmt.Errorf("VK вернул ошибку: %s", string(body))
	}

	return nil
}
