// Package actioninfo предоставляет интерфейс и функции для обработки
// и вывода сводной информации о различных видах физической активности
package actioninfo

import (
	"fmt"
	"log"
)

// DataParser описывает поведение типов, способных парсить строку с данными
// об активности и формировать сводный отчёт
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info обрабатывает набор строк с данными об активностях и выводит информацию о каждой из них
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Println(err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(info)
	}
}
