// main
package main

import (
	"fmt"
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

//https://web.telegram.org/#/im?p=@BotFather
var bot *tgbotapi.BotAPI

func main() {

	bot, err := tgbotapi.NewBotAPI("TOKEN")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	sMy := fmt.Sprintf("[나의아이디:%d] 봇이름:%s", bot.Self.ID, bot.Self.UserName)
	fmt.Println(sMy)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatid := update.Message.Chat.ID
		//sUserName := update.Message.From.UserName
		//sText := update.Message.Text
		//firstname := update.Message.From.FirstName
		//lasttname := update.Message.From.LastName

		if update.Message.Command() == "/start" || update.Message.Text == "/start" {

			sMsg := fmt.Sprintf("%s\n%s\n%s", "/help -> 나의 이미지", "/keyboard -> 나의 페이지", "/link -> 나의 링크 페이지")
			sSendMsg := tgbotapi.NewMessage(chatid, sMsg)
			_, err := bot.Send(sSendMsg)
			if err != nil {
				log.Print(err.Error())
			}

		} else if update.Message.Command() == "/keyboard" || update.Message.Text == "/keyboard" {

			opt1 := tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("ImDaeseong"), tgbotapi.NewKeyboardButton("https://github.com/ImDaeseong"))
			opt2 := tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("DaeseongBot"), tgbotapi.NewKeyboardButton("https://github.com/ImDaeseong/DaeseongBot"))
			options := tgbotapi.NewReplyKeyboard(opt1, opt2)

			sMsg := "github.com/ImDaeseong 정보를 확인하세요."
			sSendMsg := tgbotapi.NewMessage(chatid, sMsg)
			sSendMsg.ReplyMarkup = options
			_, err := bot.Send(sSendMsg)
			if err != nil {
				log.Print(err.Error())
			}

		} else if update.Message.Command() == "/link" || update.Message.Text == "/link" {

			link1 := tgbotapi.NewInlineKeyboardButtonURL("Daeseong", "https://github.com/ImDaeseong")
			link2 := tgbotapi.NewInlineKeyboardButtonURL("DaeseongBot", "https://github.com/ImDaeseong/DaeseongBot")
			Markup := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(link1, link2))

			sMsg := "나의 링크 정보를 확인하세요."
			sSendMsg := tgbotapi.NewMessage(chatid, sMsg)
			sSendMsg.ReplyMarkup = Markup
			_, err := bot.Send(sSendMsg)
			if err != nil {
				log.Print(err.Error())
			}

		} else if update.Message.Command() == "/help" || update.Message.Text == "/help" {

			imgUrl := "https://avatars0.githubusercontent.com/u/10001221?s=460&v=4"
			sSendMsg := tgbotapi.NewPhotoShare(chatid, imgUrl)
			_, err := bot.Send(sSendMsg)
			if err != nil {
				log.Print(err.Error())
			}

		} else if update.Message.Audio != nil {

		} else if update.Message.Photo != nil {

		} else if update.Message.Document != nil {

		} else if update.Message.Sticker != nil {

		} else if update.Message.Video != nil {

		} else if update.Message.Voice != nil {

		} else if update.Message.Location != nil {

			fmt.Println(update.Message.Chat.ID)
			fmt.Println(update.Message.Location)

		} else {

		}

	}

}
