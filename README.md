# rules_go test

This is an example project that uses rules_go and uses hermetic_cc_toolchain for cross-compilation. This is currently beiing used to test a Go 1.21 related bug.

`cgoexample_asm`: Demonstrates a bug in `rules_go`, where it uses Go's assembler instead of the C compiler for CGO_ENABLED=1 targets. I think it is this bug: https://github.com/bazelbuild/rules_go/issues/3411


## Works

```
bzl run \
    --incompatible_enable_cc_toolchain_resolution \
    --platforms @zig_sdk//platform:darwin_amd64 \
    --extra_toolchains @zig_sdk//toolchain:darwin_amd64 \
    //hello

bzl run \
    --incompatible_enable_cc_toolchain_resolution \
    --platforms @zig_sdk//platform:darwin_arm64 \
    --extra_toolchains @zig_sdk//toolchain:darwin_arm64 \
    //hello

bzl run \
    --incompatible_enable_cc_toolchain_resolution \
    --platforms @zig_sdk//platform:linux_amd64 \
    --extra_toolchains @zig_sdk//toolchain:linux_amd64_musl \
    //hello

bzl run \
    --incompatible_enable_cc_toolchain_resolution \
    --platforms @zig_sdk//platform:linux_arm64 \
    --extra_toolchains @zig_sdk//toolchain:linux_arm64_musl \
    //hello
```

## Does not work

bzl run --platforms @zig_sdk//platform:linux_amd64 --extra_toolchains @zig_sdk//toolchain:linux_amd64_musl //hello

```
ERROR: /private/var/tmp/_bazel_evan.jones/df0cc3210f1b4eedcd01bbea21b1eabe/external/io_bazel_rules_go/BUILD.bazel:42:7: GoStdlib external/io_bazel_rules_go/stdlib_/pkg failed: (Exit 1): builder failed: error executing command (from target @io_bazel_rules_go//:stdlib) bazel-out/darwin_arm64-opt-exec-2B5CBBC6-ST-b33d65c724e6/bin/external/go_sdk/builder_reset/builder stdlib -sdk external/go_sdk -installsuffix linux_amd64 -out ... (remaining 7 arguments skipped)

Use --sandbox_debug to see verbose messages from the sandbox and retain the sandbox build root for debugging
# runtime/cgo
linux_syscall.c:67:13: error: call to undeclared function 'setresgid'; ISO C99 and later do not support implicit function declarations [-Wimplicit-function-declaration]
linux_syscall.c:67:13: note: did you mean 'setregid'?
/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/usr/include/unistd.h:593:6: note: 'setregid' declared here
linux_syscall.c:73:13: error: call to undeclared function 'setresuid'; ISO C99 and later do not support implicit function declarations [-Wimplicit-function-declaration]
linux_syscall.c:73:13: note: did you mean 'setreuid'?
/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/usr/include/unistd.h:595:6: note: 'setreuid' declared here
stdlib: error running subcommand external/go_sdk/bin/go: exit status 1
Target //hello:hello failed to build
Use --verbose_failures to see the command lines of failed build steps.
```