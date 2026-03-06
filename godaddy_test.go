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
		}`,
		`godaddy theapitoken`,
	}

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

func TestUnmarshalCaddyFileErrors(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "missing token",
			input: `godaddy`,
		},
		{
			name: "duplicate token",
			input: `godaddy theapitoken {
				api_token theapitoken
			}`,
		},
		{
			name:  "too many inline args",
			input: `godaddy theapitoken extraarg`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dispenser := caddyfile.NewTestDispenser(tc.input)
			p := Provider{&godaddy.Provider{}}
			err := p.UnmarshalCaddyfile(dispenser)
			if err == nil {
				t.Errorf("Expected error but got nil for input: %s", tc.input)
			}
		})
	}
}
