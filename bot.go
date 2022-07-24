package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func bot() {
	updates, bot := authenticate()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		args := update.Message.CommandArguments()
		cmd := update.Message.Command()
		update.Message.CommandWithAt()

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.Command() == "start" {
			answerMessage(update, bot, update.Message.MessageID, "Mogu Mogu")
			continue
		}

		if isSign(cmd) {
			if cmd == "sign5" {
				answerMessage(update, bot, update.Message.MessageID, "This image is currently not supported as I am still "+
					"working on signs that are not at a 90° angle")
				continue
			}
			if cmd == "sign7" {
				answerMessage(update, bot, update.Message.MessageID, "This image is currently not supported as I I am still working on images with major intrusions into the text area (https://files.catbox.moe/79zoj3.png)")
				continue
			}
			imageBytes := selectImage(cmd, args)
			sendPhoto(update, bot, imageBytes)
			continue
		}

	}
}

func answerMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI, msgID int, text string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ParseMode = "html"
	msg.DisableWebPagePreview = true
	msg.ReplyToMessageID = msgID
	_, err := bot.Send(msg)
	if err != nil {
		PrintErr("Failed to send message", err, false)
	}

}

func sendMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI, text string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ParseMode = "html"
	msg.DisableWebPagePreview = true
	_, err := bot.Send(msg)
	if err != nil {
		PrintErr("Failed to send message", err, false)
	}
}

func sendPhoto(update tgbotapi.Update, bot *tgbotapi.BotAPI, data []byte) {
	pic := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileBytes{Bytes: data})
	_, err := bot.Send(pic)
	if err != nil {
		PrintErr("Failed to send photo", err, false)
	}
}

func authenticate() (tgbotapi.UpdatesChannel, *tgbotapi.BotAPI) {

	APItoken := "5267583211:AAHKKxasVDTKFrmodf_yye2rF6BXQ-EeXYI"

	b, err := tgbotapi.NewBotAPI(APItoken)

	if err != nil {
		PrintErr("Failed to connect Bot to Telegram", err, true)
	}

	b.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	log.Printf("Authorized on account %s", b.Self.UserName)

	return b.GetUpdatesChan(u), b
}
