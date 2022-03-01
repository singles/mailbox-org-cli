package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf/browser"
	"gopkg.in/headzoo/surf.v1"
)

const ACTION_DELETE = "delete"
const ACTION_RENEW = "renew"

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

func (c *Client) Create() Address {
	fm, err := c.browser.Form("#content > form")
	if fm.Submit() != nil {
		panic(err)
	}

	addresses := c.List()

	return addresses[len(addresses)-1]
}

func (c *Client) Extend(id string) Address {
	fm := c.getFormForAction(id, ACTION_RENEW)
	err := fm.Submit()
	if err != nil {
		panic(err)
	}

	for _, address := range c.List() {
		if address.Email == id {
			return address
		}
	}

	return Address{}
}

func (c *Client) Delete(id string) {
	fm := c.getFormForAction(id, ACTION_DELETE)
	err := fm.Submit()
	if err != nil {
		panic(err)
	}
}

func (c *Client) getFormForAction(id, action string) *browser.Form {
	selector := c.browser.Find(".ox-list .actions input[name=action][value=" + action + "] + input[input[type=hidden][name=id][value='" + id + "']").Parent()

	return browser.NewForm(c.browser, selector)
}
