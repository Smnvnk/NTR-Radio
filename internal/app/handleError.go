package app

import "log"

// Обработчик ошибок
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
