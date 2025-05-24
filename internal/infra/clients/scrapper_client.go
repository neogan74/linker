package clients

import "net/http"

type ScapperClinet struct {
	client *http.Client
}

func NewScrapperClient() *ScapperClinet {
	return &ScapperClinet{
		client: &http.Client{},
	}
}

func (c *ScapperClinet) SendNotifiy(tgChatID int64, url string) error {
	return nil
}
