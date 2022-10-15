build:
	mkdir -p build
	go build -o build/  cmd/gdict.go
	go build -o build/ cmd/install.go
install:
	mkdir -p build
	go build -o build/ cmd/gdict.go
	go build -o build/ cmd/install.go
	mkdir -p ~/.local/bin
	install -m 755 build/gdict ~/.local/bin
	./build/install
uninstall:
	rm ~/.local/bin/gdict
	rm ~/.local/share/gdict/dictionary.db

