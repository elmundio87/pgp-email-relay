#!/bin/bash
openssl genrsa -des3 -out server.key 1024
openssl req -new -key server.key -out server.csr
openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt
mv server.key server.key.secure
openssl rsa -in server.key.secure -out server.key

