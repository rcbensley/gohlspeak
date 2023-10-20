build:
	go clean
	@CGO_ENABLED=1 go build

release:
	goreleaser release
