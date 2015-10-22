package publickey

import "golang.org/x/net/html"
import "bytes"

func CreateQueryURL(baseurl string, email string) string {
	return baseurl + email
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
