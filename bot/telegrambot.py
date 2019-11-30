import telegram
from telegram.ext import Updater, CommandHandler
from telegram import InlineKeyboardButton, InlineKeyboardMarkup, KeyboardButton, ReplyKeyboardMarkup


TOKEN = 'TOKEN'
bot = telegram.Bot(token=TOKEN)


def sendAll(chatid):
    bot.sendMessage(chat_id=chatid, text="테스트 전달")
    bot.sendMessage(chat_id=chatid, text='https://github.com/ImDaeseong')
    bot.send_photo(chat_id=chatid, photo=open('E:\\a\\file_60.jpg', 'rb'))
    bot.send_photo(chat_id=chatid, photo='https://avatars0.githubusercontent.com/u/10001221?s=460&v=4')
    bot.send_audio(chat_id=chatid, audio=open('E:\\a\\file_64.mp3', 'rb'))
    bot.send_sticker(chat_id=chatid, sticker=open('E:\\a\\file_62.tgs', 'rb'))
    bot.send_document(chat_id=chatid, document=open('E:\\a\\file_59.doc', 'rb'))
    bot.send_video(chat_id=chatid, video=open('E:\\a\\file_51.mp4', 'rb'))
    bot.sendVoice(chat_id=chatid, voice=open('E:\\a\\file_66.oga', 'rb'))
    bot.send_sticker(chat_id=chatid, sticker='CAADBQADYwUAAmQK4AW3jYFjDvykkAI')


def start(chatid):
    sMsg = '/help -> 나의 이미지\n/keyboard -> 나의 페이지\n/link -> 나의 링크 페이지'
    bot.sendMessage(chat_id=chatid, text=sMsg)


def keyboard(chatid, sMsg):
    # 2줄에 링크
    """
    keyboard = [
        [KeyboardButton("https://github.com/ImDaeseong")],
        [KeyboardButton("https://github.com/ImDaeseong/DaeseongBot")]
    ]
    """

    # 1줄에 링크
    keyboard = [
        [KeyboardButton("https://github.com/ImDaeseong"), KeyboardButton("https://github.com/ImDaeseong/DaeseongBot")]
    ]
    reply_markup = ReplyKeyboardMarkup(keyboard, resize_keyboard=True)
    bot.sendMessage(chat_id=chatid, text=sMsg, reply_markup=reply_markup)


def link(chatid, sMsg):
    # 2줄에 링크
    """
    InlineKeyboard = [
        [InlineKeyboardButton('Daeseong', url='https://github.com/ImDaeseong')],
        [InlineKeyboardButton('DaeseongBot',  url='https://github.com/ImDaeseong/DaeseongBot')]
    ]
    """

    # 1줄에 링크
    InlineKeyboard = [
        [InlineKeyboardButton('Daeseong', url='https://github.com/ImDaeseong'),
         InlineKeyboardButton('DaeseongBot', url='https://github.com/ImDaeseong/DaeseongBot')]
    ]

    reply_markup = InlineKeyboardMarkup(InlineKeyboard)
    bot.sendMessage(chat_id=chatid, text=sMsg, reply_markup=reply_markup)


def help(chatid):
    bot.send_photo(chat_id=chatid, photo='https://avatars0.githubusercontent.com/u/10001221?s=460&v=4')


def Getfile_id(ary, findkey):
    file_id = ''

    find_1 = ary.find(findkey)
    if find_1 != -1:
        ary1 = ary[find_1:]
        # print(ary1)

        find_2 = ary1.find('file_id')
        if find_2 != -1:
            ary2 = ary1[find_2 + 7:]

            find_3 = ary2.find(',')
            if find_3 != -1:
                ary3 = ary2[:find_3]
                # print(ary3)

                file_id = str(ary3).replace("'", "").replace(":", "").strip()
                # print(file_id)
    return file_id


def GetPhotofile_id(ary):
    file_id = ''

    find_1 = ary.find('photo')
    if find_1 != -1:
        ary1 = ary[find_1:]
        # print(ary1)

        find_2 = ary1.find('[')
        find_3 = ary1.find(']')
        if find_2 != -1 and find_3 != -1:
            ary2 = ary1[find_2 + 1:find_3]
            # print(ary2)

            find_3 = ary2.find('file_id')
            if find_3 != -1:
                ary3 = ary2[find_3 + 7:]

                find_4 = ary3.find(',')
                if find_4 != -1:
                    ary4 = ary3[:find_4]
                    # print(ary4)

                    file_id = str(ary4).replace("'", "").replace(":", "").strip()
                    # print(file_id)
    return file_id


def otherType(chatid, message):
    ary = str(message)
    # print('chatid:%s message:%s' % (chatid, ary))

    file_id_sticker = Getfile_id(ary, 'sticker')
    if file_id_sticker != '':
        file = bot.get_file(file_id_sticker)
        filename = 'E:\\%s.tgs' % file_id_sticker
        print(filename)
        file.download(filename)

    file_id_video = Getfile_id(ary, 'video')
    if file_id_video != '':
        file = bot.get_file(file_id_video)
        filename = 'E:\\%s.mp4' % file_id_video
        print(filename)
        file.download(filename)

    file_id_audio = Getfile_id(ary, 'audio')
    if file_id_audio != '':
        file = bot.get_file(file_id_audio)
        filename = 'E:\\%s.mp3' % file_id_audio
        print(filename)
        file.download(filename)

    file_id_document = Getfile_id(ary, 'document')
    if file_id_document != '':
        file = bot.get_file(file_id_document)
        filename = 'E:\\%s.doc' % file_id_document
        print(filename)
        file.download(filename)

    file_id_voice = Getfile_id(ary, 'voice')
    if file_id_voice != '':
        file = bot.get_file(file_id_voice)
        filename = 'E:\\%s.oga' % file_id_voice
        print(filename)
        file.download(filename)

    file_id_Photo = GetPhotofile_id(ary)
    if file_id_Photo != '':
        file = bot.get_file(file_id_Photo)
        filename = 'E:\\%s.jpg' % file_id_Photo
        print(filename)
        file.download(filename)


def botInit():
    print('get_me:%s' % bot.get_me())

    updater = Updater(token=TOKEN, use_context=True)
    dispatcher = updater.dispatcher

    start_handler = CommandHandler('start', start)
    keyboard_handler = CommandHandler('keyboard', keyboard)
    link_handler = CommandHandler('link', link)
    help_handler = CommandHandler('help', help)
    dispatcher.add_handler(start_handler)
    dispatcher.add_handler(keyboard_handler)
    dispatcher.add_handler(link_handler)
    dispatcher.add_handler(help_handler)

    try:
        update_id = bot.getUpdates()[0].update_id
    except IndexError:
        update_id = None
    print('update_id:%s' % update_id)

    while True:
        for update in bot.getUpdates(offset=update_id, timeout=60):

            chat_id = update.message.chat_id
            # firstname = update.message.chat.first_name
            # last_name = update.message.chat.last_name
            # print('chat_id:%s' % chat_id)
            # print('first_name:%s' % firstname)
            # print('last_name:%s' % last_name)
            # print('message:%s' % update.message)
            # print('message_id:%s' % update.message.message_id)

            try:
                if update.message.text:
                    if update.message.text == '/start':
                        start(chat_id)
                    elif update.message.text == '/keyboard':
                        keyboard(chat_id, update.message.text)
                    elif update.message.text == '/link':
                        link(chat_id, update.message.text)
                    elif update.message.text == '/help':
                        help(chat_id)
                    elif update.message.text == 'send':
                        sendAll(chat_id)
                    else:
                        bot.sendMessage(chat_id=chat_id, text=update.message.text)
                else:
                    otherType(chat_id, update.message)

            except Exception as e:
                print('error:\n%s' % e)
            finally:
                update_id = update.update_id + 1


if __name__ == '__main__':
    botInit()
