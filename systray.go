package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"

	// "github.com/gen2brain/dlgs"
	"fyne.io/fyne/v2/widget"
	"github.com/getlantern/systray"
)

func onReady() {
	systray.SetTemplateIcon(Data, Data)
	systray.SetTitle("Athena Drive")
	systray.SetTooltip("Athena Drive Tooltip")
	mChangePass := systray.AddMenuItem("Change Password", "Change your encryption password")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	for {
		select {
		case <-mQuit.ClickedCh:
			systray.Quit()
		case <-mChangePass.ClickedCh:
			w := fyneApp.NewWindow("Hello World")

			w.SetContent(widget.NewLabel("Hello World!"))
			w.ShowAndRun()
			yes := false
			// yes, err := dlgs.Question("Change Password", "Are you sure you want to change your encryption password?\n Warning: This can take some time.", true)
			// if err != nil {
			// 	panic(err)
			// }
			if yes {
				fmt.Println("Changing password")
				//rclone decrypt
				decryptConfig()
				//rclone unmount
				out, err := stopRClone()
				if err != nil {
					fmt.Println("Unable to stop rclone")
					panic(err.Error())
				}
				fmt.Println(string(out))
				//rclone rename
				out, err = renameRClone("Athena-Drive:", "Athena-Old")
				if err != nil {
					fmt.Println("Unable to rename config")
					panic(err.Error())
				}
				fmt.Println(string(out))
				//get new password
				getPassword()
				createEncryptedDrive()
				//rclone move old to AthenaDrive
				copyDrive("Athena-Old:", "Athena-Drive:")
				//delete old encrypted drive
				// out, err = deleteConfig("Athena-Old:")
				// if err != nil {
				// 	fmt.Println("Unable to rename config")
				// 	panic(err.Error())
				// }
				// fmt.Println(string(out))
				//mount
				go mountDrive()
			}
		}
	}
}

func copyDrive(from string, to string) {
	cmd := exec.Command("rclone", "move", "-P", "--delete-empty-src-dirs", "--create-empty-src-dirs", from, to)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	_ = cmd.Start()
	fmt.Println("Copying drive from " + from + " to " + to)

	// print the output of the subprocess
	scanner := bufio.NewScanner(io.MultiReader(stdout, stderr))
	// scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	fmt.Println("Waiting for process to exit")
	_ = cmd.Wait()
	fmt.Print("Done")

}

func deleteConfig(config string) ([]byte, error) {
	fmt.Println("Deleting rclone config")
	// rclone.exe config decrypt
	return exec.Command("rclone", "config", "delete", config).Output()
}

func renameRClone(oldName string, newName string) ([]byte, error) {
	fmt.Println("Renaming rclone config")
	// rclone.exe config decrypt
	return exec.Command("rclone", "config", "rename", oldName, newName).Output()
}

func decryptConfig() {
	fmt.Println("Decrypting rclone config")
	// rclone.exe config decrypt
	out, err := exec.Command("rclone", "config", "decrypt").Output()
	if err != nil {
		fmt.Println("Unable to decrypt config")
		panic(err.Error())
	}
	fmt.Println(string(out))
}

func onExit() {
	// clean up here
	fmt.Println("Closing rclone")
	// rclone rc core/quit
	out, err := stopRClone()
	if err != nil {
		fmt.Println("Unable to quit")
		panic(err.Error())
	}
	exit <- true
	fmt.Println(out)
}
