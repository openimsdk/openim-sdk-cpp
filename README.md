# openim-sdk-cpp (Unity/C#)

This branch exposes a stable **C ABI** on top of [openim-sdk-core](https://github.com/openimsdk/openim-sdk-core) and ships platform-specific native libraries for **Unity / C#** via P/Invoke.

- Primary goals: seamless Unity integration; **broad Go version support** (builds with the latest stable Go); and **Windows 7 compatibility** via a Go 1.20 toolchain when targeting Win7.
- Outputs: dynamic libraries (`.dll/.so/.dylib`) and the corresponding generated C headers (from `-buildmode=c-shared`).
- Go versions: for modern OS targets, use the latest stable Go; for **Windows 7**, compile with **Go 1.20** (CGO enabled) and pin dependencies to Go 1.20–compatible versions.

---

## Targets

- **Windows**: x86 (32-bit) and x64 (64-bit)
- **Android**: `armeabi-v7a`, `arm64-v8a`
- **iOS**: `arm64` (typically static `.a`, or `.dylib` if configured)
- **Linux**: x86_64
- **macOS**: x86_64 / arm64

> Library names follow the usual platform conventions: `openimsdk_*.dll` on Windows, `libopenimsdk.*` on Unix-like platforms.

---

## Prerequisites

- **Go 1.20** with CGO enabled
    - `CGO_ENABLED=1`
   
- Platform toolchains:
    - **Windows**: MinGW-w64 (or MSVC if your scripts/toolchain are set up accordingly)
    - **Android**: Android NDK (`ANDROID_NDK_HOME` set)
    - **iOS/macOS**: Xcode + Command Line Tools
    - **Linux**: GCC or Clang
- Working module cache
    - If you see “missing go.sum entry…”, run `go mod download all` at repo root.

---

## Scripts

All scripts live under `scripts/` and generate the platform artifacts in their respective output folders.

```text
scripts/
  gen_win_dll.bat      # build Windows .dll (x86/x64), generates the .h header as well
  gen_android_so.bat   # build Android .so for armeabi-v7a / arm64-v8a
  gen_ios_dylib.sh     # build iOS static .a (or .dylib if configured)
  gen_linux_so.sh      # build Linux .so (x86_64)
  gen_mac_dylib.sh     # build macOS .dylib
