package updater

import (
	"context"
	"testing"

	"github.com/qdm12/gluetun/internal/constants/vpn"
)

func TestUpdater_FetchServers(t *testing.T) {
	t.Parallel()

	updater := &Updater{}
	
	servers, err := updater.FetchServers(context.Background(), 1)
	if err != nil {
		t.Fatalf("FetchServers failed: %v", err)
	}

	if len(servers) == 0 {
		t.Fatal("No servers returned")
	}

	// Verify first server has required fields
	server := servers[0]
	if server.VPN != vpn.OpenVPN {
		t.Errorf("Expected VPN type %s, got %s", vpn.OpenVPN, server.VPN)
	}

	if server.Hostname == "" {
		t.Error("Server hostname is empty")
	}

	if len(server.IPs) == 0 {
		t.Error("Server has no IPs")
	}

	if !server.TCP && !server.UDP {
		t.Error("Server supports neither TCP nor UDP")
	}
}