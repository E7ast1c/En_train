package telegram

import (
	"en_train/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

func InitTGBot(botToken string) (*tgbotapi.BotAPI, error) {
	botApi, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}

	botApi.Debug = true
	logrus.Infof("Authorized on account %s\n", botApi.Self.UserName)
	return botApi, nil
}

func (b *Bot) StartTGBot() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.api.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		b.chooseCondition(update)
	}
	return nil
}

type Bot struct {
	api  *tgbotapi.BotAPI
	repo *repository.Repository
}

func NewBot(api *tgbotapi.BotAPI, repo *repository.Repository) *Bot {
	return &Bot{api: api, repo: repo}
}

type SendMessage struct {
	chatId      int64
	replyMarkup tgbotapi.ReplyKeyboardMarkup
	text        string
	mode        string
	api         *tgbotapi.BotAPI
}

func NewSendMessage(chatId int64, replyMarkup tgbotapi.ReplyKeyboardMarkup, text string,
	mode string, api *tgbotapi.BotAPI) SendMessage {

	return SendMessage{
		chatId:      chatId,
		replyMarkup: replyMarkup,
		text:        text,
		mode:        mode,
		api:         api,
	}
}

func (u *SendMessage) SendMessage() {
	config := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:              u.chatId,
			ReplyMarkup:         u.replyMarkup,
		},
		Text:      u.text,
		ParseMode: u.mode,
	}

	send, err := u.api.Send(config)
	if err != nil {
		return
	}
	logrus.Infof("Send message %+v", send)
}
