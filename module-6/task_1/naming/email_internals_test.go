package naming

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"testing_tasks/task_1/naming/testdata"
)

func Test_addDomain(t *testing.T) {
	in := "dummy"
	domain := "corp.com"
	out := "dummy@corp.com"
	eml, err := addDomain(in, domain)
	require.NoError(t, err)
	assert.Equal(t, out, eml)
}

func Test_addDomain_SameName(t *testing.T) {
	in := "dummy@corp.com"
	domain := "corp.com"
	out := "dummy@corp.com"
	eml, err := addDomain(in, domain)
	require.NoError(t, err)
	assert.Equal(t, out, eml)
}

func Test_addDomain_Err(t *testing.T) {
	in := "dummy@other.com"
	domain := "corp.com"
	_, err := addDomain(in, domain)
	assert.ErrorIs(t, err, ErrAlreadyEmailWrongDomain)
}

func Test_validateName_ok(t *testing.T) {
	for _, item := range testdata.CrewData {
		err := validateName(item.InName)
		assert.NoError(t, err, "Must not have errors: '%s'", item.InName)
	}
}

func Test_validateName_Err(t *testing.T) {
	err := validateName("")
	assert.ErrorIs(t, err, ErrNameIsEmpty)
}

func Test_NameToEmail_ok(t *testing.T) {
	for _, item := range testdata.CrewData {
		out, err := NameToEmail(item.InName)
		require.NoError(t, err)
		assert.Equal(t, item.OutEmail, out)
	}
}

func Test_NameToEmail_ErrEmptyName(t *testing.T) {
	in := ""
	_, err := NameToEmail(in)
	assert.ErrorIs(t, err, ErrNameIsEmpty)
}

func Test_NameFromEmail_ok(t *testing.T) {
	for _, item := range testdata.CrewData {
		out, err := NameFromEmail(item.OutEmail)
		require.NoError(t, err)
		assert.Equal(t, out, item.OutName)
	}
}

func Test_NameFromEmail_Err(t *testing.T) {
	in := "dummy_acme.com"
	_, err := NameFromEmail(in)
	assert.ErrorIs(t, err, ErrEmailInvalidFormat)
}
