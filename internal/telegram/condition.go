package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

const back = "Назад"
const learnVerbs = "Учить глаголы"
const trainVerbs = "Тренировать глаголы"

const defaultLocation = "default"

var defaultKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(learnVerbs),
		tgbotapi.NewKeyboardButton(trainVerbs),
	),
)

var location = defaultLocation

func (b *Bot) chooseCondition(update tgbotapi.Update) {
	logrus.Debugf("location = %s\n",location)

	if location != defaultLocation {
		switch {
		case location == learnVerbs:
			b.learnVerbsCondition(update)
			break
		case location == trainVerbs:
			b.trainVerbsCondition(update)
			break
		}
		return
	}

	switch {
	case update.Message.Text == learnVerbs:
		location = learnVerbs
		msg := NewSendMessage(update.Message.Chat.ID, learnVerbsKeyboard, learnVerbs, tgbotapi.ModeHTML, b.api)
		msg.SendMessage()
		break
	case update.Message.Text == trainVerbs:
		location = trainVerbs
		msg := NewSendMessage(update.Message.Chat.ID, trainVerbsKeyboard, trainVerbs, tgbotapi.ModeHTML, b.api)
		msg.SendMessage()
		break
	default:
		msg := NewSendMessage(update.Message.Chat.ID, defaultKeyboard, "Сделай выбор", tgbotapi.ModeHTML, b.api)
		msg.SendMessage()
		break
	}
}
