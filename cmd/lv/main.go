package main

import (
	"bufio"
	"flag"
	"os"
	"strings"

	"github.com/alecthomas/chroma/quick"

	lv "github.com/gaqzi/log-viewer"
)

var (
	splitAt = flag.String(
		"sep",
		"",
		"if the string does not start with '{' "+
			"then split by this character and indent the rest of the line",
	)
	style = flag.String(
		"style",
		"base16-snazzy",
		"colorizes JSON output with this style. ",
	)
)

func main() {
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		str, err := lv.JSON(line)

		if err != nil {
			if *splitAt == "" {
				print(line)
				continue
			}

			splits := strings.SplitN(line, *splitAt, 2)
			if len(splits) != 2 {
				print(line)
				continue
			}

			str, err = lv.JSON(splits[1], true)
			if err != nil {
				print(line)
				continue
			}

			print(splits[0])
		}

		if err := quick.Highlight(
			os.Stdout,
			string(str),
			"json",
			"terminal256",
			*style,
		); err != nil {
			print(str)
		}
	}
}
