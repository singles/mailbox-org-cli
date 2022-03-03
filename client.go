package main

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf/browser"
	"gopkg.in/headzoo/surf.v1"
)

type Address struct {
	Email   string `json:"email"`
	Memo    string `json:"memo"`
	Expires string `json:"expires"`
}

type Client struct {
	browser *browser.Browser
}

type FormPayload map[string]string

func NewClient(username, password string) *Client {
	client := &Client{browser: surf.NewBrowser()}

	err := client.browser.Open("https://manage.mailbox.org/login.php?redirect=account_disposable_aliases")
	if err != nil {
		panic(err)
	}

	fm, _ := client.browser.Form("#io-ox-login-form")
	fm.Input("username", username)
	fm.Input("password", password)
	if fm.Submit() != nil {
		panic(err)
	}

	return client
}

func (c *Client) List() []Address {
	addresses := []Address{}

	c.browser.Find(".ox-list li").Each(func(_ int, s *goquery.Selection) {
		email := s.Find(".title div").Text()
		memo := s.Find(".memo #memo").AttrOr("value", "<no memo>")
		expires := s.Find(".content div").Text()

		addresses = append(addresses, Address{
			Email:   email,
			Memo:    memo,
			Expires: expires,
		})
	})

	return addresses
}

func (c *Client) Create(memo string) Address {
	c.executeAction(FormPayload{"action": "create"})

	addresses := c.List()
	newAddress := addresses[len(addresses)-1]

	if memo != "" {
		c.SetMemo(newAddress.Email, memo)
	}

	return c.findAddressByID(newAddress.Email)
}

func (c *Client) Renew(id string) Address {
	c.executeAction(FormPayload{"action": "renew", "id": id})

	return c.findAddressByID(id)
}

func (c *Client) SetMemo(id, memo string) Address {
	c.executeAction(FormPayload{"action": "edit_memo", "id": id, "memo": memo})

	return c.findAddressByID(id)
}

func (c *Client) Delete(id string) {
	c.executeAction(FormPayload{"action": "delete", "id": id})
}

func (c *Client) executeAction(data FormPayload) {
	values := url.Values{}

	for key, val := range data {
		values.Set(key, val)
	}

	c.browser.PostForm("https://manage.mailbox.org/index.php?p=account_disposable_aliases", values)
}

func (c *Client) findAddressByID(id string) Address {
	for _, address := range c.List() {
		if address.Email == id {
			return address
		}
	}

	return Address{}
}
