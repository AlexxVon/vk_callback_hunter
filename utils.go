package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

func nextTaskForCommand(commandname string) {

	var listUsers []int64

	var temp int64

	db := dbOpen()
	defer db.Close()

	res, err := db.Query("SELECT `peerid` FROM `checklist` WHERE `commandname`=?", commandname)

	if err != nil {
		fmt.Print("func nexttaskfromcommand.  ", err)
	}

	for res.Next() {
		res.Scan(&temp)
		listUsers = append(listUsers, temp)
	}

	numberTask, _ := strconv.Atoi(dbGetActualTask(commandname))

	if numberTask > countTask {
		for _, k := range listUsers {

			_ = sendMessageWithKeyboard(token, k, "Game over", gameOver())
		}
	} else {

		for _, k := range listUsers {

			_ = sendMessageWithKeyboard(token, k, "Следующее задание", gameAfterStart())
			dbGetTask(k, usersCheckIn[float64(k)])

		}
	}

}

func initRoutes() {

	route = []int{
		1, 2, 19, 17, 16, 14, 13, 12, 9, 8, 7, 6,
	}

	routes = append(routes, route[:])

	routes = append(routes, arrObratno(route))

	route = []int{
		1, 18, 20, 3, 17, 12, 14, 10, 9, 8, 7, 5,
	}

	routes = append(routes, route[:])

	routes = append(routes, arrObratno(route))

	route = []int{
		14, 19, 1, 2, 17, 15, 14, 12, 9, 8, 7, 6,
	}

	routes = append(routes, route[:])

	routes = append(routes, arrObratno(route))

	route = []int{
		8, 9, 11, 13, 14, 16, 17, 2, 19, 3, 4, 6,
	}

	routes = append(routes, route[:])

	routes = append(routes, arrObratno(route))
	for i := 0; i < len(routes); i++ {
		fmt.Println(routes[rand.IntN(len(routes)-1)])

	}

}

func arrObratno(arr []int) []int {
	var t [12]int
	for i, j := 0, len(route)-1; j >= 0; i, j = i+1, j-1 {

		t[i] = arr[j]
	}
	return t[:]

}

func dbInsertRoutsInDB() {
	list := dbGetCommandList()

	db := dbOpen()

	defer db.Close()
	var str string
	str += fmt.Sprintln()

	err, _ := db.Exec("DELETE FROM `tasklist`")
	if err != nil {
		fmt.Println("ошибка удаления данных из tasklist ", err)
	}

	for i := 0; i < len(list); i++ {
		n := rand.IntN(len(routes))

		fmt.Println("func dbInsertRoutsinDB. - n =", n)

		s := fmt.Sprintf("INSERT INTO `tasklist` (`commandname`) VALUES ('%s')", list[i])

		_, err := db.Exec(s)

		if err != nil {
			fmt.Println("ошибка добавления названия команды в базу данных func dbInsertRoutsInDB - ", err)
		}
		for j := 0; j < len(routes[n]); j++ {

			str := fmt.Sprintf("UPDATE `tasklist` SET `task%d`='img/task%d.HEIC', `pod%d1`='img/pod%d1.HEIC', `pod%d2`='img/pod%d2.HEIC',"+
				"`code%d`='%s' , `pod%d1comment`='%s',"+
				"`task%dcomment`='%s', `pod%d2comment`='%s'"+
				"WHERE `commandname`='%s'",
				j+1, routes[n][j], j+1, routes[n][j], j+1, routes[n][j],
				j+1, CODES[routes[n][j]], j+1, COMMENTS_POD1[routes[n][j]],
				j+1, COMMENTS_TASK[routes[n][j]], j+1, COMMENTS_POD2[routes[n][j]],
				list[i],
			)

			fmt.Println(str)
			_, err := db.Exec(str)

			if err != nil {

				fmt.Println("ошибка добавления заданий по маршруту func dbInsertRoutsInDB - ", err)
			}
		}
	}

}

func dbInsertingRoutes(commandname string, list []string) {

}

func randPass() string {
	letters := "0123456789"

	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.IntN(len(letters))]
	}
	return string(b)

}


func randString() string {

	// letters := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// letters := "0123456789"
	letters := "0"

	b := make([]byte, 8)
	for i := range b {
		b[i] = letters[rand.IntN(len(letters))]
	}
	return string(b)

}

func initCodes() {

	for i := 1; i < len(COMMENTS_TASK)+1; i++ {
		CODES[i] = randString()
		fmt.Println(CODES[i])
	}

	file, err := os.Create("/data/codes.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer file.Close()

	for i := 1; i < len(CODES)+1; i++ {

		writer := bufio.NewWriter(file)

		s := fmt.Sprintf("Task%d - %s\n", i, CODES[i])
		_, err = writer.WriteString(s)
		if err != nil {
			fmt.Println("Ошибка записи:", err)
		}

		// Обязательно сбрасываем буфер, чтобы данные попали в файл
		err = writer.Flush()
		if err != nil {
			fmt.Println("Ошибка сброса буфера:", err)
		}
	}

	// ## 3. Буферизованная запись (bufio.Writer)
	// Этот подход повышает производительность, если вы делаете много мелких операций записи. Данные накапливаются во внутреннем буфере, и запись происходит большими блоками. [9](https://labex.io/ru/tutorials/go-how-to-use-buffered-io-in-golang-419747)[10](https://purpleschool.ru/knowledge-base/golang/work-with-data/bufio)[4](https://habr.com/ru/companies/otus/articles/868658/)

	// **Пример:**

}

func worker() {

	for task := range taskChan {

start := time.Now()
		log.Printf("task len", taskChan)

		log.Printf("[WORKER] processing at %v (event_id=%s)", time.Now().In(location), task.Data)
	
		var client struct {
			Message json.RawMessage `json:"message"`
		}

		var msg struct {
			FromID   int    `json:"from_id"`
			Text     string `json:"text"`
			PeerID   int    `json:"peer_id"`
			RandomID int    `json:"random_id"`
		}
		if err := json.Unmarshal(task.Data, &client); err != nil {
			log.Printf("parse taskData: %v", err)
			return
		}
		// log.Printf("[DEBUG] raw object bytes: %q", string(req.Object))

		if err := json.Unmarshal(client.Message, &msg); err != nil {
			log.Printf("parse clinetMessage: %v", err)
			return
		}

		// log.Printf("[DEBUG] raw client bytes: %q", string(req.Object))

		peerID := msg.PeerID
		fromID := msg.FromID
		textRaw := msg.Text

		peerIDInt := int64(peerID)
		// fromIDInt := int64(fromID)
		text := strings.ToLower(strings.TrimSpace(textRaw))

		log.Printf("Новое сообщение: fromID=%d, peerID=%d, text=%q", fromID, peerID, text)
		log.Printf("New message from %d (peer %d): %q", msg.FromID, msg.PeerID, strings.TrimSpace(msg.Text))
		// Тут можно вызвать sendMessage(msg.PeerID, "Привет!")

		message(float64(peerIDInt), text)
log.Printf("message_new processed in %v", time.Since(start))
	}

}
