package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/mattn/go-sqlite3"
)

func dbOpen() *sql.DB {
	db, err := sql.Open("sqlite3", "/data/game.db") //для Amvery
	// db, err := sql.Open("sqlite3", "game.db")

	if err != nil {
		panic(err)
	}

	if err != nil {
		s := "Нет подключения к базе данных"
		dbLog(s, 0)
		fmt.Println(s)
		return nil
	} else {
		return db
	}
}

func dbAuthOk(id int, com string) bool {

	db := dbOpen()

	defer db.Close()

	fmt.Println("authcheckOk")

	var peer_id int
	var comName string

	err := db.QueryRow("SELECT `commandname`, `peerid` FROM `checklist` WHERE `commandname`=? AND `peerid`=? ", com, id).Scan(&comName, &peer_id)

	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}

}

func dbInitUserCheckIN() map[float64]string {
	arr := make(map[float64]string)

	var id float64
	var name string

	db := dbOpen()
	defer db.Close()

	res, err := db.Query("SELECT `commandname`, `peerid` FROM `checklist`")

	if err != nil {
		fmt.Println(err)
	}

	for res.Next() {
		err = res.Scan(&name, &id)
		arr[float64(id)] = name

	}

	return arr
}

func dbGetCommandList() []string {

	var list []string

	db := dbOpen()

	defer db.Close()

	res, err := db.Query("SELECT `commandname`, `password` FROM `auth`")

	if err != nil {
		fmt.Println(err)
	}

	for res.Next() {
		var comName string
		var pass string
		err = res.Scan(&comName, &pass)
		list = append(list, comName)
	}

	return list

}

func dbAuthCheck(comName string) bool {

	var res string
	db := dbOpen()
	defer db.Close()
	fmt.Println("authCheck")
	err := db.QueryRow("SELECT `commandname` FROM `auth` WHERE `commandname`=?", comName).Scan(&res)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println()

	if res != "" {
		return true
	} else {
		return false
	}
}

func dbCheckPassword(pass string, comName string) bool {
	db := dbOpen()

	defer db.Close()

	var p string

	err := db.QueryRow("SELECT  `password` FROM auth WHERE `commandname`=?", comName).Scan(&p)

	fmt.Println("p=", p, "  pass=", pass)
	if err != nil {
		fmt.Println(err)
	}

	if p == pass {
		fmt.Println("pass true")

		return true
	} else {
		fmt.Println("pass false")
		return false

	}

}

func dbCheckGame(peerID float64) bool {

	var res int
	db := dbOpen()
	defer db.Close()
	fmt.Println("authCheck")
	err := db.QueryRow("SELECT `peerid` FROM `checklist` WHERE `commandname`=?", peerID).Scan(&res)

	if err != nil {
		fmt.Print(err)
	}

	if res > 0 {
		return true
	} else {
		return false
	}

}

func dbSetCheckGame(id float64, comName string) {

	db := dbOpen()

	defer db.Close()

	p := int(id)

	_, err := db.Exec(fmt.Sprintf("INSERT INTO `checkList` (`commandName`, `peerId`, `auth_date`) VALUES ('%s', '%d', '%s')", comName, p, time.Now().In(location).In(location).Format(time.DateTime)))
	if err != nil {
		fmt.Println(err)
	} else {
		str := fmt.Sprintf("Пользователь %d успешно авторизовался; команда %s", int(id), comName)

		_, _ = db.Exec(fmt.Sprintf("INSERT INTO `log` (`peerid`, `date`, `event`) VALUES ('%d', '%s', '%s')", int(id), time.Now().In(location).Format(time.DateTime), str))

	}

	fmt.Println("записалось")

}

func dbExit(id int) {
	db := dbOpen()

	defer db.Close()

	var comname string

	err := db.QueryRow("SELECT `commandname` FROM `checklist` WHERE `peerid`=? ", id).Scan(&comname)

	if err != nil {
		fmt.Println(err)
	}

	_, err = db.Exec("DELETE FROM `checklist` WHERE `peerid`=?", id)

	if err != nil {
		s := fmt.Sprint("ошибка выхода игрока из игры - ", err)
		dbLog(s, 0)
		fmt.Println(err)

	} else {

		str := fmt.Sprintf("Пользователь %d покинул игру; команда %s", int(id), comname)

		_, _ = db.Exec(fmt.Sprintf("INSERT INTO `log` (`peerid`, `date`, `event`) VALUES ('%d', '%s', '%s')", int(id), time.Now().In(location).Format(time.DateTime), str))

	}

}

func dbLogStart() {

	db := dbOpen()

	defer db.Close()

	event := "Бот успешно запущен"

	_, err := db.Exec(fmt.Sprintf("INSERT INTO `log` (`peerid`, `date`, `event`) VALUES ('%d', '%s', '%s')", 0, time.Now().In(location).Format(time.DateTime), event))
	if err != nil {

		fmt.Println(err)
	}
}

func dbLog(event string, peerId int) {
	db := dbOpen()

	defer db.Close()

	event += "  Задание - " + dbGetActualTask(usersCheckIn[float64(peerId)])

	_, err := db.Exec(fmt.Sprintf("INSERT INTO `log` (`peerid`, `date`, `event`) VALUES ('%d', '%s', '%s')", peerId, time.Now().In(location).Format(time.DateTime), event))
	if err != nil {

		fmt.Println(err)
	}
}

func dbGetTask(peerid int64, commandName string) string {
	var t Task

	n := dbGetStatus(commandName)

	db := dbOpen()

	defer db.Close()

	err := db.QueryRow("SELECT `actualtask` FROM `tasklist` WHERE `commandname`=?", commandName).Scan(&n)

	str := fmt.Sprintf("SELECT `task%s`, `task%scomment`, `pod%s1`, `pod%s2`, `code%s`, `task%sresp`, `pod%s1time`, `pod%s2time` FROM `tasklist` WHERE `commandname`='%s'",
		n, n, n, n, n, n, n, n, commandName)

	err = db.QueryRow(str).Scan(
		&t.task,
		&t.taskcomment,
		&t.pod1,
		&t.pod2,
		&t.code,
		&t.taskresp,
		&t.pod1time,
		&t.pod2time,
	)

	if err != nil {
		s := fmt.Sprint("ошибка запроса в dbGetTask- ", err)
		dbLog(s, 0)

		fmt.Println(err)
	}

	sendPhoto(peerid, fmt.Sprintf("Задание %s\n%s", n, t.taskcomment), t.task)

	time_temp := time.Now().In(location).Format(time.DateTime)

	fmt.Println("func dbGetTask time.Now - ", time_temp, "status - ", n)
	s := fmt.Sprintf("UPDATE `tasklist` SET `task%sresp`='%s' WHERE `commandname`='%s'", n, time_temp, commandName)

	if t.taskresp == " " || t.taskresp == "" {
		_, err = db.Exec(s)

		if err != nil {
			s := fmt.Sprint("ошибка обновления времени отправки ответа - ", err)
			dbLog(s, 0)

			fmt.Println(err)
		}
	}

	return t.task

}

func gameIsActiv(commandname string) bool {
	db := dbOpen()

	defer db.Close()

	t := ""

	err := db.QueryRow("SELECT `task1resp` FROM `tasklist` WHERE `commandname`=? ", commandname).Scan(&t)

	if err != nil {
		s := fmt.Sprint("ошибка запроса о выдаче первого задания - ", err)
		dbLog(s, 0)

		fmt.Println(err)
	}

	if t == "" || t == " " {

		fmt.Println("func gameIsActive  false")

		fmt.Println(err)
		return false
	} else {
		fmt.Println("func gameIsActive  true")

		return true
	}
}

func dbGetActualTask(comname string) string {

	db := dbOpen()

	defer db.Close()

	var n string

	err := db.QueryRow("SELECT `actualtask` FROM `tasklist` WHERE `commandname`=?", comname).Scan(&n)

	if err != nil {
		s := fmt.Sprint("ошибка получения сведений об актуальном задании (номер задания)- ", err)
		dbLog(s, 0)

		fmt.Println(err)
		return ""
	} else {
		return n
	}

}

func dbGetStatus(commandName string) string {
	location, _ := time.LoadLocation("Europe/Moscow")

	var t Task

	n := dbGetActualTask(commandName)

	db := dbOpen()

	defer db.Close()

	err := db.QueryRow("SELECT `actualtask` FROM `tasklist` WHERE `commandname`=?", commandName).Scan(&n)

	str := fmt.Sprintf("SELECT `task%s`, `pod%s1`, `pod%s2`, `code%s`, `task%sresp`, `pod%s1time`, `pod%s2time` FROM `tasklist` WHERE `commandname`='%s'",
		n, n, n, n, n, n, n, commandName)

	err = db.QueryRow(str).Scan(
		&t.task,
		&t.pod1,
		&t.pod2,
		&t.code,
		&t.taskresp,
		&t.pod1time,
		&t.pod2time,
	)

	if err != nil {
		s := fmt.Sprint("ошибка запроса на выборку dbGetStatus- ", err)
		dbLog(s, 0)

		fmt.Println(err)
	}

	time_t, _ := time.Parse(time.DateTime, t.taskresp)

	tim := time_t.Add(time.Minute*time.Duration(intervalTimePod1) - time.Hour*3).Sub(time.Now().In(location).In(location)) //время до 1й подсказки
	fmt.Println(time_t.Add(time.Minute*time.Duration(intervalTimePod1)), time.Now().In(location))
	if tim.Minutes() < 0 {
		tim = 0
	}

	tim2 := time_t.Add(time.Minute*time.Duration(intervalTimePod2) - time.Hour*3).Sub(time.Now().In(location).Local()) //время до 2й подсказки
	if tim2.Minutes() < 0 {
		tim2 = 0
	}
	str = fmt.Sprintf("%dмин %dсек", int(tim.Minutes()), int(tim.Seconds())-int(tim.Minutes())*60)
	str2 := fmt.Sprintf("%dмин %dсек", int(tim2.Minutes()), int(tim2.Seconds())-int(tim2.Minutes())*60)
	// fmt.Println(tim.Seconds(), tim.Minutes())

	resp := fmt.Sprintf("Команда: %s\nТекущее задание №%s\nПолучено: %s\nДо 1й подсказки: %s\nДо 2й подсказки: %s",
		commandName,
		n,
		t.taskresp,
		str,
		str2,
	)

	return resp

}

func gbCheckCode(id int64, code string) bool {

	n := dbGetActualTask(usersCheckIn[float64(id)])

	var res string

	db := dbOpen()

	defer db.Close()

	str := fmt.Sprintf("SELECT `code%s` FROM `tasklist` WHERE `commandname`= '%s'", n, usersCheckIn[float64(id)])

	fmt.Println(str)

	err := db.QueryRow(str).Scan(&res)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("code from table ", res)

	if strings.ToLower(code) == strings.ToLower(res) {
		dbSetTimeTaskEnd(n, id)
		dbSetActualTask(usersCheckIn[float64(id)])

		return true

	} else {
		return false
	}

}

func dbSetTimeTaskEnd(actualtask string, id int64) {
	db := dbOpen()
	defer db.Close()

	fmt.Println("dbSetTimeTaskEnd - actualtask ", actualtask, "command name - ", usersCheckIn[float64(id)])
	str := fmt.Sprintf("UPDATE  `tasklist` SET `task%send`='%s' WHERE `commandname`='%s'", actualtask, time.Now().In(location).Format(time.DateTime), usersCheckIn[float64(id)])

	_, err := db.Exec(str)

	if err != nil {
		s := fmt.Sprint("ошибка установки времени окончании задания - ", err)
		dbLog(s, 0)

		fmt.Println(err)
	}
}

func dbSetActualTask(comname string) {

	status := dbGetActualTask(comname)

	fmt.Println(status)

	n, err := strconv.Atoi(status)

	if err != nil {
		fmt.Println(err)
	}

	db := dbOpen()
	defer db.Close()

	str := fmt.Sprintf("UPDATE  `tasklist` SET `actualtask`=%d WHERE `commandname`='%s'", n+1, comname)

	_, err = db.Exec(str)

	if err != nil {
		s := fmt.Sprint("ошибка установки актуального задания (актуал +1)- ", err)
		dbLog(s, 0)

		fmt.Println(err)
	}

}

func dbGetCommandNameFromPeerId(id int64) string {

	db := dbOpen()

	defer db.Close()

	var res string

	err := db.QueryRow("SELECT `commandname` FROM `checklist` WHERE `peerid`=?", id).Scan(&res)

	if err != nil {
		s := fmt.Sprint("ошибка запроса названия команды dbGetCommandNameFromPeerId - ", err)
		dbLog(s, 0)

		fmt.Println("dbGetCommandNameFromId ", err)
	}

	return res
}

func dbGetHelp(id int64) {

	var t Task
	n := dbGetActualTask(usersCheckIn[float64(id)])

	db := dbOpen()

	defer db.Close()

	str := fmt.Sprintf("SELECT `task%s`, `pod%s1`, `pod%s1time`, `pod%s1comment`,"+
		"`pod%s2`, `pod%s2time`, `pod%s2comment`,"+
		"`code%s`, `task%sresp` FROM `tasklist`"+
		"WHERE `commandname`='%s'",
		n, n, n, n,
		n, n, n,
		n, n, usersCheckIn[float64(id)])

	err := db.QueryRow(str).Scan(
		&t.task,
		&t.pod1,
		&t.pod1time,
		&t.pod1comment,
		&t.pod2,
		&t.pod1time,
		&t.pod2comment,
		&t.code,
		&t.taskresp,
	)

	if err != nil {
		s := fmt.Sprint("ошибка получения данных задания и подсказок из бд - ", err)
		dbLog(s, 0)

		fmt.Println("ошибка получения данных задания и подсказок из бд -", err)
	}

	time_taskresp, _ := time.Parse(time.DateTime, t.taskresp)

	tim := time_taskresp.Add(time.Minute*time.Duration(intervalTimePod1) - time.Hour*3).Sub(time.Now().In(location).Local())
	fmt.Println(time_taskresp.Add(time.Minute*time.Duration(intervalTimePod1)), time.Now().In(location))
	if tim.Minutes() < 0 {
		tim = 0
	}

	time_help1 := time_taskresp.Add(time.Minute*time.Duration(intervalTimePod1) - time.Hour*3).Sub(time.Now().In(location).Local())
	time_help2 := time_taskresp.Add(time.Minute*time.Duration(intervalTimePod2) - time.Hour*3).Sub(time.Now().In(location).Local())

	fmt.Println("\n", int(time_help1.Seconds()), "\n", int(time_help2.Seconds()))

	// fmt.Println("time +intarvalTimePod1", time_taskresp.Add(time.Minute*time.Duration(intervalTimePod1)),

	// 	"\n", time.Now().In(location).Format(time.DateTime),
	// 	"\n", time_taskresp.Add(time_help1), "\t", time.Now().In(location).After(time_taskresp.Add(time_help1)),
	// 	"\n", time_taskresp.Add(time.Minute*time.Duration(intervalTimePod2)), "\t", time.Now().In(location).Before(time_taskresp.Add(time_help2)),
	// 	"\n", time_taskresp.Add(time.Minute*time.Duration(intervalTimePod2)), "\t", time.Now().In(location).After(time_taskresp.Add(time.Minute*time.Duration(intervalTimePod2))),
	// )

	if int(time_help2.Seconds()) <= 0 {

		fmt.Println("отправляю первую подсказку команде ", usersCheckIn[float64(id)])

		str := fmt.Sprintf("pod%s1", n)
		sendPhoto(id, t.pod1comment, dbGetPathHelp(str, id))

		fmt.Println("отправляю второую подсказку команде ", usersCheckIn[float64(id)])
		str = fmt.Sprintf("pod%s2", n)
		sendPhoto(id, t.pod2comment, dbGetPathHelp(str, id))
		str = fmt.Sprintf("UPDATE `tasklist` SET `pod%s2time`='%s' WHERE `commandname`='%s'", n, time.Now().In(location).Format(time.DateTime), usersCheckIn[float64(id)])
		if t.pod2time == " " || t.pod2time == "" {
			_, err = db.Exec(str)

			if err != nil {
				s := fmt.Sprint("ошибка установки времени отправки 2й подсказки - ", err)
				dbLog(s, 0)

				fmt.Println(err)
			}
		}

	}

	if int(time_help1.Seconds()) <= 0 && int(time_help2.Seconds()) > 0 {

		fmt.Println("отправляю первую подсказку команде ", usersCheckIn[float64(id)])

		str := fmt.Sprintf("pod%s1", n)
		sendPhoto(id, t.pod1comment, dbGetPathHelp(str, id))

		str = fmt.Sprintf("UPDATE `tasklist` SET `pod%s1time`='%s' WHERE `commandname`='%s'", n, time.Now().In(location).Format(time.DateTime), usersCheckIn[float64(id)])
		if t.pod1time == " " || t.pod1time == "" {
			_, err = db.Exec(str)

			if err != nil {
				s := fmt.Sprint("ошибка установки времени отправки 1й подсказки - ", err)
				dbLog(s, 0)

				fmt.Println(err)
			}
		}

	} else {
		_ = sendMessageWithKeyboard(token, id, dbGetStatus(usersCheckIn[float64(id)]), gameAfterStart())
	}

	// str = fmt.Sprintf("%dмин %dсек", int(tim.Minutes()), int(tim.Seconds())-int(tim.Minutes())*60)

	// fmt.Println(tim.Seconds(), tim.Minutes())

}

func dbGetPathHelp(pole string, id int64) string {

	db := dbOpen()

	defer db.Close()

	str := fmt.Sprintf("SELECT `%s` FROM `tasklist` WHERE `commandname`='%s'", pole, usersCheckIn[float64(id)])

	var resp string

	err := db.QueryRow(str).Scan(&resp)

	if err != nil {
		s := fmt.Sprint("ошибка получения пути файла подсказки - ", err)
		dbLog(s, 0)

		fmt.Println("Ошибка получения пути файла подсказки -", err)
		return ""
	}

	return resp

}

func report(commandname string) string {

	var rep, t_start, t_end string

	var total_time time.Duration

	db := dbOpen()

	defer db.Close()

	for i := 1; i <= countTask; i++ {

		str := fmt.Sprintf("SELECT `task%dresp`, `task%dend` FROM `tasklist` WHERE `commandname`='%s'", i, i, commandname)

		err := db.QueryRow(str).Scan(&t_start, &t_end)

		if err != nil {

			s := fmt.Sprint("ошибка при составлении отчета - ", err)
			dbLog(s, 0)
			fmt.Println("ошибка в отчете - ", err)

			return s

		}

		time_end, err := time.Parse(time.DateTime, t_end)
		time_start, err := time.Parse(time.DateTime, t_start)
		diffTime := time_end.Sub(time_start)

		rep += fmt.Sprintf("Задание %d: %s\n", i, diffTime)
		total_time += diffTime
	}

	rep += "Общее время: " + total_time.String()

	return rep

}

func setPassForCommand() {
	db := dbOpen()
	file, err := os.Create("/data/codes.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer file.Close()

	defer db.Close()

	list := dbGetCommandList()

	for i := 0; i < len(list); i++ {

		pass := randPass()

		_, err := db.Exec(fmt.Sprintf("UPDATE `auth` SET `password`='%s' WHERE `commandName`='%s'", pass, list[i]))
		if err != nil {

			s := fmt.Sprint("ошибка записи пароля команды в базу данных", err)
			fmt.Println(s)
			dbLog(s, 0)
		}

		writer := bufio.NewWriter(file)

		s := fmt.Sprintf("%s - %s\n", list[i], pass)
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




		

}
