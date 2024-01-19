package main

import (
	"encoding/json"
	"fmt"
	"log"

	tb "gopkg.in/telebot.v3"
)

type checkoutHandler struct {
	bot *tb.Bot
}

func newCheckoutHandler(bot *tb.Bot) *checkoutHandler {
	return &checkoutHandler{bot: bot}
}

func (h *checkoutHandler) handleCheckoutError(err error, pre *tb.PreCheckoutQuery) {
	log.Println(err)

	// TODO: here it may be necessary to pass a masked error to the user
	// so as not to show implementation details

	_, sendErr := h.bot.Send(
		pre.Sender,
		fmt.Sprintf(
			"An error occurred while processing the payment. "+
				"If the problem is critical, inform the administrator: %s\n\n Checkout ID: %s",
			administratorUserName, pre.ID,
		),
	)
	if sendErr != nil {
		log.Println(sendErr)
	}
}

func (h *checkoutHandler) handleCheckout(ctx tb.Context) error {
	pre := ctx.PreCheckoutQuery()

	// check ID
	if pre == nil || pre.ID == "" {
		log.Println("ignore empty checkout")
		return nil
	}

	// parse checkout payload data
	inv := invoiceData{}
	err := json.Unmarshal([]byte(pre.Payload), &inv)
	if err != nil {
		h.handleCheckoutError(
			fmt.Errorf("failed to decode checkout payload from json: %w", err),
			pre,
		)
		return nil
	}

	// there may be other checks and operations on the user account

	// accept checkout
	if err := h.bot.Accept(pre); err != nil {
		h.handleCheckoutError(
			fmt.Errorf("accept checkout: %w", err),
			pre,
		)
		// you can return an error here if necessary
		return nil
	}

	// TODO: perform actions that will issue the product
	// to the user or carry out some other actions with the system

	// send notify to user
	msg := "*Payment was successful!*"
	if _, err := h.bot.Send(tb.ChatID(inv.UserTelegramID), msg, tgParseMode); err != nil {
		return fmt.Errorf("failed to send message about successful payment to user: %w", err)
	}

	log.Println("successfull payment!")
	return nil
}
