wget http://storage.googleapis.com/golang/go1.8.3.linux-386.tar.gz
tar -C /usr/local -xzf go1.8.3.linux-386.tar.gz
#wget http://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz
#tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz

export GOROOT="/usr/local/go"
export GOPATH="$GOROOT:$(pwd):$(pwd)/deps"
export GOBIN="$GOPATH/bin"
export PATH="$PATH:$GOROOT/bin"

go build
# (./test.sh &) 
(./fightcoder-web &)
