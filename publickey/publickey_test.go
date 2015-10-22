package publickey

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var keyserver_result = `<html xmlns="http://www.w3.org/1999/xhtml"><head>
<title>Search results for 'gmail elmundio1987 com'</title>
<meta http-equiv="Content-Type" content="text/html;charset=utf-8">
<style type="text/css">
/*<![CDATA[*/
 .uid { color: green; text-decoration: underline; }
 .warn { color: red; font-weight: bold; }
/*]]>*/
</style></head><body><h1>Search results for 'gmail elmundio1987 com'</h1><pre>Type bits/keyID     Date       User ID
</pre><hr><pre>pub  4096R/<a href="/pks/lookup?op=get&amp;search=0xC63AB6290F0E5CA5">0F0E5CA5</a> 2014-11-16 <a href="/pks/lookup?op=vindex&amp;search=0xC63AB6290F0E5CA5">Edmund Dipple &lt;elmundio1987@gmail.com&gt;</a>
</pre>
</body></html>`

var publicKey = `
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
dHQThJxttwBCU9N4xlwWMErinw==
=S+zB`

var keyserversTests = []struct {
	host  string
	query string
}{
	{"pgp.mit.edu", "/pks/lookup?op=index&exact=on&search="},
	{"sks-keyservers.net", "/pks/lookup?op=vindex&exact=on&search="},
}

func TestCreateQueryURLAppendsParametersCorrectly(t *testing.T) {
	assert.Equal(t, CreateQueryURL("http://keys.pgp.net", "/get.pgp?search=", "elmundio1987@gmail.com"), "http://keys.pgp.net/get.pgp?search=elmundio1987@gmail.com", "")
}

func TestGetFirstLinkFromHTML(t *testing.T) {
	assert.Equal(t, GetLinksFromHTML(keyserver_result)[0], "/pks/lookup?op=get&search=0xC63AB6290F0E5CA5", "")
}

func TestDownloadKeyfile(t *testing.T) {

	assert.NotEqual(t, DownloadFile("https://pgp.mit.edu/pks/lookup?op=index&exact=on&search=elmundio1987@gmail.com"), "", "")
}

func TestGetKeyFromEmail(t *testing.T) {

	for _, tt := range keyserversTests {
		assert.Contains(t, GetKeyFromEmail("elmundio1987@gmail.com", tt.host, tt.query), publicKey, "")
	}

}

func TestGetKeyWhenNoProtocolProvided(t *testing.T) {
	assert.Contains(t, GetKeyFromEmail("elmundio1987@gmail.com", "pgp.mit.edu", "/pks/lookup?op=index&exact=on&search="), publicKey, "")
}

func TestGetKeyFromWrongEmail(t *testing.T) {
	assert.Equal(t, GetKeyFromEmail("elmundio1988@gmail.com", "https://pgp.mit.edu", "/pks/lookup?op=index&exact=on&search="), "No keys Found", "")
}

func TestGetKeyWhenHostDown(t *testing.T) {
	assert.Equal(t, GetKeyFromEmail("elmundio1988@gmail.com", "https://pgp.mit.edu2", "/pks/lookup?op=index&exact=on&search="), "Invalid Host", "")
}
