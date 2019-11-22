// main
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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

			button1 := tgbotapi.NewKeyboardButton("https://github.com/ImDaeseong")
			button2 := tgbotapi.NewKeyboardButton("https://github.com/ImDaeseong/DaeseongBot")

			//한줄에
			//Markup := tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(button1, button2))

			//한줄에
			//Markup := tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{button1, button2})

			//여러줄에
			Markup := tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(button1), tgbotapi.NewKeyboardButtonRow(button2))

			sMsg := "github.com/ImDaeseong 정보를 확인하세요."
			sSendMsg := tgbotapi.NewMessage(chatid, sMsg)
			sSendMsg.ReplyMarkup = Markup
			_, err := bot.Send(sSendMsg)
			if err != nil {
				log.Print(err.Error())
			}

		} else if update.Message.Command() == "/link" || update.Message.Text == "/link" {

			link1 := tgbotapi.NewInlineKeyboardButtonURL("Daeseong", "https://github.com/ImDaeseong")
			link2 := tgbotapi.NewInlineKeyboardButtonURL("DaeseongBot", "https://github.com/ImDaeseong/DaeseongBot")

			//한줄에
			Markup := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(link1, link2))

			//여러줄에
			//Markup := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(link1), tgbotapi.NewInlineKeyboardRow(link2))

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

		} else if update.Message.Text == "/send" {

			fmt.Println("테스트 완료")

			/*
				sSend := tgbotapi.NewAudioUpload(chatid, "E:\\GoApp\\src\\audio.mp3")
				_, err := bot.Send(sSend)
				if err != nil {
					log.Print(err.Error())
				}
			*/

			/*
				sSend := tgbotapi.NewPhotoUpload(chatid, "E:\\GoApp\\src\\photo.jpg")
				_, err := bot.Send(sSend)
				if err != nil {
					log.Print(err.Error())
				}
			*/

			/*
				sSend := tgbotapi.NewStickerUpload(chatid, "E:\\GoApp\\src\\sticker.tgs")
				_, err := bot.Send(sSend)
				if err != nil {
					log.Print(err.Error())
				}
			*/

			/*
				sSend := tgbotapi.NewStickerShare(chatid, "CAADBQADYwUAAmQK4AW3jYFjDvykkAI")
				_, err := bot.Send(sSend)
				if err != nil {
					log.Print(err.Error())
				}
			*/

			/*
				sSend := tgbotapi.NewVideoUpload(chatid, "E:\\GoApp\\src\\video.mp4")
				_, err := bot.Send(sSend)
				if err != nil {
					log.Print(err.Error())
				}
			*/

			/*
				sSend := tgbotapi.NewVoiceUpload(chatid, "E:\\GoApp\\src\\voice.dat")
				_, err := bot.Send(sSend)
				if err != nil {
					log.Print(err.Error())
				}
			*/

			/*
				sSend := tgbotapi.NewDocumentUpload(chatid, "E:\\GoApp\\src\\document.doc")
				_, err := bot.Send(sSend)
				if err != nil {
					log.Print(err.Error())
				}
			*/

		} else if update.Message.Audio != nil {

			url, err := bot.GetFileDirectURL(update.Message.Audio.FileID)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println(url)

			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			file, err := os.OpenFile("audio.mp3", os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err.Error())
				resp.Body.Close()
				continue
			}
			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err.Error())
				resp.Body.Close()
				file.Close()
				continue
			}
			file.Write(bytes)
			file.Close()

		} else if update.Message.Photo != nil {

			photos := *update.Message.Photo
			photoId := photos[1].FileID
			url, err := bot.GetFileDirectURL(photoId)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println(url)

			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			file, err := os.OpenFile("photo.jpg", os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err.Error())
				resp.Body.Close()
				continue
			}
			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err.Error())
				resp.Body.Close()
				file.Close()
				continue
			}
			file.Write(bytes)
			file.Close()

		} else if update.Message.Document != nil {

			url, err := bot.GetFileDirectURL(update.Message.Document.FileID)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println(url)

			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			file, err := os.OpenFile("document.doc", os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err.Error())
				resp.Body.Close()
				continue
			}
			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err.Error())
				resp.Body.Close()
				file.Close()
				continue
			}
			file.Write(bytes)
			file.Close()

		} else if update.Message.Sticker != nil {

			url, err := bot.GetFileDirectURL(update.Message.Sticker.FileID)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println(url)

			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			file, err := os.OpenFile("sticker.tgs", os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err.Error())
				resp.Body.Close()
				continue
			}
			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err.Error())
				resp.Body.Close()
				file.Close()
				continue
			}
			file.Write(bytes)
			file.Close()

		} else if update.Message.Video != nil {

			url, err := bot.GetFileDirectURL(update.Message.Video.FileID)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println(url)

			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			file, err := os.OpenFile("video.mp4", os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err.Error())
				resp.Body.Close()
				continue
			}
			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err.Error())
				resp.Body.Close()
				file.Close()
				continue
			}
			file.Write(bytes)
			file.Close()

		} else if update.Message.Voice != nil {

			url, err := bot.GetFileDirectURL(update.Message.Voice.FileID)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			//fmt.Println(url)

			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			file, err := os.OpenFile("voice.dat", os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err.Error())
				resp.Body.Close()
				continue
			}
			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err.Error())
				resp.Body.Close()
				file.Close()
				continue
			}
			file.Write(bytes)
			file.Close()

		} else if update.Message.Location != nil {

			fmt.Println(update.Message.Location.Latitude)
			fmt.Println(update.Message.Location.Longitude)

		} else if update.CallbackQuery != nil {

			fmt.Println(update.CallbackQuery.Data)

		} else {

			sMsg := "정의된 메시지 내용이 아닙니다."
			sSendMsg := tgbotapi.NewMessage(chatid, sMsg)
			_, err := bot.Send(sSendMsg)
			if err != nil {
				log.Print(err.Error())
			}

		}

	}

}
