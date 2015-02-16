all:
	echo "nothing to do"

install:
	cd "$(SOURCE_ROOT)../" && install -g root -o root -m "0755" "gopath/bin/appalling" "/usr/local/bin/appalling"
