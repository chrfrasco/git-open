package gitremote

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
)

var configPath = "./.git/config"

// ErrNotGitRepo when current directory is not a git repo
var ErrNotGitRepo = errors.New("not a git repository")

var (
	remoteSectionRE = regexp.MustCompile(`^\[remote "(.+)"]`)
	remoteURLRE     = regexp.MustCompile(`^\turl = (.+)$`)
	remoteGitURLRE  = regexp.MustCompile(`^git@github.com:(.+)/(.+)\.git`)
)

// Remote contains the remote name & url
type Remote struct {
	Name, URL string
}

// HTTP returns the HTTP version of the remote URL
func (r Remote) HTTP() string {
	match := remoteGitURLRE.FindStringSubmatch(r.URL)
	if match != nil {
		user := match[1]
		repo := match[2]
		return fmt.Sprintf("https://github.com/%s/%s.git", user, repo)
	}

	return r.URL
}

// Parse returns the result of par
func Parse() ([]Remote, error) {
	file, err := os.Open(configPath)
	switch {
	case os.IsNotExist(err):
		return nil, ErrNotGitRepo
	case err != nil:
		return nil, err
	}
	defer file.Close()

	return parse(file)
}

func parse(file *os.File) ([]Remote, error) {
	remotes := *new([]Remote)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		f := remoteSectionRE.FindStringSubmatch(line)
		if f != nil {
			remote := Remote{}
			for scanner.Scan() {
				line := scanner.Text()
				f := remoteURLRE.FindStringSubmatch(line)
				if f != nil {
					remote.URL = f[1]
				}
			}
			remote.Name = f[1]
			remotes = append(remotes, remote)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return remotes, nil
}
