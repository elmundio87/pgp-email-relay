# CHANGELOG

## v0.3 (2015-11-1)

### Fix release

- The previous release had a nasty bug preventing the message body from showing
 if the attachment was enabled. This has been fixed

No new features, just a bug fix

## v0.2 (2015-11-1)

### Initial attachment support

- Make use of forked Email package for easier sending of attachments
- Email body can be attached for easier reading on mobile devices
- Default configuration should require less modifying
- Key generation script only requires typing the password once

I will try and get attachment relaying into the next release

## v0.1 (2015-10-31)

### Initial Release

- Email relaying
- Email encryption
- Keys downloaded from keyserver
- Error handling of malformed emails

Doesn't support attachments yet