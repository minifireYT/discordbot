package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	//"time"
	"strings"
	//"strconv"
	//"github.com/foolin/pagser"
	"./bot"
	//"reflect"
)

var BotID string
var goBot *discordgo.Session
/*func Compiler(a string) []string{

}*/
func Start() {
	goBot, err := discordgo.New("Bot " + "you token here")//連接我的機器人

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}
	BotID = u.ID
	goBot.AddHandler(messageHandler)
	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	con := make([]string,3)//con是一個有3位的slice
	con = strings.Fields(m.Content)//遇到空格時分開例如:hello world -> [hello,world]
	if m.Author.ID == BotID {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID,"ping")
	}
	if len(con) == 3{
		if con[0] == "dansearch"&&con[1] != "" &&con[2] != "" &&bot.Crawling(con[1],con[2]) != nil{
			d := bot.Crawling(con[1],con[2])
			for i := 0;i < len(d);i++{
				s.ChannelMessageSend(m.ChannelID,d[i])
			}
		}
	}
}
func main() {
	Start()
	<-make(chan struct{})
	return
}
