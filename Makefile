GOPATH="${PWD}/gopath"

all:
	echo "nothing to do"

install:
	install -g root -o root -m "0755" "$(GOPATH)/bin/appalling" "/usr/local/bin/appalling"
