package git

import (
	"os"
	"os/exec"
)

type Repo struct {
	Dir    string
	Remote string
}

func (r *Repo) cmd(name string, arg ...string) *exec.Cmd {
	cmd := exec.Command(name, arg...)
	cmd.Dir = r.Dir

	return cmd
}

func (r *Repo) Clone() error {
	os.MkdirAll(r.Dir, os.ModePerm)

	cmd := r.cmd("git", "clone", r.Remote, "-o", ".")
	return cmd.Run()
}

func (r *Repo) Checkout(branch string) {

}

func (r *Repo) Sync() {
	cmd := r.cmd("git", "pull")
	cmd.Run()
}
