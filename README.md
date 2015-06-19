# About
Prints an inline gif

# Prerequisites
iterm 2.9.20150512+
see [iterm inline images](http://iterm2.com/images.html#/section/home)

# Install
go get github.com/trotha01/gifme

# Use
gifme little puppies

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
