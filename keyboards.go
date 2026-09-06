package main

import "fmt"

func enterPass() *Keyboard {

	k := &Keyboard{

		OneTime: false,
		Buttons: [][]KeyboardButton{

			{
				{
					Action: KeyboardAction{Type: "text", Label: "Отмена"},
					Color:  "default",
				},
			},
		},
	}

	return k
}
 
func auth() *Keyboard {

	k := &Keyboard{

		OneTime: false,
		Buttons: [][]KeyboardButton{
			{
				{
					Action: KeyboardAction{Type: "text", Label: "Авторизация"},
					Color:  "primary",
				},
			},
			{
				{
					Action: KeyboardAction{Type: "text", Label: "Куда я попал и что делать?"},
					Color:  "default",
				},
			},
		},
	}

	return k

}


func gameEnteringCode () *Keyboard {
	return &Keyboard{

		OneTime: false,
		Buttons: [][]KeyboardButton{
			{
				{
					Action: KeyboardAction{Type: "text", Label: "Отмена"},
					Color:  "primary",
				},
			},
		},
	
	}

}
func game() *Keyboard {
	return &Keyboard{

		OneTime: false,
		Buttons: [][]KeyboardButton{
			{
				{
					Action: KeyboardAction{Type: "text", Label: "Начать игру"},
					Color:  "primary",
				},
			},
			// {
			// 	{
			// 		Action: KeyboardAction{Type: "text", Label: "О команде"},
			// 		Color:  "default",
			// 	},
			// },
			{
				{
					Action: KeyboardAction{Type: "text", Label: "Правила"},
					Color:  "default",
				},
			},

			{
				{
					Action: KeyboardAction{Type: "text", Label: "Выход"},
					Color:  "default",
				},
			},
		},
	}

}

func gameAfterStart() *Keyboard {
	return &Keyboard{

		OneTime: false,
		Buttons: [][]KeyboardButton{
			{
				{
					Action: KeyboardAction{Type: "text", Label: "Попросить подсказку"},
					Color:  "primary",
				},
			},
			{
				{
					Action: KeyboardAction{Type: "text", Label: "Статус"},
					Color:  "default",
				},
			},
			{
				{
					Action: KeyboardAction{Type: "text", Label: "Повторить вопрос"},
					Color:  "default",
				},
			},
			{
				{
					Action: KeyboardAction{Type: "text", Label: "Ввести код"},
					Color:  "default",
				},
			},
			{
				{
					Action: KeyboardAction{Type: "text", Label: "Выход"},
					Color:  "default",
				},
			},
		},
	}

}

func commandList(list []string) *Keyboard {

	var b [][]KeyboardButton

	fmt.Println(len(list))

	if len(list) > 15 {

		for i := 0; i < len(list); i++ {

			if i+2 >= len(list) {
				b = append(b, []KeyboardButton{

					{
						Action: KeyboardAction{Type: "text", Label: list[i]},
						Color:  "primary",
					},
					{
						Action: KeyboardAction{Type: "text", Label: list[i+1]},
						Color:  "primary",
					},
				},
				)
				i += 1

			} else {
				b = append(b, []KeyboardButton{

					{
						Action: KeyboardAction{Type: "text", Label: list[i]},
						Color:  "primary",
					},
					{
						Action: KeyboardAction{Type: "text", Label: list[i+1]},
						Color:  "primary",
					},
					{
						Action: KeyboardAction{Type: "text", Label: list[i+2]},
						Color:  "primary",
					},
				},
				)
				i += 2
			}
		}
	}

	if len(list) > 10 && len(list) <= 15 {
		for i := 0; i < len(list); i++ {

			if i+1 >= len(list) {
				b = append(b, []KeyboardButton{

					{
						Action: KeyboardAction{Type: "text", Label: list[i]},
						Color:  "primary",
					},
				},
				)
				i += 1

			} else {
				b = append(b, []KeyboardButton{

					{
						Action: KeyboardAction{Type: "text", Label: list[i]},
						Color:  "primary",
					},
					{
						Action: KeyboardAction{Type: "text", Label: list[i+1]},
						Color:  "primary",
					},
				},
				)
				i += 1
			}
		}
	}

	b = append(b, []KeyboardButton{
		{
			Action: KeyboardAction{Type: "text", Label: "Назад"},
			Color:  "primary",
		},
	})

	return &Keyboard{
		OneTime: false,
		Buttons: b,
	}

}


func gameOver() *Keyboard {

	return &Keyboard{

		OneTime: false,
		Buttons: [][]KeyboardButton{
			{
				{
					Action: KeyboardAction{Type: "text", Label: "Отчет"},
					Color:  "primary",
				},
			},
			// {
			// 	{
			// 		Action: KeyboardAction{Type: "text", Label: "О команде"},
			// 		Color:  "default",
			// 	},
			// },

			{
				{
					Action: KeyboardAction{Type: "text", Label: "Выход"},
					Color:  "default",
				},
			},
		},
	}


}
//Шаблон
// keyboard := &Keyboard{
// 	OneTime: false,
// 	Buttons: [][]KeyboardButton{
// 		{
// 			{
// 				Action: KeyboardAction{Type: "text", Label: ""},
// 				Color:  "primary",
// 			},
// 			{
// 				Action: KeyboardAction{Type: "text", Label: "Кнопка 2"},
// 				Color:  "positive",
// 			},
// 		},
// 		{
// 			{
// 				Action: KeyboardAction{Type: "text", Label: "Ещё вариант"},
// 				Color:  "default",
// 			},
// 		},
// 		{
// 			{Action: KeyboardAction{Type: "text", Label: "Ссылка"},
// 				Color: "default",
// 			},
// 		},
// 	},
// }
