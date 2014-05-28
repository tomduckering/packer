package provisioner

import (
	"testing"
)

func TestNewGuestCommands(t *testing.T) {
	_, err := NewGuestCommands("Amiga")
	if err == nil {
		t.Fatalf("Should have returned an err for unsupported OS type")
	}
}

func TestMkdir(t *testing.T) {
	guestCmd, err := NewGuestCommands(UnixOSType)
	if err != nil {
		t.Fatalf("Failed to create new GuestCommands for OS: %s", UnixOSType)
	}
	cmd := guestCmd.Mkdir("/tmp/tempdir")
	if cmd != "mkdir -p '/tmp/tempdir'" {
		t.Fatalf("Unexpected Unix mkdir cmd: %s", cmd)
	}

	guestCmd, err = NewGuestCommands(WindowsOSType)
	if err != nil {
		t.Fatalf("Failed to create new GuestCommands for OS: %s", WindowsOSType)
	}
	cmd = guestCmd.Mkdir("C:\\Windows\\Temp\\tempdir")
	if cmd != "New-Item -ItemType directory -Force -ErrorAction SilentlyContinue -Path 'C:\\Windows\\Temp\\tempdir'" {
		t.Fatalf("Unexpected Windows mkdir cmd: %s", cmd)
	}
}

func TestChmodExecutable(t *testing.T) {
	guestCmd, err := NewGuestCommands(UnixOSType)
	if err != nil {
		t.Fatalf("Failed to create new GuestCommands for OS: %s", UnixOSType)
	}
	cmd := guestCmd.ChmodExecutable("/usr/local/bin/script.sh")
	if cmd != "chmod +x '/usr/local/bin/script.sh'" {
		t.Fatalf("Unexpected Unix chmod +x cmd: %s", cmd)
	}

	guestCmd, err = NewGuestCommands(WindowsOSType)
	if err != nil {
		t.Fatalf("Failed to create new GuestCommands for OS: %s", WindowsOSType)
	}
	cmd = guestCmd.ChmodExecutable("C:\\Program Files\\SomeApp\\someapp.exe")
	if cmd != "echo 'skipping chmod C:\\Program Files\\SomeApp\\someapp.exe'" {
		t.Fatalf("Unexpected Windows chmod +x cmd: %s", cmd)
	}
}
