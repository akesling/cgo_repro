# Observed Failure

Bazel cross compiling go to Linux x86_64 from MacOS results in an error that
`libisl` is not present under `/usr/local/`.

Observed message when running `./build_and_run.sh` on MacOS 14.4.1:

    ...
    dyld[57733]: Library not loaded: /usr/local/opt/isl/lib/libisl.23.dylib
    ...

See https://gist.github.com/akesling/2b2d059cd6fc1e5257026e562b7f933a for
detailed error output.

# Reproduction Steps

1) Make sure docker and bazel are installed
2) Run `./build_and_run.sh` from the root of the repository
3) Observe the final line of output

# Expected Behavior

As when running via `bazel run //repro`, the output is expected to be some
series of `Read text from DB: Hello SQLite!` lines proportional to the number of
times the command has been run since the SQLite db was created.
