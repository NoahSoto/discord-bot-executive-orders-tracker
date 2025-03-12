package main

import bot "eggexcutive-orders/informer/bot"

func main() {
	bot.BotToken = ""
	bot.Run() // call the run function of bot/bot.go
}
