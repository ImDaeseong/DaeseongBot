﻿using System;
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

        private static async void SendMsg(long chatId, string sMsg)
        {
            await bot.SendTextMessageAsync(chatId, sMsg);
        }

        private static async void SendPhoto(string sImgPath, long chatId, string sMsg)
        {
            using (var fileStream = new FileStream(sImgPath, FileMode.Open, FileAccess.Read, FileShare.Read))
            {
                await bot.SendPhotoAsync(chatId, fileStream, sMsg);
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

                //테스트 체크
                string sMsg = textBox2.Text;
                textBox2.Text = "";
                if (sMsg == "")
                    sMsg = "이미지 전달";

                foreach (KeyValuePair<long, string> pair in dic)
                {
                    using (var fileStream = new FileStream(sImgPath, FileMode.Open, FileAccess.Read, FileShare.Read))
                    {
                        await bot.SendPhotoAsync(pair.Key, fileStream, sMsg);
                    }

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
