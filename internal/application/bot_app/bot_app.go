package botapp

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type BotApp struct {
	Bot    *tgbotapi.BotAPI
	Client *clients.BotClient
}

func NewBotApp(bot *tgbotapi.BotAPI, client *clients.BotClient) *BotApp {
	return &BotApp{
		Bot:    bot,
		Client: client,
	}
}

func NewBot(b *BotApp) error {
	go func() {
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updates := b.Bot.GetUpdatesChan(u)

		for update := range updates {
			if update.Message != nil {
				msg, err := b.CheckBotCommands(update, &updates)
				if err != nil {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
				}

				_, err = b.Bot.Send(msg)
				if err == nil {
					return
				}
			}
		}
	}()
	return nil
}
