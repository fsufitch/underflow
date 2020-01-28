package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// DebugMode is whether a debug state is set
type DebugMode bool

// UnderflowListenPort is the port that a listening Underflow Server uses
type UnderflowListenPort int

// UnderflowMasterAddr is the address that a client Underflow connects to
type UnderflowMasterAddr string

// UnderflowMinionAddrs is a list of addresses an Underflow master can find its minions at
type UnderflowMinionAddrs []string

// UnderflowMode is the mode that the Underflow process is running in
type UnderflowMode string

// Values for the UnderflowMode type
const (
	ModeServer = UnderflowMode("server")
	ModeClient = UnderflowMode("client")
)

// ProvideDebugModeFromEnvironment creates a DebugMode based on the value in the DEBUG env var
func ProvideDebugModeFromEnvironment() (DebugMode, error) {
	debugString, ok := os.LookupEnv("DEBUG")
	if !ok {
		return false, nil
	}
	mode, err := strconv.ParseBool(debugString)
	return DebugMode(mode), err
}

// ProvideUnderflowListenPortFromEnvironment creates a UnderflowListenPort from the environment, defaulting to 9876 when missing
func ProvideUnderflowListenPortFromEnvironment() (UnderflowListenPort, error) {
	portString, ok := os.LookupEnv("UF_PORT")
	if !ok {
		portString = "9876"
	}
	port, err := strconv.ParseInt(portString, 0, 0)
	return UnderflowListenPort(port), err
}

// ProvideUnderflowModeFromEnvironment provides the UnderflowMode from the environment, erroring when missing
func ProvideUnderflowModeFromEnvironment() (UnderflowMode, error) {
	mode := UnderflowMode(os.Getenv("UF_MODE"))
	switch mode {
	case "":
		return "", errors.New("UF_MODE not set")
	case ModeServer, ModeClient:
		return mode, nil
	default:
		return "", fmt.Errorf("invalid UF_MODE: %s", mode)
	}
}

// ProvideUnderflowMasterAddrFromEnvironment provides the UnderflowMasterAddr from the environment
func ProvideUnderflowMasterAddrFromEnvironment() UnderflowMasterAddr {
	return UnderflowMasterAddr(os.Getenv("UF_MASTER"))
}

// ProvideUnderflowMinionAddrsFromEnvironment provides the UnderflowMinionAddrs from the environment
func ProvideUnderflowMinionAddrsFromEnvironment() UnderflowMinionAddrs {
	return strings.Split(os.Getenv("UF_MINIONS"), ",")

}
