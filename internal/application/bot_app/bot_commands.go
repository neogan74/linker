package bot_app

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/neogan74/linker/internal/errors"
)

func (b *BotApp) CheckBotCommands(upd tgbotapi.Update, updChan *tgbotapi.UpdatesChannel) (tgbotapi.MessageConfig, error) {
	switch upd.Message.Command() {
	case "start":
		return b.StartCommand(upd)
	case "help":
		return b.HelpCommand(upd)
	case "track":
		return b.TrackCommand(upd, updChan)
	case "untrack":
		return b.UntrackCommand(upd)
	case "list":
		return b.ListCommand(upd)
	default:
		return tgbotapi.MessageConfig{}, errors.UnknownCommand{}
	}
}

func (b *BotApp) StartCommand(upd tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	err := b.Client.SaveTgChat(upd.Message.Chat.ID)
	if err != nil {
		return tgbotapi.MessageConfig{}, err
	}
	message := tgbotapi.NewMessage(upd.Message.Chat.ID, "Welcome, you are registered!")
	return message, nil
}

func (b *BotApp) HelpCommand(upd tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	message := tgbotapi.NewMessage(upd.Message.Chat.ID, "I can help you with the follwoing commands: /start, /help, /track, /untrack, /list")
	return message, nil
}

func (b *BotApp) TrackCommand(upd tgbotapi.Update, updChan *tgbotapi.UpdatesChannel) (tgbotapi.MessageConfig, error) {
	url := upd.Message.CommandArguments()
	_, err := b.Bot.Send(tgbotapi.NewMessage(upd.Message.Chat.ID, "Enter tags(space separated)"))
	if err != nil {
		return tgbotapi.MessageConfig{}, err
	}
	tags := b.CheckUpdates(updChan)
	_, err = b.Bot.Send(tgbotapi.NewMessage(upd.Message.Chat.ID, "Enter filters(space separated)"))
	if err != nil {
		return tgbotapi.MessageConfig{}, err
	}
	filters := b.CheckUpdates(updChan)
	err = b.Client.SaveLink(upd.Message.Chat.ID, url, strings.Split(filters.(string), " "), strings.Split(tags.(string), " "))
	if err != nil {
		return tgbotapi.MessageConfig{}, err
	}
	message := tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf("Now you are following %s", url))
	return message, nil
}

func (b *BotApp) UntrackCommand(upd tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	arg := upd.Message.CommandArguments()
	err := b.Client.DeleteLink(upd.Message.Chat.ID, arg)
	if err != nil {
		return tgbotapi.MessageConfig{}, err
	}

	message := tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf("Now you are not following %s", arg))
	return message, nil
}

func (b *BotApp) ListCommand(upd tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	links, err := b.Client.GetLinks(upd.Message.Chat.ID)
	if err != nil {
		return tgbotapi.MessageConfig{}, err
	}
	if len(links) == 0 {
		return tgbotapi.NewMessage(upd.Message.Chat.ID, "List of your tracks is empty"), nil
	}

	t := "There is list of your tracks:\n"
	for _, link := range links {
		t += link + "\n"
	}
	message := tgbotapi.NewMessage(upd.Message.Chat.ID, t)
	return message, nil
}

func (b *BotApp) CheckUpdates(updChan *tgbotapi.UpdatesChannel) any {
	for upd := range *updChan {
		if upd.Message != nil {
			return upd.Message.Text
		}
	}
	return ""
}
