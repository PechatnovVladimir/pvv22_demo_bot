package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/PechatnovVladimir/pvv22_demo_bot/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	productService *product.Service,
) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {

	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		args := strings.Split(update.CallbackQuery.Data, "_")
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Command: %s\n", args[0])+
				fmt.Sprintf("Offset: %s\n", args[1]),
		)
		c.bot.Send(msg)
		return
	}

	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	default:
		c.Default(update.Message)
	}

}
