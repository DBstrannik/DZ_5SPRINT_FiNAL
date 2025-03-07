package actioninfo

import "fmt"

// DataParser — интерфейс для парсинга данных и формирования информации.
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// Info обрабатывает набор данных и выводит информацию о тренировках или прогулках.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			fmt.Println("Ошибка парсинга:", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Println("Ошибка формирования информации:", err)
			continue
		}
		fmt.Println(info)
	}
}
