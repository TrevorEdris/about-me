package embedded

import (
	"fmt"
)

type (
	Link struct {
		URL  string
		Text string
		HTML string
	}
)

func NewLink(url, text string) Link {
	return Link{
		URL:  url,
		Text: text,
		HTML: fmt.Sprintf("<a href=\"%s\">%s</a>", url, text),
	}
}
