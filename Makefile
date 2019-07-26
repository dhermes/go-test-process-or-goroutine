.PHONY: help
help:
	@echo 'Makefile for `go-test-process-or-goroutine` project'
	@echo ''
	@echo 'Usage:'
	@echo '   make test    Run all Go tests'
	@echo ''

.PHONY: test
test:
	go test -count=1 ./...
