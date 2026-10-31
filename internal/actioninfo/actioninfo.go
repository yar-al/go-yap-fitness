package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	ActionInfo() (string, error)
	Parse(datastring string) (err error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Print("Ошибка парсинга: ", err)
		}
		actInfo, err := dp.ActionInfo()
		if err != nil {
			log.Print("Ошибка формирования информации о действии: ", err)
		}
		fmt.Print(actInfo)
	}
}
