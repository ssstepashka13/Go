package actioninfo

import (
	"fmt"
	"log"
)

// DataParser описывает разбор активности и получение её сводки.
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// Info выводит сводки активностей, пропуская некорректные записи.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("разбор активности: %v", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("расчёт активности: %v", err)
			continue
		}
		fmt.Println(info)
	}
}
