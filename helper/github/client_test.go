package github

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient_Success(t *testing.T) {
	key, err := os.ReadFile("testdata/test.key")
	if err != nil {
		t.Fatal(err)
	}

	client, err := NewClient(12345, string(key), "https://api.github.com/", false)
	assert.NoError(t, err)
	assert.NotNil(t, client)
}

func TestNewClient_Failure(t *testing.T) {
	client, err := NewClient(12345, "", "https://api.github.com/", false)
	assert.Error(t, err)
	assert.Nil(t, client)
}

func TestClientInstallation(t *testing.T) {
	key, err := os.ReadFile("testdata/test.key")
	if err != nil {
		t.Fatal(err)
	}

	client, err := NewClient(12345, string(key), "https://api.github.com/", false)
	assert.NoError(t, err)

	installation := client.Installation(12345)
	assert.NotNil(t, installation)
}
