build :
	go build -o redis-like main.go
run :
	go run main.go
clean :
	go clean -n
	rm -rf redis_like
fmt :
	go fmt
