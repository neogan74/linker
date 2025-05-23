package botapp

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *BotApp) CheckBotCommands(upd tgbotapi.Update, updChan *tgbotapi.UpdatesChannel) (tgbotapi.MessageConfig, error) {
	switch upd.Message.Command() {
	case "start":
		return b.StartCommand(upd)
	case "help":
		return b.HelpCommand(upd)
	case "track":
		return b.TrackCommand(upd)
	case "untrack":
		return b.UntrackCommand(upd)
	case "last":
		return b.LastCommand(upd)
	default:
		return tgbotapi.MessageConfig{}, errors.UnknownCommand{}
	}
}

func (b *BotApp) StartCommand(upd tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	panic("unimplemented")
}

func (b *BotApp) HelpCommand(upd tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	panic("unimplemented")
}

func (b *BotApp) TrackCommand(upd tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	panic("unimplemented")
}

func (b *BotApp) UntrackCommand(upd tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	panic("unimplemented")
}

func (b *BotApp) LastCommand(upd tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	panic("unimplemented")
}
