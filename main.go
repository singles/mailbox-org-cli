package main

import (
	"bufio"
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
	Memo string
}

type args struct {
	Username string          `arg:"required" help:"mailbox.org username"`
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

	switch {
	case args.List != nil:
		fmt.Println("listing addressess")
	case args.Renew != nil:
		fmt.Println("Renewing address")
	case args.Delete != nil:
		fmt.Println("Creating address")
	case args.Create != nil:
		fmt.Println("Adding disposable address")
	case args.SetMemo != nil:
		fmt.Println("Setting memo")
	default:
		p.Fail("Invalid command")
	}
}

func readPasswordFromStdin(stdin io.Reader) string {
	reader := bufio.NewReader(os.Stdin)
	password, _, _ := reader.ReadLine()

	return string(password)
}
