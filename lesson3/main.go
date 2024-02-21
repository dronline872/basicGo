package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

type URL struct {
	Name string
	Date time.Time
	Tags string
	Link string
}

type URLList struct {
	urls []URL
}

func (u *URLList) AddURL(url URL) {
	u.urls = append(u.urls, url)
}

func (u *URLList) GetURLs() []URL {
	return u.urls
}

func (u *URLList) RemoveURL(name string) {
	for i, url := range u.urls {
		if url.Name == name {
			u.urls = append(u.urls[:i], u.urls[i+1:]...)
			break
		}
	}
}

func main() {
	defer func() {
		// Завершаем работу с клавиатурой при выходе из функции
		_ = keyboard.Close()
	}()

	fmt.Println("Программа для добавления url в список")
	fmt.Println("Для выхода и приложения нажмите Esc")
	fmt.Println("Для добавления нового url нажмите a")
	fmt.Println("Для вывода списка добавленных url нажмите l")
	fmt.Println("Для удаления url из списка нажмите r")
	URLList := URLList{}
OuterLoop:
	for {
		// Подключаем отслеживание нажатия клавиш
		if err := keyboard.Open(); err != nil {
			log.Fatal(err)
		}

		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'a':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}

			// Добавление нового url в список хранения
			fmt.Println("Введите новую запись в формате <url описание теги>")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			args := strings.Fields(text)
			if len(args) < 3 {
				fmt.Println("Введите правильный аргументы в формате url описание теги")
				continue OuterLoop
			}

			// Напишите свой код здесь
			URLList.AddURL(URL{
				Name: args[1],
				Date: time.Now(),
				Tags: args[2],
				Link: args[0],
			})
			fmt.Println("URL добавлен")
		case 'l':
			// Вывод списка добавленных url. Выведите количество добавленных url и список с данными url
			// Вывод в формате
			// Имя: <Описание>
			// URL: <url>
			// Теги: <Теги>
			// Дата: <дата>

			// Напишите свой код здесь
			urls := URLList.GetURLs()
			fmt.Println("Количество добавленных url: ", len(urls))
			for _, url := range urls {
				fmt.Println("Имя: ", url.Name)
				fmt.Println("URL: ", url.Link)
				fmt.Println("Теги: ", url.Tags)
				fmt.Println("Дата: ", url.Date.Format("2006-01-02 15:04:05"))
				fmt.Println("-----------------")
			}
		case 'r':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}
			// Удаление url из списка хранения
			fmt.Println("Введите имя ссылки, которое нужно удалить")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			_ = text

			// Напишите свой код здесь
			URLList.RemoveURL(strings.Split(text, "\n")[0])
		default:
			// Если нажата Esc выходим из приложения
			if key == keyboard.KeyEsc {
				return
			}
		}
	}
}
