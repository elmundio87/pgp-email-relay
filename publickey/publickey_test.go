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
	assert.Contains(t, GetKeyFromEmail("elmundio1987@gmail.com", "https://pgp.mit.edu", "/pks/lookup?op=index&exact=on&search="), "-----BEGIN PGP PUBLIC KEY BLOCK-----", "")
}

func TestGetKeyFromWrongEmail(t *testing.T) {
	assert.Equal(t, GetKeyFromEmail("elmundio1988@gmail.com", "https://pgp.mit.edu", "/pks/lookup?op=index&exact=on&search="), "No keys Found", "")
}

func TestGetKeyWhenHostDown(t *testing.T) {
	assert.Equal(t, GetKeyFromEmail("elmundio1988@gmail.com", "https://pgp.mit.edu2", "/pks/lookup?op=index&exact=on&search="), "Invalid Host", "")
}
