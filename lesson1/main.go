package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	filePth := os.Args[1]
	var fileName, fileExt string

	// Напишите код, который выведет следующее
	// filename: <name>
	// extension: <extension>

	// Подсказка. Возможно вам понадобится функция strings.LastIndex
	// Для проверки своего решения используйте функции filepath.Base() filepath.Ext(
	// ) Они могут помочь для проверки решения

	// Ваш код здесь

	// разбиваем путь на части
	// массив с разделителем "/"
	filePathArray := strings.Split(filePth, "/")
	filenameArray := strings.Split(filePathArray[len(filePathArray)-1], ".")
	if len(filenameArray) > 1 {
		fileExt = filenameArray[len(filenameArray)-1]
	}
	fileName = strings.Split(filePathArray[len(filePathArray)-1], fmt.Sprintf(`.%s`, fileExt))[0]

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
