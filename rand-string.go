package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strconv"

	"github.com/codegangsta/cli"
)

var helpTemplate = `SUMMARY:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} [global options] strLength

VERSION:
   {{.Version}}

ARGUMENTS:
   strLength{{ "\t" }}{{ "\t" }}Length of String to Output

GLOBAL OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`

func main() {
	app := cli.NewApp()
	app.Name = "rand-string"
	app.Version = "1.0.0"
	app.Author = "James Newell"
	app.Email = "james.newell@gmail.com"
	app.Usage = "generates a cryptographically-secure random character string"
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "alphaNum, a", Usage: "use only alphaNumeric characters"},
		cli.StringFlag{Name: "filter, f", Value: "no-filter", Usage: "valid character filter (regex format)"},
		cli.BoolFlag{Name: "pass, p", Usage: "use only password-friendly characters"},
	}
	cli.AppHelpTemplate = helpTemplate

	app.Action = func(c *cli.Context) error {
		if len(c.Args()) < 1 {
			cli.ShowAppHelp(c)
			os.Exit(1)
		}

		alphaNumStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
		alphaNumCount := len(alphaNumStr)

		passStr := alphaNumStr + "!$%@#"
		passStrCount := len(passStr)

		strLeng, err := strconv.ParseInt(c.Args()[0], 10, 64)
		if err != nil {
			fmt.Println("Error parsing argument strLength (expects a number):", c.Args()[0])
			os.Exit(1)
		}

		var filter *regexp.Regexp
		if c.String("filter") != "no-filter" {
			_filter, err := regexp.Compile(c.String("filter"))
			if err != nil {
				fmt.Println("Error parsing option filter (expects a regex):", c.String("filter"))
				os.Exit(1)
			}
			filter = _filter
		}

		var i int64
		for i = 0; i < strLeng; i++ {
			if c.Bool("alphaNum") {
				randInt, _ := rand.Int(rand.Reader, big.NewInt(int64(alphaNumCount)))
				fmt.Printf("%c", alphaNumStr[randInt.Uint64()])
			} else if c.Bool("pass") {
				randInt, _ := rand.Int(rand.Reader, big.NewInt(int64(passStrCount)))
				fmt.Printf("%c", passStr[randInt.Uint64()])
			} else {
				randInt, _ := rand.Int(rand.Reader, big.NewInt(int64(93)))
				if filter == nil {
					fmt.Printf("%c", randInt.Uint64()+33)
				} else {
					s := fmt.Sprintf("%c", randInt.Uint64()+33)
					if filter.MatchString(s) {
						fmt.Print(s)
					} else {
						i = i - 1
					}
				}
			}
		}
		fmt.Println()
		return nil
	}

	app.Run(os.Args)
}
