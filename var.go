package main

import (
	"encoding/json"
	"fmt"
	"time"
)

var (
	token         = "vk1.a.b7rXyEiMmYlhoGKEM3XUF7fRGGRpO3OPlKIt-YTCz0GMy9s9I0KFclXlwAI1YIhOHv8BPvvaJIzILFHuRfL8HzonXpA_vl0rGe2Gjp2bYqp6BOQmqdtfAOmQI1H23bGp7lGaL6lux8jSOYK_tyt1pZXnI23HULtu6fIBEqJCnT-mOdHioPqBO7ImfgLHcUp4EsEjuBOJeLTQuL3K9lb83A" // групповой токен
	groupID int64 = 240576964                                                                                                                                                                                                                      // ID сообщества (без club)
)

// var VK_CALLBACK_SECRET string = "asdfkK73hdfujfbgweiu3kfsjhdbacajegfg"
// var VK_CONFIRM_STRING string = "8895848d"

type KeyboardAction struct {
	Type    string `json:"type"`
	Label   string `json:"label"`
	Payload string `json:"payload,omitempty"`
}

type KeyboardActionLink  struct {
	Type    string `json:"type"`
	Link string `json:link,omitempty"`

}
type KeyboardButton struct {
	Action KeyboardAction `json:"action"`
	Color  string         `json:"color"`
}
type KeyboardButtonLink struct {
	Action KeyboardActionLink `json:"action"`
	Color  string         `json:"color"`
}


type Keyboard struct {
	OneTime bool               `json:"one_time"`
	Buttons [][]KeyboardButton `json:"buttons"`
	Inline  bool               `json:"inline,omitempty"`
}

type KeyboardLink struct {
	OneTime bool               `json:"one_time"`
	Buttons [][]KeyboardButtonLink `json:"buttons"`
	Inline  bool               `json:"inline,omitempty"`
}


// --- Long Poll структуры (ИСПРАВЛЕНО) ---
type LongPollServer struct {
	Server string `json:"server"`
	Key    string `json:"key"`
	Ts     string `json:"ts"`
	Wait   int    `json:"wait"`
}

type LongPollUpdate struct {
	Type   string                 `json:"type"`
	Object map[string]interface{} `json:"object"`
}

type LongPollResponse struct {
	Ts      string           `json:"ts"`
	Updates []LongPollUpdate `json:"updates"` // теперь это массив структур, а не [][]interface{}
}

type UploadServerResp struct { // для отправки фото
	Server string `json:"server"`
	Photo  string `json:"photo"` // иногда приходит пустым на этом этапе
	Hash   string `json:"hash"`
}

type SavePhotoResp struct {
	Response []struct {
		ID      int `json:"id"`
		OwnerID int `json:"owner_id"`
	} `json:"response"`
}

var TIMESTARTGAME, TIMEENDGAME time.Time

func initValues() {

	var err error
	TIMESTARTGAME, err = time.Parse(time.DateTime, "2026-08-28 19:07:00")

	if err != nil {
		fmt.Println(err)
	}

	TIMEENDGAME = TIMESTARTGAME.Add(time.Hour * 4)
	usersCheckIn[0] = "0"
	usersCheckIn = dbInitUserCheckIN()
}

type Task struct {
	commandName,
	task,
	taskresp,
	taskcomment,
	pod1,
	pod1time,
	pod1comment,
	pod2,
	pod2time,
	pod2comment,
	code string
}

var location, _ = time.LoadLocation("Europe/Moscow")

var userEnterCode = make(map[float64]string)

var commanList []string

var usersCheckIn = make(map[float64]string)

var rules string = "Текст правил" //правила

var intervalTimePod1 int64 = 1

var intervalTimePod2 int64 = intervalTimePod1 * 2

var route []int
var routes [][]int

type Route struct {
	task, taskComment,
	pod1, pod1comment,
	pod2, pod2Comment string
}

var CODES = make(map[int]string)

var COMMENTS_TASK = map[int]string{
	1:  "task1",
	2:  "task2",
	3:  "task3",
	4:  "task4",
	5:  "task5",
	6:  "task6",
	7:  "task7",
	8:  "task8",
	9:  "task9",
	10: "task10",
	11: "task11",
	12: "task12",
	13: "task13",
	14: "task14",
	15: "task15",
	16: "task16",
	17: "task17",
	18: "task18",
	19: "task19",
	20: "task20",
}

var COMMENTS_POD1 = map[int]string{
	1:  "pod11",
	2:  "pod12",
	3:  "pod13",
	4:  "pod14",
	5:  "pod15",
	6:  "pod16",
	7:  "pod17",
	8:  "pod18",
	9:  "pod19",
	10: "pod110",
	11: "pod111",
	12: "pod112",
	13: "pod113",
	14: "pod114",
	15: "pod115",
	16: "pod116",
	17: "pod117",
	18: "pod118",
	19: "pod119",
	20: "pod120",
}

var COMMENTS_POD2 = map[int]string{
	1:  "pod21",
	2:  "pod22",
	3:  "pod23",
	4:  "pod24",
	5:  "pod25",
	6:  "pod26",
	7:  "pod27",
	8:  "pod28",
	9:  "pod29",
	10: "pod210",
	11: "pod211",
	12: "pod212",
	13: "pod213",
	14: "pod214",
	15: "pod215",
	16: "pod216",
	17: "pod217",
	18: "pod218",
	19: "pod219",
	20: "pod220",
}

var countTask = 5

var taskChan = make(chan TaskVK, 100)

type TaskVK struct {
	Type string
	Data json.RawMessage
}
