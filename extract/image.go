package extract

import (
	"fmt"
	"strings"

	"github.com/koffeinsource/kaffeeshare2go/data"
)

func image(i *data.Item, sourceURL string, contentType string) {
	if !(strings.Index(contentType, "image/") == 0) {
		return
	}

	fmt.Println("Running Image plugin.")

	i.ImageURL = ""
	i.Caption = sourceURL[strings.LastIndex(sourceURL, "/")+1:]
	i.Description = "<img src=\"" + sourceURL + "\">"

}