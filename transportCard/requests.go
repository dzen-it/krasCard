package transportCard

import (
	"krasCard/common"
	"net/http"
	"net/url"
	"strings"
)

func GetInfoPage(csrf, cardID string) (*http.Response, error) {
	client := &http.Client{}

	form := url.Values{}
	form.Add("card_type", "transport")
	form.Add("csrf_token", csrf)
	form.Add("card_num", cardID)

	request, err := http.NewRequest(http.MethodPost, common.ENDPOINT_POST_TRANSPORT, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.111 Safari/537.36")

	request.AddCookie(&http.Cookie{
		Name:  "csrf_token",
		Value: csrf,
	})

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, err
}
