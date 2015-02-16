cd "$(SOURCE_ROOT)"

GOPATH="$(SOURCE_ROOT)/gopath"

all:
	echo "nothing to do"

install:
	install -g root -o root -m "0755" "$(GOPATH)/bin/appalling" "/usr/local/bin/appalling"
