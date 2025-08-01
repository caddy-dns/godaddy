package godaddy

import (
	"fmt"
	"testing"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/godaddy"
)

func TestUnmarshalCaddyFile(t *testing.T) {
	tests := []string{
		`godaddy {
			api_token theapitoken
		}`}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			// given
			dispenser := caddyfile.NewTestDispenser(tc)
			p := Provider{&godaddy.Provider{}}
			// when
			err := p.UnmarshalCaddyfile(dispenser)
			// then
			if err != nil {
				t.Errorf("UnmarshalCaddyfile failed with %v", err)
				return
			}

			expectedAPIToken := "theapitoken"
			actualAPIToken := p.Provider.APIToken
			if expectedAPIToken != actualAPIToken {
				t.Errorf("Expected APIToken to be '%s' but got '%s'", expectedAPIToken, actualAPIToken)
			}
		})
	}
}
