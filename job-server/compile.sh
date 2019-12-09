
export GOPATH=$GOPATH:/Users/sebastianko/Documents/Program\ Workspace/job\ selector/job-server/
go build -ldflags "-X main.serverVersion=`date -u +.%Y%m%d.%H%M%S`" -o ./linux_startmps ./src
