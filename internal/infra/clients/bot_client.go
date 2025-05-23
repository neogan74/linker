package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/neogan74/linker/internal/api/scrapper_api/codegen"
	"github.com/neogan74/linker/internal/errors"
)

type BotClient struct {
	client *http.Client
}

func NewBotClient() *BotClient {
	return &BotClient{
		client: &http.Client{},
	}
}

func (c *BotClient) GetLinks(tgChatid int64) ([]string, error) {
	req, err := http.NewRequest("GET", "http://localhost:8081/links", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Tg-Chat-ID", strconv.Itoa(int(tgChatid)))
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.BadRequestError{}
	}

	var links codegen.ListLinksResponse
	if err = json.NewDecoder(resp.Body).Decode(&links); err != nil {
		return nil, err
	}

	var r []string
	if links.Links != nil {
		for _, l := range *links.Links {
			r = append(r, *l.Url)
		}
	}
	return r, nil
}

func (c *BotClient) SaveTgChat(tgChatID int64) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:8081/tg-chat/%d", tgChatID), nil)
	if err != nil {
		return err
	}

	httpResp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return errors.BadRequestError{}
	}

	return nil
}

func (c *BotClient) SaveLink(tgChatID int64, link string, filters []string, tags []string) error {
	r := codegen.AddLinkRequest{Link: &link, Filters: &filters, Tags: &tags}
	jsonBody, err := json.Marshal(r)
	req, err := http.NewRequest("POST", "http://localhost:8081/links", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Add("Tg-Chat-ID", strconv.Itoa(int(tgChatID)))
	httpResp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return errors.BadRequestError{}
	}
	return nil
}

func (c *BotClient) DeleteLink(tgChatID int64, link string) error {
	r := codegen.RemoveLinkRequest{Link: &link}
	jsonBody, err := json.Marshal(r)
	req, err := http.NewRequest("DELETE", "http://localhost:8081/links", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Add("Tg-Chat-ID", strconv.Itoa(int(tgChatID)))
	httpResp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		var res codegen.ApiErrorResponse
		if err := json.NewDecoder(httpResp.Body).Decode(&res); err != nil {
			return err
		}
		if *res.Code == "404" {
			return errors.UrlNotFoundError{}
		} else {
			return errors.BadRequestError{}
		}
	}
	return nil
}
