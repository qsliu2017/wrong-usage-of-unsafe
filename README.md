# Wrong Usage Of Unsafe

Rule of CodeQL [WrongUsageOfUnsafe](go/ql/src/experimental/Unsafe/UnsafeUsageOK.go)

- [slice_test.go](slice_test.go) tests slice/array address and conversion
- [columns.go](columns.go) has 2 bad / 2 good usages of unsafe
- [columns_test.go](columns_test.go) proves the correctness of 4 usages
- [columns_bench_test.go](columns_bench_test.go) compares the performance of 4 usages
