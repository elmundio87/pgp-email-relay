package pgp_encrypt

import (
  "bytes"
  "fmt"
  "github.com/cryptix/go/logging"
  "github.com/elmundio87/pgp-email-relay/email"
  "github.com/elmundio87/pgp-email-relay/publickey"
  "golang.org/x/crypto/openpgp"
  "golang.org/x/crypto/openpgp/armor"
  "html/template"
  "io"
  "io/ioutil"
  "net/mail"
  "net/smtp"
  "os"
  "path"
  "strings"
  "time"
)

const encryptionType = "PGP MESSAGE"

func HandleMail(client_data string, client_rcpt_to string, gConfig map[string]string) {

  var to = client_rcpt_to[1 : len(client_rcpt_to)-1]

  addresses := strings.Split(to, ",")

  for _, address := range addresses {

    emailData := client_data[:len(client_data)-4]
    msg, err := mail.ReadMessage(bytes.NewBuffer([]byte(emailData)))

    if err != nil {
      sendErrorReport(err, address, gConfig)
      return
    }

    headers := make(map[string]string)
    for key, value := range msg.Header {
      headers[key] = value[0]
    }

    body, _ := ioutil.ReadAll(msg.Body)

    encryptedBody := encrypt(string(body), address, gConfig)

    sendEmail(headers, encryptedBody, address, gConfig)

  }

}

// http://stackoverflow.com/a/31742265
func sendErrorReport(err error, address string, gConfig map[string]string) {

  bodyTemplate := `
Error: {{.Error}}
Time: {{.Time}}

Open issues: https://github.com/elmundio87/pgp-email-relay/issues

Feel free to submit a bug report.

`
  data := map[string]interface{}{
    "Error":   err.Error(),
    "Address": address,
    "Time":    time.Now(),
  }

  headers := make(map[string]string)
  headers["Subject"] = "Crash Report"
  headers["From"] = gConfig["REMOTE_SMTP_USER"]

  t := template.Must(template.New("email").Parse(bodyTemplate))
  buf := &bytes.Buffer{}
  if err := t.Execute(buf, data); err != nil {
    panic(err)
  }
  body := buf.String()

  sendEmail(headers, body, address, gConfig)

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

func sendEmail(headers map[string]string, body string, address string, gConfig map[string]string) {

  host := gConfig["REMOTE_SMTP_HOST"]
  port := gConfig["REMOTE_SMTP_PORT"]
  user := gConfig["REMOTE_SMTP_USER"]
  password := gConfig["REMOTE_SMTP_PASS"]

  m := email.NewMessage(headers["Subject"], body)

  fromAddress := "Unknown"
  fromHeader, ok := headers["From"]
  if ok {
    from, _ := mail.ParseAddress(fromHeader)
    fromAddress = from.Address
  }

  to, _ := mail.ParseAddress(address)
  toAddress := to.Address

  m.From = fromAddress
  m.To = []string{toAddress}

  if gConfig["PGP_ATTACH_BODY"] == "Y" {
    m.AttachData("message.asc", []byte(body))
  }

  err := email.Send(host+":"+port, smtp.PlainAuth("", user, password, host), m)
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
