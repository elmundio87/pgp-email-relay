package publickey

import "golang.org/x/net/html"
import "bytes"
import "net/http"
import "io/ioutil"
import "strings"

type HtmlOutput struct {
	body string
	code int
	err  error
}

func CreateQueryURL(host string, query string, email string) string {
	return host + query + email
}

func GetLinksFromHTML(body string) []string {
	x := bytes.NewBufferString(body)

	z := html.NewTokenizer(x)

	links := []string{}

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return links
		case tt == html.StartTagToken:
			t := z.Token()

			isAnchor := t.Data == "a"
			if isAnchor {
				for _, a := range t.Attr {
					if a.Key == "href" {
						links = append(links, a.Val)
						break
					}
				}
			}
		}
	}
}

func FormatUrl(url string) string {
	if strings.HasPrefix(url, "http") {
		return url
	} else {
		return "http://" + url
	}
}

func DownloadFile(url string) HtmlOutput {
	resp, err := http.Get(FormatUrl(url))

	if err != nil {
		return HtmlOutput{"", 404, err}
	}

	bytes, _ := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	return HtmlOutput{string(bytes), resp.StatusCode, nil}
}

func GetKeyFromEmail(email string, host string, query string) string {
	keyserverLink := CreateQueryURL(host, query, email)
	html := DownloadFile(keyserverLink)

	if html.err != nil {
		return "invalid host"
	}

	links := GetLinksFromHTML(html.body)
	if len(links) == 0 {
		return "no keys found"
	}

	keyLink := host + string(links[0])

	return DownloadFile(keyLink).body
}
