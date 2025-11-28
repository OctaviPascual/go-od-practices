package downloader

// Code taken from https://github.com/OctaviPascual/twg/blob/master/sub/downloader_test.go
// This is also interesting: https://npf.io/2015/06/testing-exec-command/

import "os/exec"

func NewDownloader() *Downloader {
	return &Downloader{
		CloneCmd: exec.Command("git", "clone"),
	}
}

type Downloader struct {
	CloneCmd *exec.Cmd
}

func (d *Downloader) Download(repo, dst string) (string, error) {
	var cmd *exec.Cmd
	if d.CloneCmd == nil {
		cmd = exec.Command("git", "clone")
	}
	name, args := d.CloneCmd.Args[0], append(d.CloneCmd.Args[1:], repo, dst)
	cmd = exec.Command(name, args...)
	cmd.Env = d.CloneCmd.Env
	bytes, err := cmd.Output()
	if err != nil {
		return string(bytes), err
	}
	return string(bytes), nil
}
