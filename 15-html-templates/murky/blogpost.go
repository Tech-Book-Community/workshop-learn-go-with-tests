package murkyblogrender

import (
	"strings"
)

type Post struct {
	Title       string
	Description string
	Body        string
	Tags        []string
}

func (p *Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}
