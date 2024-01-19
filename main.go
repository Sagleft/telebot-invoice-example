package main

import (
	"fmt"
	"log"
	"time"

	tb "gopkg.in/telebot.v3"
)

const (
	// this can be put into env parameters.
	// you can use library: https://github.com/kelseyhightower/envconfig
	administratorUserName = "@adminUserName"

	// get it from https://t.me/BotFather
	telegramPaymentToken = "000000000:TEST:00000"
	telegramBotToken     = "0000000000:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"

	// for USD: 100
	currencyMultiplier = 100

	paymentAmount      = 500
	paymentTitle       = "Invoice title"
	paymentDescription = "Invoice description"
	paymentCurrency    = "RUB"
	paymentLabel       = "Cost of the service"

	tgParseMode = tb.ModeMarkdown
)

func main() {
	bot, err := createBot(telegramBotToken)
	if err != nil {
		log.Fatalln(err)
	}

	c := newCheckoutHandler(bot)

	bot.Handle(tb.OnCheckout, c.handleCheckout)
	bot.Handle("/start", func(ctx tb.Context) error {
		// here we will create and send an invoice
		// at the first message from the user, for example
		paymentID := int64(1) // here you need to assign a payment ID
		telegramUserID := ctx.Sender().ID
		invoice, err := createInvoice(paymentID, telegramUserID)
		if err != nil {
			return fmt.Errorf("create invoice: %w", err)
		}

		// send invoice
		_, err = invoice.Send(bot, ctx.Sender(), nil)
		if err != nil {
			return fmt.Errorf("send invoice: %w", err)
		}
		return nil
	})

	bot.Start()
}

func createBot(telegramToken string) (*tb.Bot, error) {
	return tb.NewBot(tb.Settings{
		Token:  telegramToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
}
