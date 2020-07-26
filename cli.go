package main

import (
	"os"
	"bufio"
	"strconv"
)

//CLI struct
type CLI struct
{
	links []string
	scanner *bufio.Scanner
	ytd	ytDownloader
}

func newCLI() CLI {
	return CLI{scanner: bufio.NewScanner(os.Stdin), ytd: newYtDownloader()}
}

func (cli *CLI) input(task string) string {
	print(task)
	data := getLine(cli.scanner)
	return data
}

//todo code a better cli if you are not lazy
//Note You can use that https://github.com/urfave/cli/blob/master/docs/v2/manual.md
func (cli *CLI) start() {
	println("Welcome to youtubeDownloader CLI")
	for {
		println("\n")
		ytLink := cli.input("Paste the video link here: ")
		err := cli.ytd.loadVideoInfo(ytLink)
		if err != nil {
			continue
		}
		//create directory
		for {
			cli.ytd.printVideoInfo()
			input := cli.input("pick a format: ")
			if !isInt(input) {
				println("wrong input")
				continue
			}
			var filename string
			formatIndex, _ := strconv.Atoi(input)

			filename, err = cli.ytd.download(downloadDirectory, uint(formatIndex))
			if err != nil {
				println("wrong input")
				continue
			}
			println("The Video has been saved successfully inside: ", downloadDirectory + filename)
			break
		}
	}
}
