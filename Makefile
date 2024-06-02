.PHONY: build clean deploy

build:
	mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -o bin/main src/main.go || (echo "Build failed"; exit 1)
	echo '#!/bin/sh\nexec /var/task/main' > bin/bootstrap
	chmod +x bin/bootstrap
	ls -l bin/main || (echo "main binary not found"; exit 1)
	zip -j bin/function.zip bin/main bin/bootstrap || (echo "Zip creation failed"; exit 1)
	ls -l bin/function.zip || (echo "Zip file not found"; exit 1)

clean:
	rm -rf bin
	rm -f main
	rm -f function.zip

deploy: build
	sam build
	sam deploy --guided
