package main

import (
	"encoding/json"
	"errors"

	tb "gopkg.in/telebot.v3"
)

type invoiceData struct {
	SystemPaymentID int64 `json:"paymentID"`
	UserTelegramID  int64 `json:"tid"`
}

func createInvoice(
	paymentID int64,
	userTelegramID int64,
) (*tb.Invoice, error) {
	// create invoice payload
	payload := invoiceData{
		SystemPaymentID: paymentID,
		UserTelegramID:  userTelegramID,
	}

	// encode invoice payload
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.New("failed to encode invoice data: " + err.Error())
	}

	// send invoice
	totalAmount := paymentAmount * currencyMultiplier
	invoice := tb.Invoice{
		Title:       paymentTitle,
		Description: paymentDescription,
		Currency:    paymentCurrency,
		Prices: []tb.Price{
			{
				Label:  paymentLabel,
				Amount: totalAmount,
			},
		},
		Token:   telegramPaymentToken,
		Total:   totalAmount,
		Payload: string(payloadBytes),
	}
	return &invoice, nil
}
