.PHONY: test test_race test_with_mock test_ci_coverage format bench report_bench cpu_report mem_report build

COVER_OUT := coverage.out
COVER_HTML := coverage.html

test: COVER_OPTS = -covermode count
test_race: COVER_OPTS = -race -covermode atomic
test_with_mock: COVER_OPTS = -race -gcflags=all=-l -covermode atomic

test test_race test_with_mock:
	go test -v $(COVER_OPTS) -coverprofile=$(COVER_OUT) && go tool cover -html=$(COVER_OUT) -o $(COVER_HTML) && go tool cover -func=$(COVER_OUT) -o $(COVER_OUT)

test_ci_coverage:
	go test -race -gcflags=all=-l -coverprofile=coverage.txt -covermode=atomic

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
