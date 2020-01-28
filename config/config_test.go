package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDebugMode(t *testing.T) {
	// Setup
	os.Setenv("DEBUG", "1")

	// Tested code
	debug, err := ProvideDebugModeFromEnvironment()

	// Asserts
	assert.NoError(t, err)
	assert.True(t, bool(debug))
}

func TestDebugMode_Default(t *testing.T) {
	// Setup
	os.Unsetenv("DEBUG")

	// Tested code
	debug, err := ProvideDebugModeFromEnvironment()

	// Asserts
	assert.NoError(t, err)
	assert.False(t, bool(debug))
}

func TestDebugMode_Invalid(t *testing.T) {
	// Setup
	os.Setenv("DEBUG", "not a real bool")

	// Tested code
	_, err := ProvideDebugModeFromEnvironment()

	// Asserts
	assert.Error(t, err)
}

func TestUnderflowListenPort(t *testing.T) {
	os.Setenv("UF_PORT", "12345")

	port, err := ProvideUnderflowListenPortFromEnvironment()

	assert.NoError(t, err)
	assert.Equal(t, UnderflowListenPort(12345), port)
}

func TestUnderflowListenPort_Default(t *testing.T) {
	os.Unsetenv("UF_PORT")

	port, err := ProvideUnderflowListenPortFromEnvironment()

	assert.NoError(t, err)
	assert.Equal(t, UnderflowListenPort(9876), port)
}

func TestUnderflowListenPort_Invalid(t *testing.T) {
	os.Setenv("UF_PORT", "12345abc")

	_, err := ProvideUnderflowListenPortFromEnvironment()

	assert.Error(t, err)
}

func TestUnderflowMode(t *testing.T) {
	os.Setenv("UF_MODE", "client")
	mode, err := ProvideUnderflowModeFromEnvironment()
	assert.NoError(t, err)
	assert.Equal(t, ModeClient, mode)

	os.Setenv("UF_MODE", "server")
	mode, err = ProvideUnderflowModeFromEnvironment()
	assert.NoError(t, err)
	assert.Equal(t, ModeServer, mode)
}

func TestUnderflowMode_Invalid(t *testing.T) {
	os.Unsetenv("UF_MODE")
	_, err := ProvideUnderflowModeFromEnvironment()
	assert.Error(t, err)

	os.Setenv("UF_MODE", "not a real mode")
	_, err = ProvideUnderflowModeFromEnvironment()
	assert.Error(t, err)
}

func TestUnderflowMasterAddr(t *testing.T) {
	os.Setenv("UF_MASTER", "uf://abc.def:1234")
	addr := ProvideUnderflowMasterAddrFromEnvironment()
	assert.Equal(t, UnderflowMasterAddr("uf://abc.def:1234"), addr)
}

func TestUnderflowMinionAddrs(t *testing.T) {
	os.Setenv("UF_MINIONS", "uf://abc.def:1234,uf://ghi.jkl:2345,uf://mno.pqr:3456")
	addrs := ProvideUnderflowMinionAddrsFromEnvironment()
	assert.Equal(t, UnderflowMinionAddrs{"uf://abc.def:1234", "uf://ghi.jkl:2345", "uf://mno.pqr:3456"}, addrs)
}
