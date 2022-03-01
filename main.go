package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	username := flag.String("username", "", "mailbox.org username")
	flag.Parse()

	if *username == "" {
		fmt.Println("Usage: command-that-passess | mailbox-org-cli --username user@example.com")
		flag.PrintDefaults()
		os.Exit(1)
	}

	password := readPasswordFromStdin(os.Stdin)

	client := NewClient(*username, password)

	addressesJson, err := json.MarshalIndent(client.List(), "", "  ")

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
