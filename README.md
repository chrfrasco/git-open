# git-open

Open the current repo in your browser.

## Installation

```sh
$ go get github.com/chrfrasco/git-open
```

## Usage

Assuming `$GOPATH/bin` is on your path:

```sh
$ git-open # opens default by origin
```

## Create a git alias

Open `~/.gitconfig` and add this line:

```
[alias]
  ...
  open = !sh git-open
```

This adds the `open` command to git:

```sh
$ git open
```

