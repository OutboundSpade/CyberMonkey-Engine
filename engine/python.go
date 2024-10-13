package engine

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"runtime"
)

type python struct{}

func (p *python) Install() error {
	switch runtime.GOOS {
	case "linux":
		// Best effort to install Python, using the package manager
		// This is a best effort, as the package manager may not have Python
		// or the package manager may not be installed

		_, err := exec.LookPath("python3")
		if err == nil { // Python is already installed
			return nil
		}

		// check if apt is installed
		_, err = exec.LookPath("apt")
		if err == nil {
			// update the package list
			cmd := exec.Command("apt", "update")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				return err
			}
			cmd = exec.Command("apt", "install", "-y", "python3")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			return cmd.Run()
		}
		_, err = exec.LookPath("dnf")
		if err == nil {
			// update the package list
			cmd := exec.Command("dnf", "update")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				return err
			}
			cmd = exec.Command("dnf", "install", "-y", "python3")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			return cmd.Run()
		}
		_, err = exec.LookPath("yum")
		if err == nil {
			// update the package list
			cmd := exec.Command("yum", "update")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				return err
			}
			cmd = exec.Command("yum", "install", "-y", "python3")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			return cmd.Run()
		}
		_, err = exec.LookPath("pacman")
		if err == nil {
			cmd := exec.Command("pacman", "-Syu", "--noconfirm", "python")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			return cmd.Run()
		}

	case "windows":
		const PYTHON_VERSION = "3.13.0"
		// Download the installer
		const installer = "https://www.python.org/ftp/python/3.13.0/python-" + PYTHON_VERSION + "-amd64.exe"
		res, err := http.Get(installer)
		defer res.Body.Close()

		outfile, err := os.Open(path.Join(INSTALL_DIR_WINDOWS, "python-"+PYTHON_VERSION+"-amd64.exe"))
		if err != nil {
			return err
		}
		defer outfile.Close()

		_, err = io.Copy(outfile, res.Body)
		// Run the installer
		cmd := exec.Command(path.Join(INSTALL_DIR_WINDOWS, "python-"+PYTHON_VERSION+"-amd64.exe"), "/quiet", "InstallAllUsers=0", "PrependPath=1", "Include_test=0")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
	return nil
}

func (p *python) Uninstall() error {
	return nil
}

func (p *python) Info() {
	// Print information about the dependency
}
