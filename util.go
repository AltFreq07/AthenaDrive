package main

import (
	"fmt"
	"os/exec"
	"strings"
)

//CO Pilot
func checkRClone() {
	fmt.Println("Checking rclone")
	out, err := exec.Command("rclone", "version").Output()
	if err != nil {
		fmt.Println("Unable to find rclone")
		panic(err.Error())
	}
	fmt.Println(string(out[:]))
}

func mountDrive() {
	fmt.Println("Mounting Drive", driveName+":", drivePath+":")
	// rclone mount Athena-Drive: G: --allow-other --dir-cache-time 5000h --poll-interval 10s --drive-pacer-min-sleep 10ms --drive-pacer-burst 200 --vfs-cache-mode full --vfs-cache-max-size 100G --vfs-cache-max-age 5000h --vfs-cache-poll-interval 5m --vfs-read-ahead 2G --bwlimit-file 40M
	out, err := exec.Command("rclone", "mount", driveName+":", drivePath+":", "--allow-other", "--dir-cache-time", "5000h", "--poll-interval", "10s", "--drive-pacer-min-sleep", "10ms", "--drive-pacer-burst", "200", "--vfs-cache-mode", "full", "--vfs-cache-max-size", "100G", "--vfs-cache-max-age", "5000h", "--vfs-cache-poll-interval", "5m", "--vfs-read-ahead", "2G", "--bwlimit-file", "40M", "--crypt-server-side-across-configs", "--rc", "-vv").Output()
	if err != nil {
		fmt.Println("Unable to mount drive")
		panic(err.Error())
	}
	fmt.Println(string(out[:]))
}

func getMountLocation() {
	drivePath = getNextWindowsDrive()
}

func getNextWindowsDrive() string {
	fmt.Println("Getting next windows drive")
	// [char[]](67..90) | where {(get-wmiobject win32_logicaldisk | select -expand DeviceID) -notcontains "$($_):"} | Select -first 1
	// out, err := exec.Command("powershell", "-WindowStyle", "Hidden", "-C", "[char[]](67..90)", "|", "where", "{(get-wmiobject win32_logicaldisk | select -expand DeviceID) -notcontains \"$($_):\"}", "|", "Select -first 1").Output()
	out, err := exec.Command("fsutil", "fsinfo", "drives").Output()
	if err != nil {
		fmt.Println("Unable to get next windows drive")
		panic(err.Error())
	}
	drives := strings.Fields(string(out))
	return string(rune(drives[len(drives)-1][0] + 1))
	// output := string(out[:])
	// return string(output[0])
}

func stopRClone() ([]byte, error) {
	return exec.Command("rclone", "rc", "core/quit").Output()
}
