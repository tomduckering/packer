package provisioner

import (
	"fmt"
)

const UnixOSType = "unix"
const WindowsOSType = "windows"
const DefaultOSType = UnixOSType

type guestOSTypeCommand struct {
	chmodExecutable string
	mkdir           string
}

var guestOSTypeCommands = map[string]guestOSTypeCommand{
	UnixOSType: guestOSTypeCommand{
		chmodExecutable: "chmod +x '%s'",
		mkdir:           "mkdir -p '%s'",
	},
	WindowsOSType: guestOSTypeCommand{
		chmodExecutable: "echo 'skipping chmod %s'", // no-op
		mkdir:           "New-Item -ItemType directory -Force -ErrorAction SilentlyContinue -Path '%s'",
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
	return fmt.Sprintf(g.commands().chmodExecutable, path)
}

func (g *GuestCommands) Mkdir(path string) string {
	return fmt.Sprintf(g.commands().mkdir, path)
}

func (g *GuestCommands) commands() guestOSTypeCommand {
	return guestOSTypeCommands[g.GuestOSType]
}
