# Observed Failure

Bazel cross compiling go to Linux x86_64 from MacOS results in an error from
mattn/go-sqlite3 claiming cgo compilation was not enabled.

Observed message when running `./build_and_run.sh` on MacOS 14.4.1:

    Creating test table in (db.sqlite3) failed: Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub

# Reproduction Steps

1) Make sure docker and bazel are installed
2) Run `./build_and_run.sh` from the root of the repository
3) Observe the final line of output

# Expected Behavior

As when running via `bazel run //repro`, the output is expected to be some
series of `Read text from DB: Hello SQLite!` lines proportional to the number of
times the command has been run since the SQLite db was created.
