package clients

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/neogan74/linker/internal/api/bot_api/codegen"
	"github.com/neogan74/linker/internal/errors"
)

type ScapperClinet struct {
	client *http.Client
}

func NewScrapperClient() *ScapperClinet {
	return &ScapperClinet{
		client: &http.Client{},
	}
}

func (c *ScapperClinet) SendNotifiy(tgChatID int64, url string) error {
	req := codegen.LinkUpdate{
		Description: func() *string {
			s := "Link Update"
			return &s
		}(),
		Id:        nil,
		TgChatIds: &[]int64{tgChatID},
		Url:       &url,
	}
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", "http://localhost:8080/updates", bytes.NewBuffer(jsonBody))

	httpData, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	if httpData.StatusCode != http.StatusOK {
		return errors.BadRequestError{}
	}

	return nil
}
