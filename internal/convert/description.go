package convert

import (
	"bytes"
	"fmt"
	"strconv"
)

type Description struct {
	description *string
}

func NewDescription(d *string) Description {
	return Description{
		description: d,
	}
}

func (d Description) Description() string {
	if d.description == nil {
		return ""
	}

	return *d.description
}

func (d Description) Schema() []byte {
	var b bytes.Buffer

	if d.description != nil {
		quotedDescription := strconv.Quote(*d.description)

		b.WriteString(fmt.Sprintf("Description: %s,\n", quotedDescription))
		b.WriteString(fmt.Sprintf("MarkdownDescription: %s,\n", quotedDescription))
	}

	return b.Bytes()
}
