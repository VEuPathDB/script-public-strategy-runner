VERSION=$(shell git describe --tags)

build:
	env CGO_ENABLED=0 GOOS=linux go build -o bin/public-strategies -ldflags "-X main.version=${VERSION}" cmd/main.go

travis:
	env CGO_ENABLED=0 GOOS=linux go build -o bin/public-strategies -ldflags "-X main.version=${VERSION}" cmd/main.go
	cd bin && tar -czf public-strategies-linux.${TRAVIS_TAG}.tar.gz ./public-strategies && rm public-strategies

	env CGO_ENABLED=0 GOOS=darwin go build -o bin/public-strategies -ldflags "-X main.version=${VERSION}" cmd/main.go
	cd bin && tar -czf public-strategies-darwin.${TRAVIS_TAG}.tar.gz ./public-strategies && rm public-strategies

	env CGO_ENABLED=0 GOOS=windows go build -o bin/public-strategies.exe -ldflags "-X main.version=${VERSION}" cmd/main.go
	cd bin && zip -9 public-strategies-windows.${TRAVIS_TAG}.zip ./public-strategies.exe && rm public-strategies.exe
