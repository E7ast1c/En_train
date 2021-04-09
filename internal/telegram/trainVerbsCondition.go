package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

const (
	randomVerb         = "Случайный глагол"
	infinitiveVerb     = "Infinitive"
	pastSimpleVerb     = "Past Simple"
	pastParticipleVerb = "Past Participle"
)

var trainVerbsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(randomVerb),
		tgbotapi.NewKeyboardButton(infinitiveVerb),
		tgbotapi.NewKeyboardButton(pastSimpleVerb),
		tgbotapi.NewKeyboardButton(pastParticipleVerb),
		tgbotapi.NewKeyboardButton(back),
	))

func (b *Bot) trainVerbsCondition(update tgbotapi.Update) {
	switch update.Message.Text {
	case nextVerb:
		rv, err := b.repo.GetRandomVerb()
		if err != nil {
			logrus.Error(err)
			return
		}

		text := fmt.Sprintf("<b>Infinitive:</b> %s\n"+
			"<b>PastTense:</b> %s\n"+
			"<b>PastParticiple:</b> %s\n"+
			"<b>Translate:</b> %s\n",
			rv.Infinitive, rv.PastTense, rv.PastParticiple, rv.Translate)
		msg := NewSendMessage(update.Message.Chat.ID, learnVerbsKeyboard, text, tgbotapi.ModeHTML, b.api)
		msg.SendMessage()
	case infinitiveVerb:
	case pastSimpleVerb:
	case pastParticipleVerb:
	case back:
		location = defaultLocation
		msg := NewSendMessage(update.Message.Chat.ID, defaultKeyboard,
			"Возвращаемся на главный экран", tgbotapi.ModeHTML, b.api)
		msg.SendMessage()
	default:
	}
}
