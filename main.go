package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func usage(exitValue int) {
	println("USAGE")
	println("\tgo run main.go [options]")
	println("DESCRIPTION")
	println("OPTIONS")
	println("\t-C, --cli")
	println("\t\tstart youtubeDownloader in CLI mode")
	println("\t-S, --server")
	println("\t\tstart youtubeDownloader through a browser")
	os.Exit(exitValue)
}

func isInt(s string) bool {
    l := len(s)
    if strings.HasPrefix(s, "-") {
        l = l - 1
        s = s[1:]
    }

    reg := fmt.Sprintf("\\d{%d}", l)

    rs, err := regexp.MatchString(reg, s)

    if err != nil {
        return false
    }
    return rs
}

func main() {
	// flag.Parse()
	// fmt.Println(flag.Args())
	if len(os.Args) != 2 {
		usage(exitFailure)
	}
	if os.Args[1] == "-C" || os.Args[1] == "--cli" {
		cli := newCLI()
		cli.start()
	} else if os.Args[1] == "-S" || os.Args[1] == "--server" {
		api := newServer()
		api.start()
	} else {
		usage(exitFailure)
	}
}
