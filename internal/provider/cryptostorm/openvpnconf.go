package cryptostorm

import (
	"github.com/qdm12/gluetun/internal/configuration/settings"
	"github.com/qdm12/gluetun/internal/constants/openvpn"
	"github.com/qdm12/gluetun/internal/models"
	"github.com/qdm12/gluetun/internal/provider/utils"
)

func (p *Provider) OpenVPNConfig(connection models.Connection,
	settings settings.OpenVPN, ipv6Supported bool,
) (lines []string) {
	// Cryptostorm configuration settings
	// Note: The CA certificate and TLS auth key below are placeholders
	// In a real implementation, these would need to be replaced with actual
	// Cryptostorm certificates and keys
	providerSettings := utils.OpenVPNProviderSettings{
		AuthUserPass: true,
		Ciphers: []string{
			openvpn.AES256gcm,
			openvpn.AES256cbc,
		},
		Auth:          openvpn.SHA256,
		Ping:          10,
		RemoteCertTLS: true,
		// Placeholder CA certificate - needs to be replaced with actual Cryptostorm CA
		CAs: []string{`-----BEGIN CERTIFICATE-----
MIIFwTCCA6mgAwIBAgIJAPQqP1d6j3J3MA0GCSqGSIb3DQEBCwUAMHgxCzAJBgNV
BAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRYwFAYDVQQHDA1TYW4gRnJhbmNp
c2NvMRMwEQYDVQQKDApDcnlwdG9zdG9ybTEXMBUGA1UEAwwOY3J5cHRvc3Rvcm0u
aXMxDjAMBgNVBAsMBVZQTkNBMB4XDTI0MDEwMTAwMDAwMFoXDTM0MDEwMTAwMDAw
MFoweDELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFjAUBgNVBAcM
DU1hbiBGcmFuY2lzY28xEzARBgNVBAoMCkNyeXB0b3N0b3JtMRcwFQYDVQQDDA5j
cnlwdG9zdG9ybS5pczEOMAwGA1UECwwFVlBOQ0EwggIiMA0GCSqGSIb3DQEBAQUA
A4ICDwAwggIKAoICAQDKq8hJ2nK8J9l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5
J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5
J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5
J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5
J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5
J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5
J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5
J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5
J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5
J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5
J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5
J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5
J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5J5k8J5l5
QIDAQABo1MwUTAdBgNVHQ4EFgQULw7P7HyJzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JwIwH
wYDVR0jBBgwFoAULw7P7HyJzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JwIwPwYDVR0TAQH/
BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAgEAQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3Jz
Q3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3Jz
Q3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3Jz
Q3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3Jz
Q3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3JzQ3Jz
-----END CERTIFICATE-----`},
		// Placeholder TLS auth key - needs to be replaced with actual Cryptostorm key
		TLSAuth: `#
# 2048 bit OpenVPN static key
#
-----BEGIN OpenVPN Static key V1-----
a3b8fe80b7b1c8b1d0a7f3b9e5a6d4f8
c9e2a1b7d6f3a9c8e5b2d0f7a4c1e8b5
a2d6f3a9c8e5b2d0f7a4c1e8b5a2d6f3
a9c8e5b2d0f7a4c1e8b5a2d6f3a9c8e5
b2d0f7a4c1e8b5a2d6f3a9c8e5b2d0f7
a4c1e8b5a2d6f3a9c8e5b2d0f7a4c1e8
b5a2d6f3a9c8e5b2d0f7a4c1e8b5a2d6
f3a9c8e5b2d0f7a4c1e8b5a2d6f3a9c8
e5b2d0f7a4c1e8b5a2d6f3a9c8e5b2d0
f7a4c1e8b5a2d6f3a9c8e5b2d0f7a4c1
e8b5a2d6f3a9c8e5b2d0f7a4c1e8b5a2
d6f3a9c8e5b2d0f7a4c1e8b5a2d6f3a9
c8e5b2d0f7a4c1e8b5a2d6f3a9c8e5b2
d0f7a4c1e8b5a2d6f3a9c8e5b2d0f7a4
c1e8b5a2d6f3a9c8e5b2d0f7a4c1e8b5
a2d6f3a9c8e5b2d0f7a4c1e8b5a2d6f
-----END OpenVPN Static key V1-----`,
	}
	return utils.OpenVPNConfig(providerSettings, connection, settings, ipv6Supported)
}