
Example credential management.  Needs work to be made generic.

Interacts with Google APIs to download credentials.

# Initial setup

Usage:

Authenticate with Google and download a Google OAUTH2 token to the
.credentials-cli-auth file.  You should keep this token safe, and treat it as a
password:
```
  credentials-cli -A 
```
Outputs a URL which should be loaded into a web browser on the same host.
Follow prompts to complete authentication:

# Access credentials

List credentials you own:
```
  credentials-cli -L
```

Show what formats are available:

```
  credentials-cli -F CRED-NAME
```

Downloads a credential.  VPN creds are stored as a set of .ovpn
configuration files.  Web creds are stored as a .p12 file.  The P12
file password is shown on stdout:
```
  credentials-cli -D CRED-NAME
```

# Other formats

## PEM

Save as PEM:
```
  credentials-cli -D CRED-NAME -f pem
```
If cred is a web cred or VPN cred, save as a bundle of PEMs instead of a
P12 or OpenVPN configuration file.

## MacOS / iOS

Save as .mobileconfig which configures a MacOS or iOS device for IPsec VPN:
```
  credentials-cli -D CRED-NAME -f mobileconfig-ipsec
```

Save as .mobileconfig which configures an iOS device for OpenVPN.  Only tested
with OpenVPN Connect on iPad:
```
  credentials-cli -D CRED-NAME -f mobileconfig-ovpn
```

Save as .mobileconfig which configures an iOS with web credentials:
```
  credentials-cli -D CRED-NAME -f mobileconfig
```

## Android

Save as .sswan which configures an Android device for StrongSwan.
```
  credentials-cli -D CRED-NAME -f sswan
```

## Linux

Save as a bunch of StrongSwan configuration files which can be used to
configure StrongSwan natively.  Use this for Linux.
```
  credentials-cli -D CRED-NAME -f sswan-basic
```

## Admin access from a service account

If you're working with a service account, the JSON key file can be specified
with the -k option:
```
  credentials-cli -L -k JSON-KEY-FILE
  credentials-cli -D CRED-NAME -k JSON-KEY-FILE
```

# Mobileconfig signing

If you're creating a mobileconfig file, you can sign it with a signing
key/cert:
```
  credentials-cli -D CRED-NAME -f mobileconfig -S -C signing.crt -K signing.pem
  credentials-cli -D CRED-NAME -f mobileconfig-ipsec -S -C signing.crt -K signing.pem
```

# MacOS admin

While you're enjoying the command-line experience, did you know you can
add a VPN profile on MacOS thus:
```
    sudo profiles -I -U mark -f -F mark-mac.mobileconfig
```
and list profiles:
```
  profiles -L
```
and remove a profile:
```
  sudo profiles -U mark -D mark-mac.device.local -f
```
On MacOS you can also import a P12 file into the keychain:
```
  security import web-cert.p12 -P my-p12-password
```

# Credential creation

```
  credentials-cli --create-vpn -u user@example.com -i bunchy-mac
  credentials-cli --create-web -u user@example.com -i 'User Smith (ABC)'
  credentials-cli --revoke-vpn -u user@example.com -i bunchy-mac
  credentials-cli --revoke-web -u user@example.com -i 'User Smith (ABC)'
  credentials-cli --revoke-all -u user@example.com
```
