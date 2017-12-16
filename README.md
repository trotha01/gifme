# About
Prints an inline gif

# Prerequisites
iterm 2.9.20150512+
see [iterm inline images](http://iterm2.com/images.html#/section/home)

# Install
go get github.com/trotha01/gifme

# Use
```gifme little puppies```

**Help**
``` gifme -h ```
```
gifsearch is a way to find gifs

Usage:
  gifsearch [flags]

  Flags:
    -c, --count int       number of gifs to return (default 1)
    -e, --engine string   gif engine to use 'giphy' or 'tenor'. If not specified, Tenor is searched first and Gifme if there is an error from Tenor
```

# Example
![](/how_to_gifme.gif)

# Detailed Install
## Install the latest iterm
- go to https://www.iterm2.com/downloads.html
- scroll to bottom, and download the latest nightly build.

## Install golang
- ```brew install go``` (or go to https://golang.org/doc/install)
- ```mkdir -p ~/go/src```
- Add to .bash_profile or .bashrc (osx or linux respectively):
```
export GOPATH=~/go
export PATH="$PATH:$GOPATH/bin"
export GOROOT="/usr/local/go"
```

## Install gifme
``` go get github.com/trotha01/gifme ```
