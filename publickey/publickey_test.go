package publickey

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateQueryURLAppendsParametersCorrectly(t *testing.T) {
	assert.Equal(t, CreateQueryURL("http://keys.pgp.net/get.pgp?search=", "elmundio1987@gmail.com"), "http://keys.pgp.net/get.pgp?search=elmundio1987@gmail.com", "")
}

func TestGetResultPageHTML(t *testing.T) {

}
