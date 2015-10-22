#!/bin/bash

export bin=goguerrilla
export installdir=/etc/pgp_email_relay
export GOPATH=/tmp/go

if [[ $EUID -ne 0 ]]; then
  echo "You must be a root user" 2>&1
  exit 1
fi

if [ ! -f ${bin} ]; then
	echo "Binary missing, running build script..."
	./build.sh
fi

if [ ! -f ${bin}.conf ]; then
	echo "Missing configuration file: ${bin}.conf"
	exit 1
fi

echo "Creating install directory ${installdir}"
mkdir -p $installdir

if [ ! -f server.key ]; then
	echo "server.key missing, generating certificate and key"
	./generate_keys.sh
fi

echo "Moving files to ${installdir}"
cp $bin server.{key,crt} ${bin}.conf $installdir/
echo "Done!"
