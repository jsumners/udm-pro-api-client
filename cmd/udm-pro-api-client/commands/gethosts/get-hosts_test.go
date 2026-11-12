package gethosts

import (
	"bytes"
	"testing"

	"github.com/jsumners/udm-pro-api-client/cmd/udm-pro-api-client/internal/config"
	"github.com/stretchr/testify/assert"
)

func Test_outputHostAliases(t *testing.T) {
	input := []config.HostAlias{
		{Name: "foo", IpAddress: "1.1.1.1"},
		{Name: "bar", IpAddress: "2.2.2.2"},
	}
	found := bytes.Buffer{}
	outputHostAliases(input, &found)

	expected := "# host aliases\n1.1.1.1\tfoo\n2.2.2.2\tbar\n\n"
	assert.Equal(t, expected, found.String())
}
