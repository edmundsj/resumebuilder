package main
import (
	"os"
	"log"
	e "errors"
)

func parseArgs() (string, []string, error) {
	if len(os.Args) < 2 {
		return "", []string{}, e.New("You must supply a filename of your resume")
	}
	filename := os.Args[1]
	tags := os.Args[2:]
	return filename, tags, nil
}

func main() {
	filename, tags, err := parseArgs()
	if err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}	
	log.Println("Running with filename", filename)
	log.Println("Using tags:", tags)
		
	contents, _ := ParseMarkdown(filename)
	log.Println(contents)
	
}
