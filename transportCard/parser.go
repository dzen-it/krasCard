package transportCard

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type TransportCard struct {
	CurrentBalance string
	ErrorText      string
}

func ParseInfo(page *http.Response) (*TransportCard, error) {
	doc, err := goquery.NewDocumentFromResponse(page)
	if err != nil {
		return nil, err
	}

	sel := doc.Find("div.main-col > table:nth-child(5) > tbody > tr:nth-child(2) > td:nth-child(2)")

	tc := new(TransportCard)
	tc.CurrentBalance = sel.Contents().Text()
	tc.ErrorText = parseError(doc)

	return tc, err
}

func parseError(doc *goquery.Document) string {
	sel := doc.Find("div.main-col > div.alert.alert-error > p")
	text := sel.Contents().Text()
	return text
}
