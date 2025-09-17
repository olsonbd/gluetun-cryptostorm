package cryptostorm

import (
	"github.com/qdm12/gluetun/internal/configuration/settings"
	"github.com/qdm12/gluetun/internal/models"
	"github.com/qdm12/gluetun/internal/provider/utils"
)

func (p *Provider) GetConnection(selection settings.ServerSelection, ipv6Supported bool) (
	connection models.Connection, err error,
) {
	// Cryptostorm supports OpenVPN on TCP 443 and UDP 1194
	defaults := utils.NewConnectionDefaults(443, 1194, 0) // No WireGuard support
	return utils.GetConnection(p.Name(),
		p.storage, selection, defaults, ipv6Supported, p.randSource)
}