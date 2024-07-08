compile:
	GOOS=windows GOARCH=amd64 go build ./cmd/store-by-gin-redis-postgress/
	mv ./store-by-gin-redis-postgress.exe ./bin/store-by-gin-redis-postgress_win_amd64.exe
	GOOS=darwin GOARCH=amd64 go build ./cmd/store-by-gin-redis-postgress/
	mv ./store-by-gin-redis-postgress ./bin/store-by-gin-redis-postgress_darwin_arm64
	GOOS=linux GOARCH=amd64 go build ./cmd/store-by-gin-redis-postgress/
	mv ./store-by-gin-redis-postgress ./bin/store-by-gin-redis-postgress_linux_amd64
	GOOS=darwin GOARCH=amd64 go build ./cmd/store-by-gin-redis-postgress/
	mv ./store-by-gin-redis-postgress ./bin/store-by-gin-redis-postgress_darwin_amd64