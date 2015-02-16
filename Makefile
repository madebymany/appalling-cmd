all:
	echo "nothing to do"

install:
	cd "$(SOURCE_ROOT)/../" && install -g root -o root -m "0755" "gopath/bin/appalling-cmd" "/usr/local/bin/appalling-cmd"
