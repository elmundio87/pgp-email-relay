
PGP Email Relay
====================

A simple SMTP relay that will encrypt any emails that it receives, and will relay
the message via a remote SMTP server.

The intended purpose of this project is to allow an application that sends automated emails to encrypt them before they reach the recipient, without modifying the source code of the application itself.

Most of the SMTP code is based on [Go Guerrilla](https://github.com/flashmob/go-guerrilla) by Flashmob

Public keys are downloaded from a keyserver (the keyserver URL is configurable). You can also place your own keys into the key cache folder, these will not be overwritten.


Building
===========================

To build, you will need do the following;

1. Install Golang 
2. Run the build script ```./build.sh```


Before you run the server
===========================

1. Rename smtp.conf.sample to smtp.conf and modify accordingly
2. Run the key generation script ```./generate_keys.sh```


Configuration
============================================
The configuration is in strict JSON format. Here is an annotated configuration.
Copy smtp.conf.sample to smtp.conf

| Config Option  | Purpose  | Example |
|---|---|---|
|REMOTE_SMTP_USER|Remote SMTP server username| user@remotehost.com |
|REMOTE_SMTP_PASS|Remote SMTP server password| p@55w0rd |
|REMOTE_SMTP_HOST|Remote SMTP server hostname| smtp.remotehost.com |
|REMOTE_SMTP_PORT|Which port the remote SMTP server is listening on| 25 |
|PGP_KEYSERVER|The PGP keyserver that will be used to cache keys from| pgp.mit.edu |
|PGP_KEYSERVER_QUERY|The URL query that is used to search for keys| /pks/lookup?op=index&exact=on&search= |
|PGP_KEY_FOLDER|Where keys are cached| /tmp/key_cache |
|PGP_ATTACH_BODY|Attach the contents of the encrypted email body. Makes it easier to decrypt on mobile clients.| Y or N |
|GM_ALLOWED_HOSTS|Which domains accept mail| localhost, mail.yourhost.com |
|GM_PRIMARY_MAIL_HOST|Given in the SMTP greeting| mail.yourhost.com |
|GSMTP_HOST_NAME|Given in the SMTP greeting| mail.yourhost.com |
|GSMTP_LOG_FILE"|Not used yet| N/A |
|GSMTP_MAX_SIZE|Max size of DATA command| 15728640 |
|GSMTP_PRV_KEY|Private key for TLS|server.key|
|GSMTP_PUB_KEY|Public key for TLS|server.cert|
|GSMTP_TIMEOUT|TCP connection timeout in seconds|100|
|GSMTP_VERBOSE|Enable Debugging|Y or N|
|GSTMP_LISTEN_INTERFACE|What IP:PORT to listen on|127.0.0.1:25|
|GM_MAX_CLIENTS|Max clients that can be handled|500|
|NGINX_AUTH_ENABLED| Enable Nginx authentication|Y or N|
|NGINX_AUTH|If using Nginx proxy, choose an ip and port to serve Auth requsts for Nginx|127.0.0.1:8025|
|SGID|Group id of the user from /etc/passwd|508|
|GUID|Uid from /etc/passwd|504|

Using Nginx as a proxy
=========================================================
Nginx can be used to proxy SMTP traffic for GoGuerrilla SMTPd

Why proxy SMTP?

 *	Terminate TLS connections: Golang is not there yet when it comes to TLS.
At present, only a partial implementation of TLS is provided (as of Nov 2012). 
OpenSSL on the other hand, used in Nginx, has a complete implementation of
SSL v2/v3 and TLS protocols.
 *	Could be used for load balancing and authentication in the future.

 1.	Compile nginx with --with-mail --with-mail_ssl_module

 2.	Configuration:

	
		mail {
	        auth_http 127.0.0.1:8025/; # This is the URL to GoGuerrilla's http service which tells Nginx where to proxy the traffic to 								
	        server {
	                listen  15.29.8.163:25;
	                protocol smtp;
	                server_name  ak47.example.com;
	
	                smtp_auth none;
	                timeout 30000;
					smtp_capabilities "SIZE 15728640";
					
					# ssl default off. Leave off if starttls is on
	                #ssl                  on;
	                ssl_certificate      /etc/ssl/certs/ssl-cert-snakeoil.pem;
	                ssl_certificate_key  /etc/ssl/private/ssl-cert-snakeoil.key;
	                ssl_session_timeout  5m;
	                ssl_protocols  SSLv2 SSLv3 TLSv1;
	                ssl_ciphers  HIGH:!aNULL:!MD5;
	                ssl_prefer_server_ciphers   on;
					# TLS off unless client issues STARTTLS command
	                starttls on;
	                proxy on;
	        }
		}
	
			
Assuming that Guerrilla SMTPd has the following configuration settings:

	"GSMTP_MAX_SIZE"		  "15728640",
	"NGINX_AUTH_ENABLED":     "Y",
	"NGINX_AUTH":             "127.0.0.1:8025", 


Starting / Command Line usage
==========================================================

All command line arguments are optional

	-config="smtp.conf": Path to the configuration file
	 -if="": Interface and port to listen on, eg. 127.0.0.1:2525
	 -v="n": Verbose, [y | n]

Starting from the command line (example)

	/usr/bin/nohup /home/elmundio87/pgp-email-relaay -config=/home/elmundio87/smtp.conf 2>&1 &

This will place goguerrilla in the background and continue running
