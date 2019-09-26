cmd := goradamsa

version := 0.1.0-alpha
commit := $(shell git rev-parse --short HEAD 2>/dev/null)
release := $(if $(commit),$(version)+$(commit),$(version))
goflags := -ldflags "-s -w -X main.release=$(release)"

code := main.go

.PHONY: all clean

all: $(code)
	go build -o $(cmd) $(goflags)

clean:
	rm -f $(cmd)
