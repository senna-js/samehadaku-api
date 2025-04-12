// this file contains function to get the stream iframe
package utility

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

// form body struct for getting the iframe
type IFrameBody struct {
	// Action is optional
	Action string
	// Post number
	Post string
	// Nume is optional
	Nume string
	// Type is optional
	ResponseType string
}

// extract the struct into form body string
//
// if `i.Action`, i.Nume, and i.Post is not set, it will fill with the default value
func (i *IFrameBody) String() string {
	payload := url.Values{}

	if i.Action == "" {
		i.Action = "player_ajax"
	}

	if i.Nume == "" {
		i.Nume = "1"
	}

	if i.ResponseType == "" {
		i.ResponseType = "schtml"
	}

	payload.Add("action", i.Action)
	payload.Add("post", i.Post)
	payload.Add("nume", i.Nume)
	payload.Add("type", i.ResponseType)

	return payload.Encode()
}

// get iframe url based on the opsi. field `referer` is optional.
func GetIFrameURL(referer string, i IFrameBody) (response string, err error) {
	req, err := http.NewRequest("POST", ADMIN_URL, strings.NewReader(i.String()))
	if err != nil {
		return
	}

	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.7")
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("sec-ch-ua", "\"Chromium\";v=\"134\", \"Not:A-Brand\";v=\"24\", \"Brave\";v=\"134\"")
	req.Header.Set("sec-ch-ua-mobile", "?1")
	req.Header.Set("sec-ch-ua-platform", "\"Android\"")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("referer", referer)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return string(body), nil
}
