#!/usr/bin/ruby

require 'net/smtp'

SMTPHOST = 'localhost'
FROM = '"Your Email" <youremail@replace.example.com>'

def send(to, subject, message)
body = <<EOF
From: #{FROM}
To: #{to}
Subject: #{subject}

#{message}
EOF
  Net::SMTP.start(SMTPHOST) do |smtp|
    smtp.send_message body, FROM, to
  end
end
    
send('elmundio1987@gmail.com', 'testing', 'This is a message!')
