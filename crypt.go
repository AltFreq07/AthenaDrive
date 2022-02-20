package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gen2brain/dlgs"
)

func confEncrypted() bool {
	out, err := exec.Command("rclone", "config", "encrypted").Output()
	if err != nil {
		fmt.Println("Error checking config encryption")
		os.Exit(1)
	}
	res, err := strconv.ParseBool(strings.TrimSuffix(string(out[:]), "\n"))
	if err != nil {
		fmt.Println("Error checking config encryption")
		os.Exit(1)
	}
	return res
}

//create encrypted drive config
func createEncryptedDrive() {
	// .\rclone.exe config create Athena-Drive crypt remote=Athena-Base:Encrypted password=$pass1 password2=$pass2 server-side-across-configs=true
	out, err := exec.Command("rclone", "config", "create", driveName, "crypt", "remote=Base-Athena:Encrypted", "password="+os.Getenv("RCLONE_CONFIG_PASS"), "server-side-across-configs=true", "--obscure").Output()
	if err != nil {
		fmt.Println("rclone", "config", "create", driveName, "crypt", "remote=Base-Athena:Encrypted", "password="+os.Getenv("RCLONE_CONFIG_PASS"), "server-side-across-configs=true")
		fmt.Println("Unable to create drive config")
		panic(err.Error())
	}
	output := string(out[:])
	fmt.Println(output)
}

func getPassword() {
	var passMessage string
	if confEncrypted() {
		passMessage = "Enter a password:"
	} else {
		passMessage = "Create a password:"
	}
	pass, accept, err := dlgs.Password("Encryption", passMessage)

	if err != nil || !accept {
		fmt.Println("Error with password input")
		os.Exit(1)
	}
	os.Setenv("RCLONE_CONFIG_PASS", pass)
	encryptConfig()
}

func encryptConfig() {
	if !confEncrypted() {
		out, err := exec.Command("rclone", "config", "encrypt", os.Getenv("RCLONE_CONFIG_PASS")).Output()
		if err != nil {
			fmt.Println("Error encrypting config")
			os.Exit(1)
		}
		output := string(out[:])
		fmt.Println(output)
	}
}
