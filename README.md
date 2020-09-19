# axhsare-go

Server API of `axshare` for hosting and sharing Axure RP prototypes

Web Project: https://github.com/XanderCheung/axshare

## Build Step
```shell script
# config
cp config/.env.example config/.env

# build
go build -ldflags '-w -s' cmd/axshare_go.go
# build linux app
# env GOOS=linux go build -ldflags '-w -s' cmd/axshare_go.go

# run
./axshare_go
```

## Software to be installed
`wget` - Download files from QiNiu

`unar` - Extract the contents of archive files

## more
```shell script
# kill
pgrep -f ./axshare_go | xargs kill -9;true

# nohup 
nohup ./axshare_go >> log/production.log &
```