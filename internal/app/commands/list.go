package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outMsgText := "Список продуктов: \n\n"
	products := c.productService.List()
	for _, p := range products {
		outMsgText += p.Title
		outMsgText += "\n"

	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outMsgText)
	c.bot.Send(msg)
}
