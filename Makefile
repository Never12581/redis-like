build :
	go build -o redis_like main.go
run :
	go run main.go
clean :
	go clean -n
	rm -rf redis-like nohup.out redis_like
fmt :
	go fmt
