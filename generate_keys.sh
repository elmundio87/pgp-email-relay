#!/bin/bash

echo "Creating server key and certificate"
echo "Enter a password: (Output Suppressed)"
read -s password

echo ${password}
openssl genrsa -des3 -out server.key -passout pass:${password} 1024 
openssl req -new -key server.key -out server.csr -passin pass:${password}
openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt -passin pass:${password}
mv server.key server.key.secure
echo "Decrypting key for the server to use it"
openssl rsa -in server.key.secure -out server.key -passin pass:${password}

echo "=============================================================================="
echo "You now have a server.key and server.cert. Edit the GSMTP_PRV_KEY and"
echo "GSMTP_PUB_KEY properties in the config file if you want to store these keys"
echo "outside of the installation directory"
echo "=============================================================================="