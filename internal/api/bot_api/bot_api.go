package bot_api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/neogan74/linker/internal/api/bot_api/codegen"
	"github.com/neogan74/linker/internal/application/bot_app"
)

type BotHandler struct {
	a *bot_app.BotApp
}

func NewBotHandler(a *bot_app.BotApp) codegen.ServerInterface {
	return &BotHandler{a}
}

func GetApiErrorResponse(code *string, desc *string) codegen.ApiErrorResponse {
	return codegen.ApiErrorResponse{
		Code:        code,
		Description: desc,
	}
}

func (h *BotHandler) PostUpdates(w http.ResponseWriter, r *http.Request) {
	log.Println("Post")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		code := "400"
		desc := err.Error()
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(GetApiErrorResponse(&code, &desc))
		return
	}

	var req codegen.LinkUpdate
	err = json.Unmarshal(d, &req)
	if err != nil {
		code := "400"
		desc := err.Error()
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(GetApiErrorResponse(&code, &desc))
		return
	}

	_, err = h.a.Bot.Send(tgbotapi.NewMessage((*req.TgChatIds)[0], fmt.Sprintf("Notification %s", *req.Url)))
	if err != nil {
		code := "400"
		desc := err.Error()
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(GetApiErrorResponse(&code, &desc))
		return
	}

	w.WriteHeader(http.StatusOK)
}
