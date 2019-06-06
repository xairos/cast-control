APP_NAME=cast-control

build:
	go build -v -o $(APP_NAME)

build-all:
	for GOOS in darwin linux windows; do \
    	for GOARCH in 386 amd64; do \
        	go build -v -o $(APP_NAME)-$$GOOS-$$GOARCH; \
    	done \
	done

clean:
	find . -type f  -name "$(APP_NAME)*" ! -name "*.*"  -delete; \
	go clean
