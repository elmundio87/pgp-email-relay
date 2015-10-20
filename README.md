
PGP Email Relay
====================

A simple SMTP relay that will encrypt any emails that it recieves, and will relay
the message via a remote SMTP server.

Most of the SMTP code is based on [Go Guerrilla](https://github.com/flashmob/go-guerrilla) by Flashmob

Building
===========================

To build, you will need do the following;

1) Install Golang 
2) Run the build script ```./build.sh```


Before you run the server
===========================

1) Rename goguerrilla.conf.sample to goguerrilla.conf and modify accordinly
2) Run the key generation script ```./generate_keys.sh```


Configuration
============================================
The configuration is in strict JSON format. Here is an annotated configuration.
Copy goguerrilla.conf.sample to goguerrilla.conf


	{
	    "GM_ALLOWED_HOSTS":"example.com,sample.com,foo.com,bar.com", // which domains accept mail
	    "GM_MAIL_TABLE":"new_mail", // name of new email table
	    "GM_PRIMARY_MAIL_HOST":"mail.example.com", // given in the SMTP greeting
	    "GSMTP_HOST_NAME":"mail.example.com", // given in the SMTP greeting
	    "GSMTP_LOG_FILE":"/dev/stdout", // not used yet
	    "GSMTP_MAX_SIZE":"131072", // max size of DATA command
	    "GSMTP_PRV_KEY":"/etc/ssl/private/example.com.key", // private key for TLS
	    "GSMTP_PUB_KEY":"/etc/ssl/certs/example.com.crt", // public key for TLS
	    "GSMTP_TIMEOUT":"100", // tcp connection timeout
	    "GSMTP_VERBOSE":"N", // set to Y for debugging
	    "GSTMP_LISTEN_INTERFACE":"5.9.7.183:25",
	    "GM_MAX_CLIENTS":"500", // max clients that can be handled
			"NGINX_AUTH_ENABLED":"N",// Y or N
			"NGINX_AUTH":"127.0.0.1:8025", // If using Nginx proxy, choose an ip and port to serve Auth requsts for Nginx
	    "SGID":"508",// group id of the user from /etc/passwd
			"GUID":"504" // uid from /etc/passwd
			"REMOTE_SMTP_USER":"user@remotehost.com", //user to log into remote SMTP server
    	"REMOTE_SMTP_PASS":"password", //password of remote SMTP user 
    	"REMOTE_SMTP_HOST":"smtp.remotehost.com", //remote SMTP server host
    	"REMOTE_SMTP_PORT":"25" //which port to use when connecting to the remote SMTP server
	}

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

	-config="goguerrilla.conf": Path to the configuration file
	 -if="": Interface and port to listen on, eg. 127.0.0.1:2525
	 -v="n": Verbose, [y | n]

Starting from the command line (example)

	/usr/bin/nohup /home/mike/goguerrilla -config=/home/mike/goguerrilla.conf 2>&1 &

This will place goguerrilla in the background and continue running

You may also put another process to watch your goguerrilla process and re-start it
if something goes wrong.