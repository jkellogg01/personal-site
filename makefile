TARGET := server

run: build
	./server/bin/$(TARGET)

build:
	cd client; npm run build;
	go build -C server -o bin/$(TARGET) main.go