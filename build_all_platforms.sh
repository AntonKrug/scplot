echo "Setting Bash strict mode"
# http://redsymbol.net/articles/unofficial-bash-strict-mode/
set -e
set -u
set -o pipefail
echo

echo "Go version and enviroment"
#go version
#go env
echo

echo "Supported distributions:"
#go tool dist list
echo

echo "Installing depedencies"
# go get github.com/logrusorgru/aurora
# go get github.com/dustin/go-humanize
# go get gonum.org/v1/plot
# go get gonum.org/v1/plot/plotter
# go get gonum.org/v1/plot/plotutil
# go get gonum.org/v1/plot/vg
#go get github.com/google/gxui
#go get github.com/google/gxui/drivers/gl
#go get github.com/google/gxui/themes/dark


TS="${TS:-`date +"%Y%m%d-%H%M%S"`}" 
echo "Timestamp is: $TS"

echo "Creating clean build folder"
mkdir -p release
rm -r ./release
mkdir -p release

echo "Building Windows"
GOOS=windows GOARCH=amd64 go build -o release/scplot-windows-x86-64.exe
#GOOS=windows GOARCH=amd64 go build -o release/scplot-windows-x86-64-$TS.exe
# GOOS=windows GOARCH=386 go build -o release/scplot-windows-x86-32-$TS.exe
# GOOS=windows GOARCH=arm go build -o release/scplot-windows-arm-$TS.exe

# echo "Building Linux"
# GOOS=linux GOARCH=amd64 go build -o release/scplot-linux-x86-64-$TS
# GOOS=linux GOARCH=386 go build -o release/scplot-linux-x86-32-$TS
# GOOS=linux GOARCH=arm64 go build -o release/scplot-linux-arm-64-$TS
# GOOS=linux GOARCH=arm go build -o release/scplot-linux-arm-32-$TS

# echo "Building FreeBSD"
# GOOS=freebsd GOARCH=amd64 go build -o release/scplot-freebsd-x86-64-$TS
# GOOS=freebsd GOARCH=386 go build -o release/scplot-freebsd-x86-32-$TS

# echo "Building MacOS"
# GOOS=darwin GOARCH=amd64 go build -o release/scplot-macos-x86-64-$TS
# GOOS=darwin GOARCH=386 go build -o release/scplot-macos-x86-32-$TS

echo
echo "Finished artifacts:"
ls -la ./release
