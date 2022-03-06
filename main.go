package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/alexflint/go-arg"
)

type ListCommand struct{}

type RenewCommand struct {
	ID string `arg:"required"`
}

type DeleteCommand struct {
	ID string `arg:"required"`
}

type SetMemoCommand struct {
	ID   string `arg:"required"`
	Memo string `arg:"required"`
}

type CreateCommand struct {
	Memo string `default:""`
}

type args struct {
	Username string          `arg:"required,env:MAILBOX_ORG_USERNAME" help:"mailbox.org username"`
	List     *ListCommand    `arg:"subcommand:list" help:"list dispossable addresses"`
	Renew    *RenewCommand   `arg:"subcommand:renew" help:"renew dispossable address"`
	Delete   *DeleteCommand  `arg:"subcommand:delete" help:"delete dispossable address"`
	SetMemo  *SetMemoCommand `arg:"subcommand:set-memo" help:"set-memo on existing dispossable address"`
	Create   *CreateCommand  `arg:"subcommand:create" help:"create new dispossable address with optional memo"`
}

func (args) Description() string {
	return "Commad line \"client\" for mailbox.org dispossable addresses feature"
}

func main() {
	var args args
	p := arg.MustParse(&args)

	stdin := os.Stdin
	stat, _ := stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("You must pipe password into this command, exiting.")
		os.Exit(1)
	}

	client, err := NewClient(args.Username, readPasswordFromStdin(stdin))

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	var data interface{}
	switch {
	case args.List != nil:
		data = client.List()
	case args.Renew != nil:
		data = client.Renew(args.Renew.ID)
	case args.Delete != nil:
		client.Delete(args.Delete.ID)
	case args.Create != nil:
		data = client.Create(args.Create.Memo)
	case args.SetMemo != nil:
		data = client.SetMemo(args.SetMemo.ID, args.SetMemo.Memo)
	default:
		p.Fail("Invalid command")
	}

	if data != nil {
		output, err := json.MarshalIndent(data, "", "  ")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(output))
	}

}

func readPasswordFromStdin(stdin io.Reader) string {
	reader := bufio.NewReader(os.Stdin)
	password, _, _ := reader.ReadLine()

	return string(password)
}
