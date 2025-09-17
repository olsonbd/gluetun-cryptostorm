# Cryptostorm VPN Provider

This provider supports the Cryptostorm VPN service.

## Configuration

Set your environment variables:

```bash
VPN_SERVICE_PROVIDER=cryptostorm
OPENVPN_USER=your_username
OPENVPN_PASSWORD=your_password
```

## Features

- **OpenVPN Support**: TCP port 443, UDP port 1194
- **Encryption**: AES-256-GCM, AES-256-CBC
- **Multiple Locations**: US, Canada, UK, Germany, Netherlands
- **Authentication**: Username/password authentication

## Server Selection

You can select servers by:
- Country (e.g., `SERVER_COUNTRIES=United States`)
- City (e.g., `SERVER_CITIES=New York`)
- Hostname (e.g., `SERVER_HOSTNAMES=newyork.cryptostorm.is`)

## Notes

- This provider uses placeholder certificates in the current implementation
- In production, you would need to replace the CA certificate and TLS auth key with actual Cryptostorm credentials
- Server list is currently static but could be enhanced to fetch from Cryptostorm's API

## Cryptostorm Documentation

For more information about Cryptostorm VPN service, visit: https://cryptostorm.is/nix