package flags

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/versent/saml2aws/pkg/cfg"
)

func TestRoleSuppliedIsTrueWithNonEmptyRole(t *testing.T) {

	commonFlags := &CommonFlags{RoleArn: "arn:aws:iam::456456456456/role:myrole"}

	expected := true
	actual := commonFlags.RoleSupplied()

	assert.Equal(t, expected, actual)
}

func TestRoleSuppliedIsFalseWithEmptyRole(t *testing.T) {

	commonFlags := &CommonFlags{RoleArn: ""}

	expected := false
	actual := commonFlags.RoleSupplied()

	assert.Equal(t, expected, actual)
}

func TestOverrideAllFlags(t *testing.T) {

	commonFlags := &CommonFlags{
		IdpProvider:          "ADFS",
		MFA:                  "mymfa",
		SkipVerify:           true,
		URL:                  "https://id.example.com",
		Username:             "myuser",
		AmazonWebservicesURN: "urn:amazon:webservices",
	}
	idpa := &cfg.IDPAccount{
		Provider:             "Ping",
		MFA:                  "none",
		SkipVerify:           false,
		URL:                  "https://id.test.com",
		Username:             "test123",
		AmazonWebservicesURN: "urn:govcloud:webservices",
	}

	expected := &cfg.IDPAccount{
		Provider:             "ADFS",
		MFA:                  "mymfa",
		SkipVerify:           true,
		URL:                  "https://id.example.com",
		Username:             "myuser",
		AmazonWebservicesURN: "urn:amazon:webservices",
	}
	ApplyFlagOverrides(commonFlags, idpa)

	assert.Equal(t, expected, idpa)
}

func TestNoOverrides(t *testing.T) {

	commonFlags := &CommonFlags{
		IdpProvider:          "",
		MFA:                  "",
		SkipVerify:           false,
		URL:                  "",
		Username:             "",
		AmazonWebservicesURN: "",
	}
	idpa := &cfg.IDPAccount{
		Provider:             "Ping",
		MFA:                  "none",
		SkipVerify:           false,
		URL:                  "https://id.test.com",
		Username:             "test123",
		AmazonWebservicesURN: "urn:govcloud:webservices",
	}

	expected := &cfg.IDPAccount{
		Provider:             "Ping",
		MFA:                  "none",
		SkipVerify:           false,
		URL:                  "https://id.test.com",
		Username:             "test123",
		AmazonWebservicesURN: "urn:govcloud:webservices",
	}
	ApplyFlagOverrides(commonFlags, idpa)

	assert.Equal(t, expected, idpa)
}
