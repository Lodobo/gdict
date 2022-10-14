build:
	go build cmd/gdict.go
	go build cmd/install.go
install:
	go build cmd/gdict.go
	go build cmd/install.go
	install -m 755 gdict ~/.local/bin
	./install

