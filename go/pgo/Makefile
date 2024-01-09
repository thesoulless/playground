default: build.pgo

.PHONY: run
run:
	@go run ./

.PHONY: profile
profile:
	@curl -o cpu.pprof "http://localhost:9091/debug/pprof/profile?seconds=30"

.PHONY: build.pgo
build.pgo: clean.pgo
	@go build -o bin/pgo
	@go version -m bin/pgo

.PHONY: build
build: clean.build
	@mv default.pgo x.pgo
	@go build -o bin/npgo
	@go version -m bin/npgo
	@mv x.pgo default.pgo

.PHONY: clean.pgo
clean.pgo:
	@rm -rf bin/pgo

.PHONY: clean.build
clean.build:
	@rm -rf bin/npgo

.PHONY: diff
diff:
	@go install golang.org/x/perf/cmd/benchstat@latest
	benchstat nopgo.txt withpgo.txt
