package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	driveNameEn string = "Encrypted"
	driveName   string = "Athena-Drive"
	drivePath   string
)

func createBaseDrive() {
	//rclone config create drive-test drive env_auth true -vv --drive-client-id=619469146244-e97afg89glcs2a1omneeqeigdk6n1l0o.apps.googleusercontent.com --drive-client-secret=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 config_change_team_drive=true config_team_drive=true
	out, err := exec.Command("rclone", "config", "create", "Base-Athena", "drive", "env_auth", "true", "-vv", "--drive-client-id=619469146244-pcj383pubudn3ksf5a622m53jkvem77f.apps.googleusercontent.com", "--drive-client-secret=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", "config_change_team_drive=true", "config_team_drive=true").Output()
	if err != nil {
		fmt.Println("rclone", "config", "create", "Base-Athena", "env_auth", "true", "-vv", "--drive-client-id=619469146244-e97afg89glcs2a1omneeqeigdk6n1l0o.apps.googleusercontent.com", "--drive-client-secret=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", "config_change_team_drive=true", "config_team_drive=true")
		fmt.Println("Unable to create drive config")
		panic(err.Error())
	}
	output := string(out[:])
	fmt.Println(output)
}

func checkConfig() {
	fmt.Println("Checking rclone config")
	getPassword()
	out, err := exec.Command("rclone", "config", "dump").Output()
	if err != nil {
		fmt.Println("Incorrect password")
		os.Exit(1)
	}
	mp := make(map[string]interface{})

	// Decode JSON into our map
	json.Unmarshal([]byte(out[:]), &mp)
	fmt.Println(mp)
	baseConfigured := false
	encConfigured := false
	for k := range mp {
		if strings.HasPrefix(k, "Base-Athena") {
			baseConfigured = true
		}
		if strings.HasPrefix(k, "Athena-Drive") {
			encConfigured = true
		}
	}
	if !baseConfigured {
		createBaseDrive()
	}
	if !encConfigured {
		createEncryptedDrive()
	}
}
