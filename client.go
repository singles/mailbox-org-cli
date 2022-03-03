package main

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf/browser"
	"gopkg.in/headzoo/surf.v1"
)

const MANAGEMENT_URL = "https://manage.mailbox.org/index.php?p=account_disposable_aliases"

type Address struct {
	Email   string `json:"email"`
	Memo    string `json:"memo"`
	Expires string `json:"expires"`
}

type Client struct {
	browser *browser.Browser
}

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
	body := url.Values{"action": []string{"create"}}
	c.browser.PostForm(MANAGEMENT_URL, body)

	addresses := c.List()
	newAddress := addresses[len(addresses)-1]

	if memo != "" {
		c.SetMemo(newAddress.Email, memo)
	}

	for _, address := range c.List() {
		if address.Email == newAddress.Email {
			return address
		}
	}

	return Address{}
}

func (c *Client) Renew(id string) Address {
	body := url.Values{"action": []string{"renew"}, "id": []string{id}}
	c.browser.PostForm(MANAGEMENT_URL, body)

	for _, address := range c.List() {
		if address.Email == id {
			return address
		}
	}

	return Address{}
}

func (c *Client) SetMemo(id, memo string) Address {
	body := url.Values{"action": []string{"edit_memo"}, "id": []string{id}, "memo": []string{memo}}
	c.browser.PostForm(MANAGEMENT_URL, body)

	for _, address := range c.List() {
		if address.Email == id {
			return address
		}
	}

	return Address{}
}

func (c *Client) Delete(id string) {
	body := url.Values{"action": []string{"delete"}, "id": []string{id}}
	c.browser.PostForm(MANAGEMENT_URL, body)
}
