cd /Users/sebastianko/Documents/Program\ Workspace/job\ selector/job-server/

rm -rf ./target

export GOPATH=$GOPATH:/Users/sebastianko/Documents/Program\ Workspace/job\ selector/job-server/
GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -ldflags "-X main.serverVersion=`date -u +.%Y%m%d.%H%M%S`" -o ./linux_startmps ./src

mkdir ./target
mv ./linux_startmps ./target
cp ./run_mps.sh ./target
