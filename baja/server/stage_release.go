package server

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"strings"
	"sync"
)

type StageBuild struct {
	currentCommitID string
	isBuilding      bool
	sync.Mutex
}

type Repo struct {
	uri          string
	branch       string
	lastCommitID string
}

func NewRepo(repo string) (*Repo, error) {
	path := strings.Split(repo, "/")
	if len(path) <= 3 {
		return nil, errors.New("Invalid repo and branch")
	}

	r := Repo{
		uri:    fmt.Sprintf("https://github.com/%s/%s.git", path[0], path[1]),
		branch: path[2],
	}

	return &r, nil
}

var stageBuild = &StageBuild{
	currentCommitID: "",
	isBuilding:      false,
}

// Allow us to view any branch of code from any accessible git repository
func (s *Server) StageRelease(c echo.Context) error {
	stageBuild.Lock()
	if stageBuild.isBuilding == false {
		go checkAndBuild(c.Param("repo"))
	}
	stageBuild.Unlock()

	return c.String(http.StatusOK, "Hello, World!")
}

// sync a git repo to local, refetch if not existing

func syncRepo(repo) {

}

func checkAndBuild(repo string) {
	// F
}
