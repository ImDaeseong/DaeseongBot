// main
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

//https://web.telegram.org/#/im?p=@BotFather
var bot *tgbotapi.BotAPI

func getfileNameUrl(sUrl string) string {

	nIndex := strings.LastIndex(sUrl, "/")
	if nIndex == -1 {
		return ""
	}

	return string(sUrl[nIndex+1:])
}

func AudioDownload(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {

	url, err := bot.GetFileDirectURL(message.Audio.FileID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	filename := getfileNameUrl(url)
	if filename == "" {
		return
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
		resp.Body.Close()
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		resp.Body.Close()
		file.Close()
		return
	}
	file.Write(bytes)
	file.Close()
}

func PhotoDownload(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {

	photos := *message.Photo
	photoId := photos[1].FileID
	url, err := bot.GetFileDirectURL(photoId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	filename := getfileNameUrl(url)
	if filename == "" {
		return
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
		resp.Body.Close()
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		resp.Body.Close()
		file.Close()
		return
	}
	file.Write(bytes)
	file.Close()
}

func DocumentDownload(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {

	url, err := bot.GetFileDirectURL(message.Document.FileID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	filename := getfileNameUrl(url)
	if filename == "" {
		return
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
		resp.Body.Close()
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		resp.Body.Close()
		file.Close()
		return
	}
	file.Write(bytes)
	file.Close()
}

func StickerDownload(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {

	url, err := bot.GetFileDirectURL(message.Sticker.FileID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	filename := getfileNameUrl(url)
	if filename == "" {
		return
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
		resp.Body.Close()
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		resp.Body.Close()
		file.Close()
		return
	}
	file.Write(bytes)
	file.Close()
}

func VideoDownload(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {

	url, err := bot.GetFileDirectURL(message.Video.FileID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	filename := getfileNameUrl(url)
	if filename == "" {
		return
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
		resp.Body.Close()
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		resp.Body.Close()
		file.Close()
		return
	}
	file.Write(bytes)
	file.Close()
}

func VoiceDownload(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {

	url, err := bot.GetFileDirectURL(message.Voice.FileID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	filename := getfileNameUrl(url)
	if filename == "" {
		return
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
		resp.Body.Close()
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		resp.Body.Close()
		file.Close()
		return
	}
	file.Write(bytes)
	file.Close()
}

func SendAudio(bot *tgbotapi.BotAPI, chatid int64, sPath string) {
	sSend := tgbotapi.NewAudioUpload(chatid, sPath)
	_, err := bot.Send(sSend)
	if err != nil {
		log.Print(err.Error())
	}
}

func SendPhoto(bot *tgbotapi.BotAPI, chatid int64, sPath string) {
	sSend := tgbotapi.NewPhotoUpload(chatid, sPath)
	_, err := bot.Send(sSend)
	if err != nil {
		log.Print(err.Error())
	}
}

func SendSticker(bot *tgbotapi.BotAPI, chatid int64, sPath string) {
	sSend := tgbotapi.NewStickerUpload(chatid, sPath)
	_, err := bot.Send(sSend)
	if err != nil {
		log.Print(err.Error())
	}
}

func SendStickerShare(bot *tgbotapi.BotAPI, chatid int64) {
	sSend := tgbotapi.NewStickerShare(chatid, "CAADBQADYwUAAmQK4AW3jYFjDvykkAI")
	_, err := bot.Send(sSend)
	if err != nil {
		log.Print(err.Error())
	}
}

func SendVideo(bot *tgbotapi.BotAPI, chatid int64, sPath string) {
	sSend := tgbotapi.NewVideoUpload(chatid, sPath)
	_, err := bot.Send(sSend)
	if err != nil {
		log.Print(err.Error())
	}
}

func SendVoice(bot *tgbotapi.BotAPI, chatid int64, sPath string) {
	sSend := tgbotapi.NewVoiceUpload(chatid, sPath)
	_, err := bot.Send(sSend)
	if err != nil {
		log.Print(err.Error())
	}
}

func SendDocument(bot *tgbotapi.BotAPI, chatid int64, sPath string) {
	sSend := tgbotapi.NewDocumentUpload(chatid, sPath)
	_, err := bot.Send(sSend)
	if err != nil {
		log.Print(err.Error())
	}
}

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

			go func() {
				sMsg := fmt.Sprintf("%s\n%s\n%s\n%s", "/help -> 나의 이미지", "/keyboard -> 나의 페이지", "/link -> 나의 링크 페이지", "/close -> 키보드 제거")
				sSendMsg := tgbotapi.NewMessage(chatid, sMsg)
				_, err := bot.Send(sSendMsg)
				if err != nil {
					log.Print(err.Error())
				}
			}()

		} else if update.Message.Command() == "/keyboard" || update.Message.Text == "/keyboard" {

			go func() {
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
			}()

		} else if update.Message.Command() == "/link" || update.Message.Text == "/link" {

			go func() {
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
			}()

		} else if update.Message.Command() == "/help" || update.Message.Text == "/help" {

			go func() {
				imgUrl := "https://avatars0.githubusercontent.com/u/10001221?s=460&v=4"
				sSendMsg := tgbotapi.NewPhotoShare(chatid, imgUrl)
				_, err := bot.Send(sSendMsg)
				if err != nil {
					log.Print(err.Error())
				}
			}()

		} else if update.Message.Text == "/close" {

			go func() {
				sMsg := "키보드를 삭제합니다."
				sSendMsg := tgbotapi.NewMessage(chatid, sMsg)
				sSendMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				_, err := bot.Send(sSendMsg)
				if err != nil {
					log.Print(err.Error())
				}
			}()

		} else if update.Message.Text == "/send" {

			go SendAudio(bot, chatid, "D:\\DaeseongGolang\\src\\file_42.mp3")

			go SendPhoto(bot, chatid, "D:\\DaeseongGolang\\src\\file_43.jpg")

			go SendSticker(bot, chatid, "D:\\DaeseongGolang\\src\\file_41.tgs")

			go SendStickerShare(bot, chatid)

			go SendVideo(bot, chatid, "D:\\DaeseongGolang\\src\\file_47.mp4")

			go SendVoice(bot, chatid, "D:\\DaeseongGolang\\src\\file_45.dat")

			go SendDocument(bot, chatid, "D:\\DaeseongGolang\\src\\file_48.doc")

		} else if update.Message.Audio != nil {

			go AudioDownload(bot, update.Message)

		} else if update.Message.Photo != nil {

			go PhotoDownload(bot, update.Message)

		} else if update.Message.Document != nil {

			go DocumentDownload(bot, update.Message)

		} else if update.Message.Sticker != nil {

			go StickerDownload(bot, update.Message)

		} else if update.Message.Video != nil {

			go VideoDownload(bot, update.Message)

		} else if update.Message.Voice != nil {

			go VoiceDownload(bot, update.Message)

		} else if update.Message.Location != nil {

			go func() {
				fmt.Println(update.Message.Location.Latitude)
				fmt.Println(update.Message.Location.Longitude)
			}()

		} else if update.CallbackQuery != nil {

			fmt.Println(update.CallbackQuery.Data)

		} else {

			go func() {
				sMsg := "정의된 메시지 내용이 아닙니다."
				sSendMsg := tgbotapi.NewMessage(chatid, sMsg)
				_, err := bot.Send(sSendMsg)
				if err != nil {
					log.Print(err.Error())
				}
			}()

		}

	}

}
