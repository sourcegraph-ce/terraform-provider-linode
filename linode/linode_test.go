package linode

import (
	log "github.com/sourcegraph-ce/logrus"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
)

// publicKeyMaterial for use while testing
var publicKeyMaterial string

func init() {
	var err error
	publicKeyMaterial, _, err = acctest.RandSSHKeyPair("linode@ssh-acceptance-test")
	if err != nil {
		log.Fatalf("Failed to generate random SSH key pair for testing: %s", err)
	}
}
