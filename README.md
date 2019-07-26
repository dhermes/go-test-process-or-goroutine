# go-test-process-or-goroutine

> Does `go test` run package tests in subprocesses or goroutines?

This repo has 3 packages that all compete over a single value in one of the
packages (`winkin.sharedValue`); by running the tests all at once with
`go test`, we can see if there will be contention for this value (across
goroutines) or if `go test` puts test runs in separate processes:

```
$ make test
go test -v -count=1 ./...
=== RUN   TestSetSpecificValue
blinkin_test: pid=8589, ppid=8551--- PASS: TestSetSpecificValue (0.00s)
PASS
ok      github.com/dhermes/go-test-process-or-goroutine/blinkin 0.004s
=== RUN   TestSetSpecificValue
 winkin_test: pid=8590, ppid=8551--- PASS: TestSetSpecificValue (0.00s)
PASS
ok      github.com/dhermes/go-test-process-or-goroutine/nod     0.004s
=== RUN   TestSetValue_GetValue
    nod_test: pid=8588, ppid=8551--- PASS: TestSetValue_GetValue (0.00s)
PASS
ok      github.com/dhermes/go-test-process-or-goroutine/winkin  0.005s
```
