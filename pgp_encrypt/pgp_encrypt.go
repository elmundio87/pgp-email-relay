package pgp_encrypt

import(
"bytes"
"os"
"path"
"github.com/elmundio87/pgp-email-relay/publickey"
"fmt"
"io"
"io/ioutil"
"golang.org/x/crypto/openpgp"
"golang.org/x/crypto/openpgp/armor"
"github.com/cryptix/go/logging"
"net/smtp"
"net/mail"
)

const encryptionType = "PGP MESSAGE"

func HandleMail(client_data string, client_rcpt_to string, gConfig map[string]string){
    
    var to = client_rcpt_to[1 : len(client_rcpt_to)-1]

    emailData := client_data[:len(client_data)-4]
    msg, _ := mail.ReadMessage(bytes.NewBuffer([]byte(emailData)))

    headers := msg.Header

    headersString := ""

    for key, value := range headers {
      fmt.Println("Key:", key, "Value:", value)
      headersString = headersString + key + ": " + value[0] + "\n"
    }

    headersString = headersString + "\n"

    body, _ := ioutil.ReadAll(msg.Body)

    encryptedBody := encrypt(string(body), to, gConfig)

    sendEmail(headersString+encryptedBody, to, gConfig)
}

func encrypt(input string, email string, gConfig map[string]string) string {

  os.MkdirAll(gConfig["PGP_KEY_FOLDER"], 0777)
  keyfileName := path.Join(gConfig["PGP_KEY_FOLDER"], email+".asc")
  keyfileExists, _ := exists(keyfileName)
  if !keyfileExists {

    key := publickey.GetKeyFromEmail(email, gConfig["PGP_KEYSERVER"], gConfig["PGP_KEYSERVER_QUERY"])
    if key == "no keys found" {
      return key + " on keyserver " + gConfig["PGP_KEYSERVER"] + " from query " + gConfig["PGP_KEYSERVER"] + gConfig["PGP_KEYSERVER_QUERY"] + email
    }

    if key == "invalid host" {
      return gConfig["PGP_KEYSERVER"] + " is offline and your key has not previously been cached."
    }

    f, err := os.Create(keyfileName)
    if err != nil {
      fmt.Println(err)
    }
    n, err := io.WriteString(f, key)
    if err != nil {
      fmt.Println(n, err)
    }
    f.Close()
  }

  to, err := os.Open(keyfileName)
  logging.CheckFatal(err)

  defer to.Close()

  entitylist, err := openpgp.ReadArmoredKeyRing(to)

  buf := new(bytes.Buffer)
  w, _ := armor.Encode(buf, encryptionType, nil)
  plaintext, _ := openpgp.Encrypt(w, entitylist, nil, nil, nil)

  fmt.Fprintf(plaintext, input)
  plaintext.Close()
  w.Close()

  return buf.String()

}

func sendEmail(body string, email string, gConfig map[string]string) {
  // Set up authentication information.
  auth := smtp.PlainAuth(
    "",
    gConfig["REMOTE_SMTP_USER"],
    gConfig["REMOTE_SMTP_PASS"],
    gConfig["REMOTE_SMTP_HOST"],
  )
  // Connect to the server, authenticate, set the sender and recipient,
  // and send the email all in one step.
  err := smtp.SendMail(
    gConfig["REMOTE_SMTP_HOST"]+":"+gConfig["REMOTE_SMTP_PORT"],
    auth,
    "vuze@elmund.io",
    []string{email},
    []byte(body),
  )
  logging.CheckFatal(err)
}

func exists(path string) (bool, error) {
  _, err := os.Stat(path)
  if err == nil {
    return true, nil
  }
  if os.IsNotExist(err) {
    return false, nil
  }
  return true, err
}
