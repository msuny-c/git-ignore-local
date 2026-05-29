package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const localIgnoreFile = ".local.gitignore"

func main() {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	out, err := cmd.Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: not inside a git repository")
		os.Exit(1)
	}

	gitRoot := strings.TrimSpace(string(out))

	if err := os.Chdir(gitRoot); err != nil {
		fmt.Fprintf(os.Stderr, "error: cannot change to git root: %v\n", err)
		os.Exit(1)
	}

	if _, err := os.Stat(localIgnoreFile); os.IsNotExist(err) {
		if err := os.WriteFile(localIgnoreFile, []byte(""), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "error: cannot create %s: %v\n", localIgnoreFile, err)
			os.Exit(1)
		}
	}

	excludePath := filepath.Join(".git", "info", "exclude")
	content, err := os.ReadFile(excludePath)
	if err != nil && !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "error: cannot read exclude file: %v\n", err)
		os.Exit(1)
	}

	if !strings.Contains(string(content), localIgnoreFile) {
		f, err := os.OpenFile(excludePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: cannot open exclude file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()

		if len(content) > 0 && !bytes.HasSuffix(content, []byte("\n")) {
			f.WriteString("\n")
		}
		f.WriteString(localIgnoreFile + "\n")
	}

	configCmd := exec.Command("git", "config", "--local", "core.excludesfile", localIgnoreFile)
	if err := configCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: cannot set git config: %v\n", err)
		os.Exit(1)
	}
}
