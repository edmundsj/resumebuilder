package main
import "github.com/russross/blackfriday/v2"
import "os"

func ParseMarkdown(filename string) ([]byte, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}

	return blackfriday.Run(contents), nil
}
