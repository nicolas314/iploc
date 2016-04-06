

iploc: iploc.go
	go build -ldflags "-s" iploc.go
	strip iploc
