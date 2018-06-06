package gitremote

import (
	"strings"
	"testing"
)

var singleRemoteConf = `
[remote "origin"]
	url = git@github.com/chrfrasco/git-open.git
`
var singleRemoteConfOutput = []Remote{
	Remote{"origin", "git@github.com/chrfrasco/git-open.git"},
}

var twoRemoteConf = `
[remote "origin"]
	url = git@github.com/chrfrasco/git-open.git
[remote "heroku"]
	url = https://git.heroku.com/chrfrasco/git-open.git
`
var twoRemoteConfOutput = []Remote{
	Remote{"origin", "git@github.com/chrfrasco/git-open.git"},
	Remote{"heroku", "https://git.heroku.com/chrfrasco/git-open.git"},
}

type testpair struct {
	config  string
	remotes []Remote
}

var tests = []testpair{
	{``, []Remote{}},
	{singleRemoteConf, singleRemoteConfOutput},
	{twoRemoteConf, twoRemoteConfOutput},
}

func parseString(config string) ([]Remote, error) {
	return parse(strings.NewReader(config))
}

func TestParse(t *testing.T) {
	for _, pair := range tests {
		remotes, err := parseString(pair.config)
		if err != nil {
			t.Error(err)
		}
		for i, remote := range remotes {
			if remote != pair.remotes[i] {
				t.Error("For", pair.config,
					"expected", pair.remotes,
					"got", remotes)
			}
		}
	}
}
