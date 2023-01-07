package infakt

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const (
	Token = "12345678-1234-1234-1234-1234567890ab"
)

func TestResourceProvider_Configure(t *testing.T) {
	infaktProvider := Provider()

	raw := map[string]interface{}{
		"token": Token,
	}

	err := infaktProvider.Configure(terraform.NewResourceConfigRaw(raw))
	if err != nil {
		t.Fatalf("err: %s", err)
	}
}
