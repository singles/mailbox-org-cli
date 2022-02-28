package main

import (
	"encoding/json"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
)

type Address struct {
	Email   string `json:"email"`
	Memo    string `json:"memo"`
	Expires string `json:"expires"`
}

func main() {
	bow := surf.NewBrowser()
	err := bow.Open("https://manage.mailbox.org/login.php?redirect=account_disposable_aliases")
	if err != nil {
		panic(err)
	}

	fm, _ := bow.Form("#io-ox-login-form")
	fm.Input("username", "xxx")
	fm.Input("password", "xxx")
	if fm.Submit() != nil {
		panic(err)
	}

	addresses := []Address{}

	bow.Find(".ox-list li").Each(func(_ int, s *goquery.Selection) {
		email := s.Find(".title div").Text()
		memo := s.Find(".memo #memo").AttrOr("value", "<no memo>")
		expires := s.Find(".content div").Text()

		addresses = append(addresses, Address{
			Email:   email,
			Memo:    memo,
			Expires: expires,
		})
	})

	addressesJson, err := json.Marshal(addresses)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(addressesJson))
}
