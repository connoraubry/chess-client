# Chess Client

Client for interacting with chess server.

## Installation

go 1.21 required.

```bash
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
tar -C /usr/local/go -xf go1.21.0.linux-amd64.tar.gz
echo "export PATH="$PATH:/usr/local/go/bin"
```

After installing go, download packages
```bash
go get -u . 
```
And build from source
```bash
make
```


## Running

Either run the binary or build at runtime
```bash
./bin/client

go run main.go
```

## Usage

While in the client, create a game by running `new game` or join an existing game using `join {GAME ID}`.

For additional commands, run `help`.
