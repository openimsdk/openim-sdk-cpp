//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var (
	soName  = "libopenimsdk" //
	outPath = "../shared/"
	goSrc   = "go" //
)

var Default = Build

// BuildAll compiles the project for all platforms.
func Build() {
	if err := BuildAndroid(); err != nil {
		fmt.Println("Error building for Android:", err)
	}
	if err := BuildIOS(); err != nil {
		fmt.Println("Error building for iOS:", err)
	}
	if err := BuildLinux(); err != nil {
		fmt.Println("Error building for Linux:", err)
	}
	if err := BuildWindows(); err != nil {
		fmt.Println("Error building for Windows:", err)
	}
}

// BuildAndroid compiles the project for Android.
func BuildAndroid() error {
	architectures := []struct {
		Arch, API string
	}{
		{"arm", "16"},
		{"arm64", "21"},
		{"386", "16"},
		{"amd64", "21"},
	}

	for _, arch := range architectures {
		if err := buildAndroid(outPath+"android", arch.Arch, arch.API); err != nil {
			fmt.Printf("Failed to build for %s: %v\n", arch.Arch, err)
		}
	}
	return nil
}

// BuildIOS compiles the project for iOS.
func BuildIOS() error {
	fmt.Println("Building for iOS...")
	outPath += "ios"
	arch := os.Getenv("GOARCH")

	if len(arch) == 0 {
		arch = runtime.GOARCH
	}

	os.Setenv("GOOS", "darwin")
	os.Setenv("GOARCH", arch)
	os.Setenv("CGO_ENABLED", "1")
	os.Setenv("CC", "clang")

	cmd := exec.Command("go", "build", "-buildmode=c-shared", "-o", outPath+"/"+soName+".dylib", ".")
	cmd.Dir = goSrc
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to build for iOS: %v\n", err)
		return err
	}
	fmt.Println("Build for iOS completed successfully.")
	return nil
}

// BuildLinux compiles the project for Linux.
func BuildLinux() error {
	fmt.Println("Building for Linux...")

	outPath += "linux"
	arch := os.Getenv("GOARCH")
	cc := os.Getenv("CC")
	cxx := os.Getenv("CXX")

	if len(arch) == 0 {
		arch = runtime.GOARCH
	}

	if len(cc) == 0 {
		cc = "gcc"
	}

	if len(cxx) != 0 {
		os.Setenv("CXX", cxx)
	}

	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", arch)
	os.Setenv("CGO_ENABLED", "1")
	os.Setenv("CC", cc) //

	cmd := exec.Command("go", "build", "-buildmode=c-shared", "-trimpath", "-ldflags=-s -w", "-o", outPath+"/"+soName+".so", ".")
	cmd.Dir = goSrc
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to build for Linux: %v\n", err)
		return err
	}
	fmt.Println("Build for Linux completed successfully.")
	return nil
}

// BuildWindows compiles the project for Windows.
func BuildWindows() error {
	fmt.Println("Building for Windows...")

	outPath += "windows"

	os.Setenv("GOOS", "windows")
	os.Setenv("GOARCH", "amd64")
	os.Setenv("CGO_ENABLED", "1")
	os.Setenv("CC", "gcc")

	cmd := exec.Command("go", "build", "-buildmode=c-shared", "-trimpath", "-ldflags=-s -w", "-o", outPath+"/"+soName+".dll", ".")
	cmd.Dir = goSrc
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to build for Windows: %v\n", err)
		return err
	}
	fmt.Println("Build for Windows completed successfully.")
	return nil
}

// buildAndroid builds the Android library for the specified architecture.
func buildAndroid(aOutPath, arch, apiLevel string) error {
	fmt.Printf("Building for %s...\n", arch)

	ndkPath := os.Getenv("ANDROID_NDK_HOME")
	osSuffix := ""
	if runtime.GOOS == "windows" {
		osSuffix = ".cmd" //
	}

	ccBasePath := ndkPath + "/toolchains/llvm/prebuilt/" + runtime.GOOS + "-x86_64/bin/"

	var cc string
	switch arch {
	case "arm":
		cc = ccBasePath + "armv7a-linux-androideabi" + apiLevel + "-clang" + osSuffix
	case "arm64":
		cc = ccBasePath + "aarch64-linux-android" + apiLevel + "-clang" + osSuffix
	case "386":
		cc = ccBasePath + "i686-linux-android" + apiLevel + "-clang" + osSuffix
	case "amd64":
		cc = ccBasePath + "x86_64-linux-android" + apiLevel + "-clang" + osSuffix
	}

	env := []string{
		"CGO_ENABLED=1",
		"GOOS=android",
		"GOARCH=" + arch,
		"CC=" + cc,
	}
	cmd := exec.Command("go", "build", "-buildmode=c-shared", "-trimpath", "-ldflags=-s -w", "-o", aOutPath+"/"+arch+"/"+soName+".so", ".")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Dir = goSrc
	cmd.Env = append(os.Environ(), env...)
	return cmd.Run()
}
