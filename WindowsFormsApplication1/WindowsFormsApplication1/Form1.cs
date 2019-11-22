using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;
using System.IO;
using System.Threading.Tasks;

using Telegram.Bot;
using Telegram.Bot.Types;
using Telegram.Bot.Types.Enums;
using Telegram.Bot.Types.ReplyMarkups;
using Telegram.Bot.Types.InputFiles;

namespace WindowsFormsApplication1
{
    public partial class Form1 : Form
    {
        private static readonly TelegramBotClient bot = new TelegramBotClient("TOKEN");
        
        private static Dictionary<long, string> dic = new Dictionary<long, string>();
                        
        public Form1()
        {
            InitializeComponent();
        }

        private void Form1_Load(object sender, EventArgs e)
        {
            try
            {
                bot.OnMessage += bot_OnMessage;
                bot.OnMessageEdited += bot_OnMessageEdited;
                bot.OnCallbackQuery += bot_OnCallbackQuery;
                bot.OnReceiveError += bot_OnReceiveError;

                var me = bot.GetMeAsync().Result;
                this.Text = me.Username;

                bot.StartReceiving();
            }
            catch (Exception ex) 
            {
                Console.WriteLine(ex.Message.ToString());
            }
        }        

        private void Form1_FormClosed(object sender, FormClosedEventArgs e)
        {
            try
            {
                dic.Clear();

                bot.StopReceiving();
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message.ToString());
            }
        }

        private async void bot_OnMessage(object sender, Telegram.Bot.Args.MessageEventArgs e)
        {
            if (e.Message.Type == Telegram.Bot.Types.Enums.MessageType.Text)
            {
                string sMessage = e.Message.Text;
                long chatid = e.Message.Chat.Id;
                string firstname = e.Message.Chat.FirstName;
                string lasttname = e.Message.Chat.LastName;


                //나에게 메시지를 보내온 chatid 저장
                if (!dic.ContainsKey(chatid))
                {
                    dic.Add(chatid, firstname);
                }
                

                if (sMessage == "/start")
                {
                    StringBuilder sMsg = new StringBuilder();
                    sMsg.AppendLine("/help -> 나의 이미지");
                    sMsg.AppendLine("/keyboard -> 나의 페이지");
                    sMsg.AppendLine("/link -> 나의 링크 페이지");
                    await bot.SendTextMessageAsync(chatid, sMsg.ToString());
                }
                else if (sMessage == "/keyboard")
                {
                    ReplyKeyboardMarkup ReplyKeyboard = new[] {
                        new[] { "https://github.com/ImDaeseong" },
                        new[] { "https://github.com/ImDaeseong/DaeseongBot" },
                    };
                    await bot.SendTextMessageAsync(chatid, "github.com/ImDaeseong 정보를 확인하세요.", replyMarkup: ReplyKeyboard);
                }
                else if (sMessage == "/link") 
                {
                    var InlineKeyboard = new InlineKeyboardMarkup(new[]
                    {
                        new[] { InlineKeyboardButton.WithUrl("Daeseong", "https://github.com/ImDaeseong") },
                        new[] { InlineKeyboardButton.WithCallbackData("ImDaeseong") }
                    });
                    await bot.SendTextMessageAsync(chatid, "나의 링크 정보를 확인하세요.", replyMarkup: InlineKeyboard);                
                }
                else if (sMessage == "/help")
                {
                    await bot.SendPhotoAsync(chatid, "https://avatars0.githubusercontent.com/u/10001221?s=460&v=4", sMessage);
                }
                else
                {
                    await bot.SendTextMessageAsync(chatid, sMessage);
                }

                Invoke(new MethodInvoker(delegate()
                {
                    textBox1.Text += "[ " + e.Message.Date.TimeOfDay + 3 + " ] 받은내용 - " + sMessage + " [" + firstname + "]" + Environment.NewLine;
                }));           
            }
            else if (e.Message.Type == Telegram.Bot.Types.Enums.MessageType.Audio)
            {
                //Console.WriteLine("Audio");

                var fileInfo = await bot.GetFileAsync(e.Message.Audio.FileId);
                using (var stream = new FileStream(e.Message.Audio.FileId + ".mp3", FileMode.Create))
                {
                    await bot.GetInfoAndDownloadFileAsync(fileInfo.FileId, stream);
                }                           
            }
            else if (e.Message.Type == Telegram.Bot.Types.Enums.MessageType.Photo)
            {
                //Console.WriteLine("Photo");

                var fileInfo = await bot.GetFileAsync(e.Message.Photo[1].FileId);
                using (var stream = new FileStream(e.Message.Photo[1].FileId + ".jpg", FileMode.Create))
                {
                    await bot.GetInfoAndDownloadFileAsync(fileInfo.FileId, stream);
                }                 
            }
            else if (e.Message.Type == Telegram.Bot.Types.Enums.MessageType.Document)
            {
                //Console.WriteLine("Document");

                var fileInfo = await bot.GetFileAsync(e.Message.Document.FileId);
                using (var stream = new FileStream(e.Message.Document.FileName, FileMode.Create))
                {
                    await bot.GetInfoAndDownloadFileAsync(fileInfo.FileId, stream);
                }
            }
            else if (e.Message.Type == Telegram.Bot.Types.Enums.MessageType.Sticker)
            {                
                //Console.WriteLine("Sticker");

                var fileInfo = await bot.GetFileAsync(e.Message.Sticker.FileId);
                using (var stream = new FileStream(e.Message.Sticker.FileId, FileMode.Create))
                {
                    await bot.GetInfoAndDownloadFileAsync(fileInfo.FileId, stream);
                }
            }
            else if (e.Message.Type == Telegram.Bot.Types.Enums.MessageType.Video)
            {
                //Console.WriteLine("Video");

                var fileInfo = await bot.GetFileAsync(e.Message.Video.FileId);
                using (var stream = new FileStream(e.Message.Video.FileId, FileMode.Create))
                {
                    await bot.GetInfoAndDownloadFileAsync(fileInfo.FileId, stream);
                }
            }
            else if (e.Message.Type == Telegram.Bot.Types.Enums.MessageType.Voice)
            {
                //Console.WriteLine("Voice");

                var fileInfo = await bot.GetFileAsync(e.Message.Voice.FileId);
                using (var stream = new FileStream(e.Message.Voice.FileId, FileMode.Create))
                {
                    await bot.GetInfoAndDownloadFileAsync(fileInfo.FileId, stream);
                }
            }
            else if (e.Message.Type == Telegram.Bot.Types.Enums.MessageType.Location)
            {
                string sMsg = string.Format("Latitude:{0} Longitude:{1}", e.Message.Location.Latitude, e.Message.Location.Longitude);
                Console.WriteLine(sMsg);
            }
            else
            {
                await bot.SendTextMessageAsync(e.Message.Chat.Id, "메시지 내용이 테스트가 아닙니다.", replyToMessageId: e.Message.MessageId);
            }
        }

        private static async void bot_OnCallbackQuery(object sender, Telegram.Bot.Args.CallbackQueryEventArgs e)
        {
            var callbackQuery = e.CallbackQuery;
            string messageBody = callbackQuery.Data;
            var chatId = callbackQuery.Message.Chat.Id;
            
            await bot.SendTextMessageAsync(chatId, messageBody);
        }

        private void bot_OnReceiveError(object sender, Telegram.Bot.Args.ReceiveErrorEventArgs e)
        {
            Console.WriteLine("bot_OnReceiveError");
        }

        private void bot_OnMessageEdited(object sender, Telegram.Bot.Args.MessageEventArgs e)
        {
            Console.WriteLine("bot_OnMessageEdited");
        }
        
        private bool IsExistFile(string sLocalPath)
        {
            FileInfo f = new FileInfo(sLocalPath);
            if (f.Exists)
                return true;
            else
                return false;
        }

        private string fileExtName(string strFilename)
        {
            int nPos = strFilename.LastIndexOf('.');
            int nLength = strFilename.Length;
            if (nPos < nLength)
                return strFilename.Substring(nPos + 1, (nLength - nPos) - 1);
            return string.Empty;
        }

        private string fileExeName(string sFileFullPath)
        {
            string sFileName = sFileFullPath.Substring(sFileFullPath.LastIndexOf("\\") + 1);
            return sFileName;
        }
                
        private static async void SendMsg(long chatId, string sMsg)
        {
            await bot.SendTextMessageAsync(chatId, sMsg);
        }

        private static async void SendPhoto(long chatId, string sPath, string sMsg)
        {
            try
            {
                using (var fileStream = new FileStream(sPath, FileMode.Open, FileAccess.Read, FileShare.Read))
                {
                    await bot.SendPhotoAsync(chatId, new InputOnlineFile(fileStream, Path.GetFileName(sPath)), sMsg);                    
                }
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message.ToString());
            }
        }

        private static async void SendAudio(long chatId, string sPath, string sMsg)
        {
            try
            {
                using (var fileStream = new FileStream(sPath, FileMode.Open, FileAccess.Read, FileShare.Read))
                {
                    await bot.SendAudioAsync(chatId, new InputOnlineFile(fileStream, Path.GetFileName(sPath)), sMsg);
                }
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message.ToString());
            }
        }

        private static async void SendSticker(long chatId, string sPath)
        {
            try
            {
                using (var fileStream = new FileStream(sPath, FileMode.Open, FileAccess.Read, FileShare.Read))
                {
                    await bot.SendStickerAsync(chatId, new InputOnlineFile(fileStream, Path.GetFileName(sPath)));
                }
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message.ToString());
            }
        }

        private static async void SendVideo(long chatId, string sPath)
        {
            try
            {
                using (var fileStream = new FileStream(sPath, FileMode.Open, FileAccess.Read, FileShare.Read))
                {
                    await bot.SendVideoAsync(chatId, new InputOnlineFile(fileStream, Path.GetFileName(sPath)));
                }
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message.ToString());
            }
        }

        private static async void SendVoice(long chatId, string sPath, string sMsg)
        {
            try
            {
                using (var fileStream = new FileStream(sPath, FileMode.Open, FileAccess.Read, FileShare.Read))
                {
                    await bot.SendVoiceAsync(chatId, new InputOnlineFile(fileStream, Path.GetFileName(sPath)), sMsg);
                }
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message.ToString());
            }
        }
        
        private async void button1_Click(object sender, EventArgs e)
        {
            try
            {
                string sMsg = textBox2.Text;
                textBox2.Text = "";

                if (sMsg == "") return;
                
                foreach (KeyValuePair<long, string> pair in dic)
                {
                    //Console.WriteLine("{0}, {1}", pair.Key, pair.Value);

                    await bot.SendTextMessageAsync(pair.Key, sMsg);

                    Invoke(new MethodInvoker(delegate()
                    {
                        textBox1.Text += " 전달내용 - " + sMsg + " [" + pair.Value + "]" + Environment.NewLine;
                    }));
                }
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message.ToString());
            }            
        }               

        private async void button2_Click(object sender, EventArgs e)
        {
            try
            {
                //이미지 체크
                string sImgPath = "";
                OpenFileDialog fileDialog = new OpenFileDialog();
                fileDialog.RestoreDirectory = true;
                if (fileDialog.ShowDialog() == DialogResult.OK)
                    sImgPath = fileDialog.FileName;

                if (sImgPath == "") return;
                if (!IsExistFile(sImgPath)) return;
                if (fileExtName(sImgPath).ToLower() != "png" && fileExtName(sImgPath).ToLower() != "bmp" && fileExtName(sImgPath).ToLower() != "gif" && fileExtName(sImgPath).ToLower() != "jpg") return;

                string sFileName = fileExeName(sImgPath);

                //테스트 체크
                string sMsg = textBox2.Text;
                textBox2.Text = "";
                if (sMsg == "")
                    sMsg =  " 이미지:" + sFileName;
                else
                    sMsg += " 이미지:" + sFileName; 

                foreach (KeyValuePair<long, string> pair in dic)
                {
                    SendPhoto(pair.Key, sImgPath, sMsg);

                    //SendAudio(pair.Key, sImgPath, sMsg);
                    //SendSticker(pair.Key, sImgPath);
                    //SendVideo(pair.Key, sImgPath);
                    //SendVoice(pair.Key, sImgPath, sMsg);
                                                                                
                    Invoke(new MethodInvoker(delegate()
                    {
                        textBox1.Text += " 전달내용 - " + sMsg + " [" + pair.Value + "]" + Environment.NewLine;
                    }));
                }
                
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message.ToString());
            }
        }

    }
}
