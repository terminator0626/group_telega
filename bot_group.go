package main

import 
(
	"log"
	"go"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main()
{
	// Прописываем токен бота
	bot, err := tgbotapi.NewBotAPI("7270476740:AAE1f27K-BJbgkkeJWi8_LbiiZG9aY6cH-Q")
	
	if err != nil 
	{
		log.Fatal(err)
	}

	// Установка параметров для получения обновлений из всех чатов
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Получение обновлений из всех чатов 
	updates, err := bot.GetUpdatesChan(u)
	if err != nil
	{
		log.Fatal(err)
	}

	//Обработка полученных обновлений 
	for update := range updates
	{
		// Проверка, что сообщение полученно из группового чата
		if update.Message != nill && update.Message.Chat.Type == "group"
		{
			// Обработка сообщений из групповых чатов, так же анализ текста и выполнение необъодимых действий 

		}
	}

	// Ожидание сигнала для завершения работы 
	// Можно использовать os.Signal для  этого
	// Например:
	// sig := make(chan os.Signal, 1)
	// signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	// <-sig

}