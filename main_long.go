package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// func getLongPollServer(token string, groupID int64) (*LongPollServer, error) {
// 	vkURL := "https://api.vk.com/method/groups.getLongPollServer"

// 	params := url.Values{}
// 	params.Set("access_token", token)
// 	params.Set("group_id", strconv.Itoa(int(groupID)))
// 	params.Set("v", "5.199")

// 	resp, err := http.PostForm(vkURL, params)
// 	if err != nil {
// 		return nil, fmt.Errorf("ошибка запроса getLongPollServer: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("чтение ответа: %w", err)
// 	}

// 	var result struct {
// 		Response LongPollServer `json:"response"`
// 	}
// 	if err := json.Unmarshal(body, &result); err != nil {
// 		return nil, fmt.Errorf("не удалось распарсить ответ VK: %w", err)
// 	}

// 	return &result.Response, nil
// }

// func sendMessageWithKeyboard(token string, peerID int64, text string, keyboard *Keyboard) error {
// 	token = strings.TrimSpace(token)
// 	if token == "" {
// 		return fmt.Errorf("токен пустой")
// 	}

// 	vkURL := "https://api.vk.com/method/messages.send"

// 	params := url.Values{}
// 	params.Set("access_token", token)
// 	params.Set("v", "5.199")
// 	params.Set("peer_id", strconv.FormatInt(peerID, 10))

// 	params.Set("message", text)

// 	randomID := time.Now().UnixNano()
// 	params.Set("random_id", strconv.FormatInt(randomID, 10))

// 	if keyboard != nil {
// 		kbJSON, err := json.Marshal(keyboard)
// 		if err != nil {
// 			return fmt.Errorf("ошибка сериализации клавиатуры: %w", err)
// 		}
// 		params.Set("keyboard", string(kbJSON))
// 	}

// 	resp, err := http.PostForm(vkURL, params)
// 	if err != nil {
// 		return fmt.Errorf("HTTP запрос: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return fmt.Errorf("чтение ответа: %w", err)
// 	}

// 	fmt.Println("Ответ VK на send:", string(body))

// 	var vkResp map[string]interface{}
// 	if err := json.Unmarshal(body, &vkResp); err != nil {
// 		return fmt.Errorf("не удалось распарсить ответ VK: %w", err)
// 	}

// 	if _, ok := vkResp["error"]; ok {
// 		return fmt.Errorf("VK вернул ошибку: %s", string(body))
// 	}

// 	return nil
// }

// func message(peerID, fromID float64, textRaw string) {

// 	peerIDInt := int64(peerID)
// 	fromIDInt := int64(fromID)
// 	text := strings.ToLower(strings.TrimSpace(textRaw))

// 	log.Printf("Новое сообщение: fromID=%d, peerID=%d, text=%q", fromIDInt, peerIDInt, text)

// 	//Проверка авторизации. Если пользователь не авторизован

// 	//отправляется предолжение авторизации

// 	// fmt.Println("\n\n\n", usersCheckIn)
// 	// fmt.Println("peerid float ", peerID, "\npeerid int ", peerIDInt)
// 	// sendPhoto(peerIDInt, "1.jpg")

// 	time.LoadLocation("Europe/Moscow")

// 	if peerID == 12934363 {
// 		switch text {
// 		case "обнулить":
// 			initRoutes()
// 			initCodes()
// 			dbInsertRoutsInDB()

// 		}
// 	}

// 	if dbAuthOk(int(peerID), usersCheckIn[peerID]) {

// 		//проверяет закончена ли авторизация пользователя
// 		//(ищет в таблице базы данных CheckList - совпадения из userchecklist по peerID и названию команды)
// 		//если все ок запускается функционал "ИГРЫ"
// 		numberTask, _ := strconv.Atoi(dbGetActualTask(usersCheckIn[peerID]))

// 		if numberTask > countTask {

// 			str := fmt.Sprint("Команда " + usersCheckIn[peerID] + " закончила игру")
// 			fmt.Println(str)
// 			dbLog(str, int(peerID))
// 			switch text {

// 			case "выход":
// 				dbExit(int(peerID))
// 				delete(usersCheckIn, peerID)
// 				_ = sendMessageWithKeyboard(token, peerIDInt, "Вы покинули игру", auth())

// 			case "отчет":

// 				_ = sendMessageWithKeyboard(token, peerIDInt, report(usersCheckIn[peerID]), gameOver())

// 			default:
// 				_ = sendMessageWithKeyboard(token, peerIDInt, "Game over", gameOver())

// 			}

// 		} else {

// 			if time.Now().In(location).Before(TIMESTARTGAME) {
// 				fmt.Println("текущее время после анонсированного начала игры\nклавиатура с кнопкой начать")

// 				_ = sendMessageWithKeyboard(token, peerIDInt, fmt.Sprintf("Старт игры: %s", TIMESTARTGAME.Format(time.DateTime)), game())

// 			} else if !getEnteringCodeStatus(peerID) {

// 				fmt.Println("getEnteringCodeStatus false")

// 				switch text {
// 				case "попросить подсказку":
// 					dbGetHelp(peerIDInt)

// 				case "выход":

// 					dbExit(int(peerID))
// 					delete(usersCheckIn, peerID)
// 					_ = sendMessageWithKeyboard(token, peerIDInt, "Вы покинули игру", auth())

// 				case "начать игру":
// 					dbGetTask(peerIDInt, usersCheckIn[peerID])
// 					_ = sendMessageWithKeyboard(token, peerIDInt, "Задание выдано!\nВперед!", gameAfterStart())

// 				case "повторить вопрос":

// 					dbGetTask(peerIDInt, usersCheckIn[peerID])
// 					_ = sendMessageWithKeyboard(token, peerIDInt, "Вперед!", gameAfterStart())

// 				case "статус":
// 					str := dbGetStatus(usersCheckIn[peerID])
// 					_ = sendMessageWithKeyboard(token, peerIDInt, str, gameAfterStart())

// 				case "ввести код":
// 					fmt.Printf("пользователь %d из команды %s пытается ввести код", peerIDInt, usersCheckIn[peerID])
// 					_ = sendMessageWithKeyboard(token, peerIDInt, "Отправьте мне код", gameEnteringCode())
// 					userEnterCode[peerID] = dbGetCommandNameFromPeerId(peerIDInt)

// 				case "правила":
// 					if gameIsActiv(usersCheckIn[peerID]) {
// 						_ = sendMessageWithKeyboard(token, peerIDInt, rules, gameAfterStart())
// 					} else {
// 						_ = sendMessageWithKeyboard(token, peerIDInt, rules, game())
// 					}

// 				default:
// 					if gameIsActiv(usersCheckIn[peerID]) {
// 						fmt.Println("задание уже было выдано/ команда ", usersCheckIn[peerID])
// 						_ = sendMessageWithKeyboard(token, peerIDInt, "Выбирайте команды из списка", gameAfterStart())

// 					} else {
// 						fmt.Println("задание еще не было выдано / команда", usersCheckIn[peerID])
// 						_ = sendMessageWithKeyboard(token, peerIDInt, "Выбирайте команды из списка", game())
// 					}
// 				}
// 			} else if getEnteringCodeStatus(peerID) {
// 				fmt.Println("getEnteringCodeStatus true")

// 				switch text {
// 				case "отмена":
// 					delete(userEnterCode, peerID)
// 					fmt.Println("id пользователя удален из userEnterCode")
// 					_ = sendMessageWithKeyboard(token, peerIDInt, "Выбирайте команды из списка", gameAfterStart())

// 				default:

// 					if gbCheckCode(peerIDInt, text) {
// 						fmt.Println("code is correct")
// 						dbLog(fmt.Sprintf("пользователь ввел верный код -- %s", text), int(peerIDInt))
// 						delete(userEnterCode, peerID)
// 						nextTaskForCommand(usersCheckIn[peerID])
// 					} else {
// 						fmt.Println("code is not correct")

// 						_ = sendMessageWithKeyboard(token, peerIDInt, "Неверный код", gameEnteringCode())
// 						dbLog(fmt.Sprintf("пользователь пытался ввести неверный код -- %s", text), int(peerIDInt))

// 					}
// 				}
// 			}
// 		}

// 	} else {

// 		if dbAuthCheck(textRaw) {
// 			//проверка на наличие команды в списке
// 			// если в списке то записывает ip_peer в мапу userCheckIn
// 			//
// 			usersCheckIn[peerID] = textRaw

// 		}

// 		var b bool

// 		for key, _ := range usersCheckIn {
// 			if key == peerID {
// 				b = true
// 				continue
// 			} else {
// 				b = false
// 			}
// 		}

// 		if b {

// 			if dbCheckPassword(textRaw, usersCheckIn[peerID]) {
// 				// если пароль верный - записывает команду и id_peer в
// 				// базу данных таблица checkgame
// 				// и удаляет id_peer из мапы userChekIn
// 				// если пароль неверный - просит ввести еще раз
// 				//

// 				dbSetCheckGame(peerID, usersCheckIn[peerID])

// 				if gameIsActiv(usersCheckIn[peerID]) {

// 					fmt.Println("gameIsActiv true  line 215")
// 					_ = sendMessageWithKeyboard(token, peerIDInt, "Выбирайте команды из списка", gameAfterStart())

// 				} else {
// 					_ = sendMessageWithKeyboard(token, peerIDInt, "Выбирайте команды из списка", game())

// 				}

// 			} else {

// 				if text == "отмена" {
// 					_ = sendMessageWithKeyboard(token, peerIDInt, "Привет! Чем могу помочь?", auth())
// 					delete(usersCheckIn, peerID)

// 					return
// 				}

// 				_ = sendMessageWithKeyboard(token, peerIDInt, "Введите пароль", enterPass())

// 			}

// 		} else {

// 			switch text {
// 			case "привет", "хай", "hello", "отмена":
// 				_ = sendMessageWithKeyboard(token, peerIDInt, "Привет! Чем могу помочь?", auth())
// 			case "авторизация":
// 				_ = sendMessageWithKeyboard(token, peerIDInt, "Выбери команду из списка", commandList(dbGetCommandList()))
// 			case "назад":
// 				_ = sendMessageWithKeyboard(token, peerIDInt, "Чем могу помочь?", auth())
// 			case "Введите пароль":
// 				fmt.Println("case введите пароль")
// 			case "куда я попал и что делать?":
// 				_ = sendMessageWithKeyboard(token, peerIDInt, "Тут будет описание игры", auth())
// 			default:
// 				_ = sendMessageWithKeyboard(token, peerIDInt, "Чем могу помочь?", auth())

// 			}
// 		}
// 	}
// }

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

// 				// handleMessage(update)

// 				select {
// 				case taskChan <- TaskVK{Data: update}:
// 				default:
// 					log.Println("queue full")
// 				}

// 			}
// 		}
// 	}

// }

// func main() {
// 	initValues()

// 	// initRoutes()
// 	// initCodes()
// 	// dbInsertRoutsInDB()

// 	log.Println("Запуск Long Poll бота...")
// 	dbLogStart()
// 	go worker()
// 	// fmt.Println(TIMESTARTGAME, TIMEENDGAME)
// 	longPollLoop()
// }

// func handleMessage(update LongPollUpdate) {
// 	obj := update.Object

// 	// Достаём message из object.message
// 	msg, ok := obj["message"].(map[string]interface{})
// 	if !ok {
// 		log.Printf("Нет поля message в событии")
// 		return
// 	}

// 	for k, j := range msg {

// 		fmt.Println(k, "   ", j)
// 	}
// 	peerID, ok1 := msg["peer_id"].(float64)
// 	fromID, ok2 := msg["from_id"].(float64)
// 	textRaw, ok3 := msg["text"].(string)

// 	if !(ok1 && ok2 && ok3) {
// 		log.Printf("Не все поля сообщения найдены")
// 		return
// 	}

// 	peerIDInt := int64(peerID)
// 	fromIDInt := int64(fromID)
// 	text := strings.ToLower(strings.TrimSpace(textRaw))

// 	log.Printf("Новое сообщение: fromID=%d, peerID=%d, text=%q", fromIDInt, peerIDInt, text)

// 	//Проверка авторизации. Если пользователь не авторизован

// 	//отправляется предолжение авторизации

// 	// fmt.Println("\n\n\n", usersCheckIn)
// 	// fmt.Println("peerid float ", peerID, "\npeerid int ", peerIDInt)
// 	// sendPhoto(peerIDInt, "1.jpg")

// 	time.LoadLocation("Europe/Moscow")

// 	if peerID == 12934363 {
// 		switch text {
// 		case "обнулить":
// 			initRoutes()
// 			initCodes()
// 			dbInsertRoutsInDB()

// 		}
// 	}

// 	if dbAuthOk(int(peerID), usersCheckIn[peerID]) {

// 		//проверяет закончена ли авторизация пользователя
// 		//(ищет в таблице базы данных CheckList - совпадения из userchecklist по peerID и названию команды)
// 		//если все ок запускается функционал "ИГРЫ"
// 		numberTask, _ := strconv.Atoi(dbGetActualTask(usersCheckIn[peerID]))

// 		if numberTask > countTask {

// 			str := fmt.Sprint("Команда " + usersCheckIn[peerID] + " закончила игру")
// 			fmt.Println(str)
// 			dbLog(str, int(peerID))
// 			switch text {

// 			case "выход":
// 				dbExit(int(peerID))
// 				delete(usersCheckIn, peerID)
// 				_ = sendMessageWithKeyboard(token, peerIDInt, "Вы покинули игру", auth())

// 			case "отчет":

// 				_ = sendMessageWithKeyboard(token, peerIDInt, report(usersCheckIn[peerID]), gameOver())

// 			default:
// 				_ = sendMessageWithKeyboard(token, peerIDInt, "Game over", gameOver())

// 			}

// 		} else {

// 			if time.Now().In(time.Local).Before(TIMESTARTGAME) {
// 				fmt.Println("текущее время после анонсированного начала игры\nклавиатура с кнопкой начать")

// 				_ = sendMessageWithKeyboard(token, peerIDInt, fmt.Sprintf("Старт игры: %s", TIMESTARTGAME.Format(time.DateTime)), game())

// 			} else if !getEnteringCodeStatus(peerID) {

// 				fmt.Println("getEnteringCodeStatus false")

// 				switch text {
// 				case "попросить подсказку":
// 					dbGetHelp(peerIDInt)

// 				case "выход":

// 					dbExit(int(peerID))
// 					delete(usersCheckIn, peerID)
// 					_ = sendMessageWithKeyboard(token, peerIDInt, "Вы покинули игру", auth())

// 				case "начать игру":
// 					dbGetTask(peerIDInt, usersCheckIn[peerID])
// 					_ = sendMessageWithKeyboard(token, peerIDInt, "Задание выдано!\nВперед!", gameAfterStart())

// 				case "повторить вопрос":

// 					dbGetTask(peerIDInt, usersCheckIn[peerID])
// 					_ = sendMessageWithKeyboard(token, peerIDInt, "Вперед!", gameAfterStart())

// 				case "статус":
// 					str := dbGetStatus(usersCheckIn[peerID])
// 					_ = sendMessageWithKeyboard(token, peerIDInt, str, gameAfterStart())

// 				case "ввести код":
// 					fmt.Printf("пользователь %d из команды %s пытается ввести код", peerIDInt, usersCheckIn[peerID])
// 					_ = sendMessageWithKeyboard(token, peerIDInt, "Отправьте мне код", gameEnteringCode())
// 					userEnterCode[peerID] = dbGetCommandNameFromPeerId(peerIDInt)

// 				case "правила":
// 					if gameIsActiv(usersCheckIn[peerID]) {
// 						_ = sendMessageWithKeyboard(token, peerIDInt, rules, gameAfterStart())
// 					} else {
// 						_ = sendMessageWithKeyboard(token, peerIDInt, rules, game())
// 					}

// 				default:
// 					if gameIsActiv(usersCheckIn[peerID]) {
// 						fmt.Println("задание уже было выдано/ команда ", usersCheckIn[peerID])
// 						_ = sendMessageWithKeyboard(token, peerIDInt, "Выбирайте команды из списка", gameAfterStart())

// 					} else {
// 						fmt.Println("задание еще не было выдано / команда", usersCheckIn[peerID])
// 						_ = sendMessageWithKeyboard(token, peerIDInt, "Выбирайте команды из списка", game())
// 					}
// 				}
// 			} else if getEnteringCodeStatus(peerID) {
// 				fmt.Println("getEnteringCodeStatus true")

// 				switch text {
// 				case "отмена":
// 					delete(userEnterCode, peerID)
// 					fmt.Println("id пользователя удален из userEnterCode")
// 					_ = sendMessageWithKeyboard(token, peerIDInt, "Выбирайте команды из списка", gameAfterStart())

// 				default:

// 					if gbCheckCode(peerIDInt, text) {
// 						fmt.Println("code is correct")
// 						dbLog(fmt.Sprintf("пользователь ввел верный код -- %s", text), int(peerIDInt))
// 						delete(userEnterCode, peerID)
// 						nextTaskForCommand(usersCheckIn[peerID])
// 					} else {
// 						fmt.Println("code is not correct")

// 						_ = sendMessageWithKeyboard(token, peerIDInt, "Неверный код", gameEnteringCode())
// 						dbLog(fmt.Sprintf("пользователь пытался ввести неверный код -- %s", text), int(peerIDInt))

// 					}
// 				}
// 			}
// 		}

// 	} else {

// 		if dbAuthCheck(textRaw) {
// 			//проверка на наличие команды в списке
// 			// если в списке то записывает ip_peer в мапу userCheckIn
// 			//
// 			usersCheckIn[peerID] = textRaw

// 		}

// 		var b bool

// 		for key, _ := range usersCheckIn {
// 			if key == peerID {
// 				b = true
// 				continue
// 			} else {
// 				b = false
// 			}
// 		}

// 		if b {

// 			if dbCheckPassword(textRaw, usersCheckIn[peerID]) {
// 				// если пароль верный - записывает команду и id_peer в
// 				// базу данных таблица checkgame
// 				// и удаляет id_peer из мапы userChekIn
// 				// если пароль неверный - просит ввести еще раз
// 				//

// 				dbSetCheckGame(peerID, usersCheckIn[peerID])

// 				if gameIsActiv(usersCheckIn[peerID]) {

// 					fmt.Println("gameIsActiv true  line 215")
// 					_ = sendMessageWithKeyboard(token, peerIDInt, "Выбирайте команды из списка", gameAfterStart())

// 				} else {
// 					_ = sendMessageWithKeyboard(token, peerIDInt, "Выбирайте команды из списка", game())

// 				}

// 			} else {

// 				if text == "отмена" {
// 					_ = sendMessageWithKeyboard(token, peerIDInt, "Привет! Чем могу помочь?", auth())
// 					delete(usersCheckIn, peerID)

// 					return
// 				}

// 				_ = sendMessageWithKeyboard(token, peerIDInt, "Введите пароль", enterPass())

// 			}

// 		} else {

// 			switch text {
// 			case "привет", "хай", "hello", "отмена":
// 				_ = sendMessageWithKeyboard(token, peerIDInt, "Привет! Чем могу помочь?", auth())
// 			case "авторизация":
// 				_ = sendMessageWithKeyboard(token, peerIDInt, "Выбери команду из списка", commandList(dbGetCommandList()))
// 			case "назад":
// 				_ = sendMessageWithKeyboard(token, peerIDInt, "Чем могу помочь?", auth())
// 			case "Введите пароль":
// 				fmt.Println("case введите пароль")
// 			case "куда я попал и что делать?":
// 				_ = sendMessageWithKeyboard(token, peerIDInt, "Тут будет описание игры", auth())
// 			default:
// 				_ = sendMessageWithKeyboard(token, peerIDInt, "Чем могу помочь?", auth())

// 			}
// 		}
// 	}
// }
