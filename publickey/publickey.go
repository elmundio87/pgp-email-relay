package publickey

import "golang.org/x/net/html"
import "bytes"
import "net/http"
import "io/ioutil"

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

func DownloadQueryResult(query string) string {
	resp, _ := http.Get(query)
	bytes, _ := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	return string(bytes)
}

func DownloadKeyfile(url string) string {
	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	return string(bytes)
}

func GetKeyFromEmail(email string, host string, query string) string {
	keyserverLink := CreateQueryURL(host, query, email)
	html := DownloadQueryResult(keyserverLink)
	links := GetLinksFromHTML(html)
	return DownloadKeyfile(host + string(links[0]))
}
