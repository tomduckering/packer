package provisioner

import (
	"fmt"
	"strings"
)

const UnixOSType = "unix"
const WindowsOSType = "windows"
const DefaultOSType = UnixOSType

type guestOSTypeCommand struct {
	chmodExecutable string
	mkdir           string
	removeDir       string
}

var guestOSTypeCommands = map[string]guestOSTypeCommand{
	UnixOSType: guestOSTypeCommand{
		chmodExecutable: "chmod +x '%s'",
		mkdir:           "mkdir -p '%s'",
		removeDir:       "rm -rf '%s'",
	},
	WindowsOSType: guestOSTypeCommand{
		chmodExecutable: "echo 'skipping chmod %s'", // no-op
		mkdir:           "New-Item -ItemType directory -Force -ErrorAction SilentlyContinue -Path %s",
		removeDir:       "rm %s -recurse -force",
	},
}

type GuestCommands struct {
	GuestOSType string
}

func NewGuestCommands(osType string) (*GuestCommands, error) {
	_, ok := guestOSTypeCommands[osType]
	if !ok {
		return nil, fmt.Errorf("Invalid osType: \"%s\"", osType)
	}
	return &GuestCommands{GuestOSType: osType}, nil
}

func (g *GuestCommands) ChmodExecutable(path string) string {
	return fmt.Sprintf(g.commands().chmodExecutable, g.escapePath(path))
}

func (g *GuestCommands) CreateDir(path string) string {
	return fmt.Sprintf(g.commands().mkdir, g.escapePath(path))
}

func (g *GuestCommands) RemoveDir(path string) string {
	return fmt.Sprintf(g.commands().removeDir, g.escapePath(path))
}

func (g *GuestCommands) commands() guestOSTypeCommand {
	return guestOSTypeCommands[g.GuestOSType]
}

func (g *GuestCommands) escapePath(path string) string {
	if g.GuestOSType == UnixOSType {
		return path
	}
	return strings.Replace(path, " ", "` ", -1)
}
