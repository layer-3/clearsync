tests:
	go test -v -covermode count -coverprofile coverage.out && go tool cover -html coverage.out -o coverage.html && go tool cover -func coverage.out -o coverage.out

tests_race:
	go test -v -race -covermode atomic -coverprofile coverage.out && go tool cover -html coverage.out -o coverage.html && go tool cover -func coverage.out -o coverage.out

tests_with_mock:
	go test -v -race -gcflags=all=-l -covermode atomic -coverprofile coverage.out && go tool cover -html coverage.out -o coverage.html && go tool cover -func coverage.out -o coverage.out

format:
	go fmt .

bench:
	go test -bench . -benchmem -cpu 1

report_bench:
	go test -cpuprofile cpu.prof -memprofile mem.prof -bench . -cpu 1

cpu_report:
	go tool pprof cpu.prof

mem_report:
	go tool pprof mem.prof

build:
	go build -v ./...
