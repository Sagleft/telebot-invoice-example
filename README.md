# telebot-invoice-example
An example of making a payment using library tucnak/telebot: create, send and handle invoice

## The meaning of the project

There is very little information and examples on making payments with libraries on Golang to Telegram, so I decided to make a comprehensive example showing how payment works, how an invoice is created and received.

This code consists of three parts:

1. Creating a bot that processes the `/start` command.
2. Creating an invoice and sending it to the user.
3. Processing payment and completing an action, as well as sending a notification to the user about successful payment.
