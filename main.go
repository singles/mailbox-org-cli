package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
)

type Address struct {
	Email   string `json:"email"`
	Memo    string `json:"memo"`
	Expires string `json:"expires"`
}

func main() {
	username := flag.String("username", "", "mailbox.org username")
	flag.Parse()

	if *username == "" {
		fmt.Println("Usage: command-that-passess | mailbox-org-cli --username user@example.com")
		flag.PrintDefaults()
		os.Exit(1)
	}

	password := readPasswordFromStdin(os.Stdin)

	bow := surf.NewBrowser()
	err := bow.Open("https://manage.mailbox.org/login.php?redirect=account_disposable_aliases")
	if err != nil {
		panic(err)
	}

	fm, _ := bow.Form("#io-ox-login-form")
	fm.Input("username", *username)
	fm.Input("password", password)
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

func readPasswordFromStdin(stdin io.Reader) string {
	reader := bufio.NewReader(os.Stdin)
	password, _, _ := reader.ReadLine()

	return string(password)
}
