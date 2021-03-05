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
		"if the line doesn't format as JSON then try to split by this sep and format",
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
		line := scanner.Text() + "\n"
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

			print(splits[0] + "\n")
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
