package main

import "strings"

const StartTagToken = "{"
const EndTagToken = "}"
const TagSeparator = ","

func Tags(line string) ([]string, error) {
	beginTokenIndex := strings.Index(line, StartTagToken) 
	endTokenIndex := strings.Index(line, EndTagToken)
	if endTokenIndex == -1 || beginTokenIndex == -1 {
		return []string{}, nil
	}

	tagContents := line[beginTokenIndex+1:endTokenIndex]
	tagContents = strings.ReplaceAll(tagContents, " ", "")
	tags := strings.Split(tagContents, TagSeparator)
	return tags, nil
}
