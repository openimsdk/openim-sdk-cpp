package main

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

var Default = GenGo

// Aliases is alias for mage, like `mage go` is `mage GenGo`
var Aliases = map[string]any{
	"go":    GenGo,
	"java":  GenJava,
	"js":    GenJS,
	"ts":    GenTS,
	"rs":    GenRust,
	"swift": GenSwift,
	"ap":    AllProtobuf,

	"android": BuildAndroid,
	"ios":     BuildiOS,
	"mac":     BuildmacOS,
	"linux":   BuildLinux,
	"windows": BuildWindows,
	"h12":     BuildHarmonyOS_API12,
	"h9":      BuildHarmonyOS_API9,
	"al":      AllDynamicLib,

	"docs": GenDocs,
}

/* Protocol Generate */

// Langeuage target
// Define output directories for each target language
const (
	GO     = "go"
	JAVA   = "java"
	CSharp = "csharp"
	JS     = "js"
	TS     = "ts"
	RS     = "rust"
	Swift  = "swift"
)

// protoModules lists all the protobuf modules to be processed for code generation.
var protoModules = []string{
	"common",
	"conversation",
	"error",
	"event",
	"ffi",
	"group",
	"init",
	"interop",
	"message",
	"relation",
	"shared",
	"third",
	"user",
}

// proto files directory path
// ignore Separator, just append folder name, like "./pb/proto" is (".", "pb", "proto")

var protoDir = filepath.Join(".", "proto") // "./proto"

/*
protoc --go_out=:./ --go_opt=module=github.com/openimsdk/openim-sdk-core/v3/proto *.proto
protoc --go_out=./${name} --go_opt=module=github.com/openimsdk/openim-sdk-core/v3/proto/go/${name} proto/${name}.proto
*/

/*
TypeScript requires installing `ts-proto` via a package manager
// JavaScript requires installing `protoc-gen-js` via a package manager (Use TypeScript first.)
*/

// Generate code for all languages (Go, Java, C#, JS, TS) from protobuf files.
func AllProtobuf() error {
	if err := GenGo(); err != nil {
		return err
	}
	if err := GenJava(); err != nil {
		return err
	}
	if err := GenCSharp(); err != nil {
		return err
	}
	if err := GenJS(); err != nil {
		return err
	}
	if err := GenTS(); err != nil {
		return err
	}
	return nil
}

// Generate documentation for sdk protobuf interfaces.
func GenDocs() error {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)
	log.Println("Generating documentation from proto files")

	protoc, err := getToolPath("protoc")
	if err != nil {
		return err
	}

	docsOutDir := filepath.Join(protoDir, "docs")

	for _, module := range protoModules {
		if err := os.MkdirAll(filepath.Join(docsOutDir, module), 0755); err != nil {
			return err
		}

		args := []string{
			"--proto_path=" + protoDir,
			// "--doc_out=" + filepath.Join(docsOutDir, module),
			"--doc_out=" + filepath.Join(docsOutDir),
			"--doc_opt=markdown," + strings.Join([]string{module, "md"}, "."),
			filepath.Join("proto", module) + ".proto",
			// filepath.Join("proto", module+".sdkapi") + ".proto",
		}
		log.Println(protoc, args)
		cmd := exec.Command(protoc, args...)
		connectStd(cmd)
		if err := cmd.Run(); err != nil {
			log.Printf("Error generating documentation for module %s: %v\n", module, err)
			continue
		}
	}

	return nil
}

// Generate Go code from protobuf files.
func GenGo() error {
	log.SetOutput(os.Stdout)
	// log.SetFlags(log.Lshortfile)
	log.Println("Generating Go code from proto files")

	goOutDir := filepath.Join(protoDir, GO)

	protoc, err := getToolPath("protoc")
	if err != nil {
		return err
	}

	for _, module := range protoModules {
		if err := os.MkdirAll(filepath.Join(goOutDir, module), 0755); err != nil {
			return err
		}

		args := []string{
			"--proto_path=" + protoDir,
			"--go_out=" + filepath.Join(goOutDir, module),
			//"--go-grpc_out=" + filepath.Join(goOutDir, module),
			"--go_opt=module=github.com/openimsdk/openim-sdk-core/v3/proto/" + strings.Join([]string{GO, module}, "/"),
			//"--go-grpc_opt=module=github.com/openimsdk/openim-sdk-core/v3/proto/" + strings.Join([]string{GO, module}, "/"),
			filepath.Join("proto", module) + ".proto",
		}
		//log.Println(protoc, args)
		cmd := exec.Command(protoc, args...)
		connectStd(cmd)
		if err := cmd.Run(); err != nil {
			log.Printf("Error generating Go code for module %s: %v\n", module, err)
			continue
		}
	}

	if err := removeOmitemptyTags(); err != nil {
		log.Println("Remove Omitempty is Error", err)
		return err
	} else {
		log.Println("Remove Omitempty is Success")
	}

	return nil
}

// Generate Java code from protobuf files.
func GenJava() error {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)
	log.Println("Generating Java code from proto files")

	javaOutDir := filepath.Join(protoDir, JAVA)

	protoc, err := getToolPath("protoc")
	if err != nil {
		return err
	}

	for _, module := range protoModules {
		if err := os.MkdirAll(filepath.Join(javaOutDir, module), 0755); err != nil {
			return err
		}

		args := []string{
			"--proto_path=" + protoDir,
			"--java_out=lite:" + filepath.Join(javaOutDir, module),
			filepath.Join("proto", module) + ".proto",
		}

		cmd := exec.Command(protoc, args...)
		connectStd(cmd)
		if err := cmd.Run(); err != nil {
			log.Printf("Error generating Java code for module %s: %v\n", module, err)
			continue
		}
	}

	return nil
}

// Generate C# code from protobuf files.
func GenCSharp() error {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)
	log.Println("Generating C# code from proto files")

	csharpOutDir := filepath.Join(protoDir, CSharp)

	protoc, err := getToolPath("protoc")
	if err != nil {
		return err
	}

	for _, module := range protoModules {
		if err := os.MkdirAll(filepath.Join(csharpOutDir, module), 0755); err != nil {
			return err
		}

		args := []string{
			"--proto_path=" + protoDir,
			"--csharp_out=" + filepath.Join(csharpOutDir, module),
			filepath.Join("proto", module) + ".proto",
		}

		cmd := exec.Command(protoc, args...)
		connectStd(cmd)

		if err := cmd.Run(); err != nil {
			log.Printf("Error generating C# code for module %s: %v\n", module, err)
			continue
		}
	}

	return nil
}

// Use TypeScript first. JavaScript need check it useful.

// Generate JavaScript code from protobuf files.
func GenJS() error {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)
	log.Println("Generating JavaScript code from proto files")

	jsOutDir := filepath.Join(protoDir, JS)

	protoc, err := getToolPath("protoc")
	if err != nil {
		return err
	}

	for _, module := range protoModules {
		if err := os.MkdirAll(filepath.Join(jsOutDir, module), 0755); err != nil {
			return err
		}

		args := []string{
			"--proto_path=" + protoDir,
			"--js_out=import_style=commonjs,binary:" + filepath.Join(jsOutDir, module),
			filepath.Join("proto", module) + ".proto",
		}

		cmd := exec.Command(protoc, args...)
		connectStd(cmd)
		if err := cmd.Run(); err != nil {
			log.Printf("Error generating JavaScript code for module %s: %v\n", module, err)
			continue
		}
	}

	return nil
}

// Generate Harmony JavaScript code from protobuf files.
// Note: please install pbjs and pbts command first
// Reference Link: https://ohpm.openharmony.cn/#/cn/detail/@ohos%2Fprotobufjs
func GenHarmonyTS() error {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)

	log.Println("Start Gen JS")
	// 生成js
	outJSFile := "proto.js"
	args := make([]string, 0)
	args = append(args, "-t", "static-module")
	args = append(args, "-w", "es6")
	args = append(args, "-o", outJSFile)
	for _, module := range protoModules {
		protoFile := fmt.Sprintf("%s\\%s.proto", "proto", module)
		args = append(args, protoFile)
	}
	jscmd := exec.Command("pbjs", args...)
	jscmd.Env = os.Environ()
	connectStd(jscmd)

	log.Println("Run: " + jscmd.String())
	if err := jscmd.Run(); err != nil {
		log.Panicf("Error generating Harmony JavaScript code : %v\n", err)
	}
	// 生成TS定义
	outTSDefFile := "proto.d.ts"
	tscmd := exec.Command("pbts",
		outJSFile,
		"-o", outTSDefFile,
	)
	tscmd.Env = os.Environ()
	connectStd(tscmd)

	log.Println("Run: " + tscmd.String())
	if err := tscmd.Run(); err != nil {
		log.Panicf("Error generating Harmony TS define code : %v\n", err)
	}

	// 修改生成的文件
	// 1.将生成的js文件中的 import * as $protobuf from "protobufjs/minimal";
	// 修改为   import { index } from "@ohos/protobufjs"; const $protobuf = index;

	// 2.将生成的.d.ts文件中的 import * as $protobuf from "protobufjs";
	// 修改为  import * as $protobuf from "@ohos/protobufjs";

	// 3.在生成的js文件中 import * as $protobuf from "@ohos/protobufjs";这行代码下方添加如下代码
	// import Long from 'long';
	// $protobuf.util.Long=Long
	// $protobuf.configure()
	replaceStr := func(filePath, oldStr, newStr string) {
		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Panic("failed to read file: %w", err)
		}

		originalContent := string(content)
		modifiedContent := strings.Replace(originalContent, oldStr, newStr, 1) // 只替换一次

		if originalContent == modifiedContent {
			return
		}
		err = os.WriteFile(filePath, []byte(modifiedContent), 0644)
		if err != nil {
			log.Panic("failed to write file: %w", err)
		}
	}
	replaceStr(outJSFile, "import * as $protobuf from \"protobufjs/minimal\";", "import { index } from \"@ohos/protobufjs\"; \nconst $protobuf = index; \n import Long from 'long';\n$protobuf.util.Long=Long \n$protobuf.configure()")
	replaceStr(outTSDefFile, "import * as $protobuf from \"protobufjs\";\nimport Long = require(\"long\");", "import * as $protobuf from \"@ohos/protobufjs\"\nimport Long from 'long';")
	return nil
}

// Generate TypeScript code from protobuf files.
func GenTS() error {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)
	log.Println("Generating TypeScript code from proto files")

	tsOutDir := filepath.Join(protoDir, TS)

	protoc, err := getToolPath("protoc")
	if err != nil {
		return err
	}

	tsProto := filepath.Join(".", "node_modules", ".bin", "protoc-gen-ts_proto")

	if runtime.GOOS == "windows" {
		tsProto = filepath.Join(".", "node_modules", ".bin", "protoc-gen-ts_proto.cmd")
	}

	if _, err := os.Stat(tsProto); err != nil {
		log.Println("tsProto Not Found. Error: ", err, " tsProto Path: ", tsProto)
		return err
	}

	for _, module := range protoModules {
		if err := os.MkdirAll(filepath.Join(tsOutDir, module), 0755); err != nil {
			return err
		}

		args := []string{
			"--proto_path=" + protoDir,
			"--plugin=protoc-gen-ts_proto=" + tsProto,
			"--ts_proto_opt=esModuleInterop=true,messages=true,outputJsonMethods=false,outputPartialMethods=false,outputClientImpl=false,outputEncodeMethods=false,useOptionals=messages",
			"--ts_proto_out=" + filepath.Join(tsOutDir, module),
			filepath.Join("proto", module) + ".proto",
		}

		cmd := exec.Command(protoc, args...)
		connectStd(cmd)
		if err := cmd.Run(); err != nil {
			log.Printf("Error generating TypeScript code for module %s: %v\n", module, err)
			continue
		}
	}

	return nil
}

// Generate Rust code from protobuf files.
func GenRust() error {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)
	log.Println("Generating Rust code from proto files")

	rsOutDir := filepath.Join(protoDir, RS)

	protoc, err := getToolPath("protoc")
	if err != nil {
		return err
	}

	rustgRPC, err := getRustToolPath("grpc_rust_plugin")
	if err != nil {
		return err
	}

	for _, module := range protoModules {
		if err := os.MkdirAll(filepath.Join(rsOutDir, module), 0755); err != nil {
			return err
		}

		args := []string{
			"--proto_path=" + protoDir,
			"--rust_out=kernel=upb:" + filepath.Join(rsOutDir, module),
			"--grpc_out=" + filepath.Join(rsOutDir, module),
			"--plugin=protoc-gen-grpc=" + rustgRPC,
			"--rust_opt=experimental-codegen=enabled",
			filepath.Join("proto", module) + ".proto",
		}

		cmd := exec.Command(protoc, args...)
		connectStd(cmd)
		if err := cmd.Run(); err != nil {
			log.Printf("Error generating Rust code for module %s: %v\n", module, err)
			continue
		}
	}

	return nil
}

// Generate Swift code from protobuf files.
func GenSwift() error {
	// Configure logging
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)
	log.Println("Generating Swift code from proto files")

	swiftOutDir := filepath.Join(protoDir, Swift)

	// Find protoc and Swift plugin paths
	protoc, err := getToolPath("protoc")
	if err != nil {
		return err
	}

	// Iterate over proto modules to generate Swift code
	for _, module := range protoModules {
		modulePath := filepath.Join(protoDir, module+".proto")
		outputPath := filepath.Join(swiftOutDir, module)

		// Ensure the output directory for the module exists
		if err := os.MkdirAll(outputPath, 0755); err != nil {
			return err
		}

		// Prepare protoc command
		args := []string{
			"--proto_path=" + protoDir,
			"--swift_out=" + outputPath,
			"--swift_opt=Visibility=" + "Public",
			modulePath,
		}

		cmd := exec.Command(protoc, args...)
		connectStd(cmd) // Connect command's output to standard output for logging

		// Run the command and handle errors
		if err := cmd.Run(); err != nil {
			log.Printf("Error generating Swift code for module %s: %v\n", module, err)
			continue
		}
		log.Printf("Successfully generated Swift code for module %s\n", module)
	}

	return nil
}

/* Tools */

var bindlingDir = filepath.Join(".", "bindings")

// Generate WebAssembly (wasm) file from go code(./bindings/wasm)
func Wasm() error {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)
	log.Println("Generating WebAssembly file")

	wasmDir := filepath.Join(bindlingDir, "wasm")

	cmd := exec.Command("go", "build", "-trimpath", "-ldflags", "-s -w", "-o", filepath.Join(wasmDir, "output", "openIM.wasm"), filepath.Join(wasmDir, "main.go"))
	cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")

	connectStd(cmd)
	if err := cmd.Run(); err != nil {
		log.Printf("Error generating WebAssembly file: %v\n", err)
	}

	return nil
}

// ffi_c

var soName = "libopenimsdk"

var outPath = filepath.Join(".", "output")
var goSrc = filepath.Join(".", bindlingDir, "ffi_c")

// Builds the project for all platforms and generates Go dynamic libraries
// (e.g., .so, .dylib, .dll) for each platform (Android, iOS, Linux, Windows).
func AllDynamicLib() {
	if err := BuildAndroid(); err != nil {
		fmt.Println("Error building for Android:", err)
	}
	if err := BuildiOS(); err != nil {
		fmt.Println("Error building for iOS:", err)
	}
	if err := BuildLinux(); err != nil {
		fmt.Println("Error building for Linux:", err)
	}
	if err := BuildWindows(); err != nil {
		fmt.Println("Error building for Windows:", err)
	}
}

// Compiles the project for Android and generates a Go dynamic library for Android.
func BuildAndroid() error {
	architectures := []struct {
		GoArch, API, ArchName string
	}{
		{"arm", "16", "armeabi-v7a"},
		{"arm64", "21", "arm64-v8a"},
		{"386", "16", "x86"},
		{"amd64", "21", "x86_64"},
	}

	androidOut := filepath.Join(outPath, "android")

	for _, arch := range architectures {
		if err := os.MkdirAll(filepath.Join(goSrc, androidOut, arch.ArchName), 0755); err != nil {
			return err
		}

		if err := buildAndroid(androidOut, arch.GoArch, arch.API, arch.ArchName); err != nil {
			fmt.Printf("Failed to build for Android %s: %v\n", arch.ArchName, err)
		}
	}
	return nil
}

// Compiles the project for iOS and generates a Go static library for iOS.
func BuildiOS() error {
	log.SetOutput(os.Stdout)
	// log.SetFlags(log.Lshortfile)
	log.Println("Building for iOS...")

	os.Setenv("GOOS", "ios")
	os.Setenv("GOARCH", "arm64")
	os.Setenv("CGO_ENABLED", "1")
	os.Setenv("CC", "clang")

	sdkPathCmd := exec.Command("xcrun", "--sdk", "iphoneos", "--show-sdk-path")
	sdkPathOutput, err := sdkPathCmd.Output()
	if err != nil {
		return fmt.Errorf("failed to get SDK path: %v", err)
	}

	sdkPath := strings.TrimSpace(string(sdkPathOutput))
	cFlags := fmt.Sprintf("-arch arm64 -miphoneos-version-min=9.0 -isysroot %s", sdkPath)
	ldFlags := fmt.Sprintf("-arch arm64 -miphoneos-version-min=9.0 -isysroot %s", sdkPath)
	os.Setenv("CGO_CFLAGS", cFlags)
	os.Setenv("CGO_LDFLAGS", ldFlags)

	iosOut := filepath.Join(outPath, "ios")

	if err := os.MkdirAll(filepath.Join(goSrc, iosOut), 0755); err != nil {
		return err
	}

	log.Println(filepath.Join(goSrc, iosOut))

	buildTags := "darwin,!macos"

	cmd := exec.Command(
		"go",
		"build",
		"-buildmode=c-archive",
		"-trimpath",
		"-ldflags=-s -w",
		"-tags", buildTags,
		"-o", filepath.Join(iosOut, strings.Join([]string{soName, "a"}, ".")), ".")
	cmd.Dir = goSrc
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return err
	}
	log.Println("Build for iOS completed successfully.")
	return nil
}

// Compiles the project for macOS and generates a Go dynamic library for macOS.
func BuildmacOS() error {
	log.SetOutput(os.Stdout)
	// log.SetFlags(log.Lshortfile)
	log.Println("Building for iOS...")

	arch := os.Getenv("GOARCH")

	if len(arch) == 0 {
		arch = runtime.GOARCH
	}

	os.Setenv("GOOS", "darwin")
	os.Setenv("GOARCH", arch)
	os.Setenv("CGO_ENABLED", "1")
	os.Setenv("CC", "clang")

	iosOut := filepath.Join(outPath, "macOS")

	if err := os.MkdirAll(filepath.Join(goSrc, iosOut), 0755); err != nil {
		return err
	}

	log.Println(filepath.Join(goSrc, iosOut))

	cmd := exec.Command("go", "build", "-buildmode=c-shared", "-o", filepath.Join(iosOut, strings.Join([]string{soName, "dylib"}, ".")), ".")
	cmd.Dir = goSrc
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return err
	}
	log.Println("Build for iOS completed successfully.")
	return nil
}

// Compiles the project for Linux and generates a Go dynamic library for Linux.
func BuildLinux() error {
	log.SetOutput(os.Stdout)
	// log.SetFlags(log.Lshortfile)
	log.Println("Building for Linux...")

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

	linuxOut := filepath.Join(outPath, "linux")

	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", arch)
	os.Setenv("CGO_ENABLED", "1")
	os.Setenv("CC", cc) //

	if err := os.MkdirAll(filepath.Join(goSrc, linuxOut), 0755); err != nil {
		return err
	}

	cmd := exec.Command("go", "build", "-buildmode=c-shared", "-trimpath", "-ldflags=-s -w", "-o", filepath.Join(linuxOut, strings.Join([]string{soName, "so"}, ".")), ".")
	cmd.Dir = goSrc
	cmd.Env = os.Environ()

	connectStd(cmd)
	if err := cmd.Run(); err != nil {
		log.Printf("Failed to build for Linux: %v\n", err)
		return err
	}
	log.Println("Build for Linux completed successfully.")
	return nil
}

// Compiles the project for Windows and generates a Go dynamic library for Windows.
func BuildWindows() error {
	log.SetOutput(os.Stdout)
	// log.SetFlags(log.Lshortfile)
	log.Println("Building for Windows...")

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

	windowsOut := filepath.Join(outPath, "windows")

	os.Setenv("GOOS", "windows")
	os.Setenv("GOARCH", arch)
	os.Setenv("CGO_ENABLED", "1")
	os.Setenv("CC", cc)

	if err := os.MkdirAll(filepath.Join(goSrc, windowsOut), 0755); err != nil {
		return err
	}
	cmd := exec.Command("go", "build", "-buildmode=c-shared", "-trimpath", "-ldflags=-s -w", "-o", filepath.Join(windowsOut, strings.Join([]string{soName, "dll"}, ".")), ".")
	cmd.Dir = goSrc
	cmd.Env = os.Environ()

	connectStd(cmd)

	if err := cmd.Run(); err != nil {
		log.Printf("Failed to build for Windows: %v\n", err)
		return err
	}
	log.Println("Build for Windows completed successfully.")
	return nil
}

// buildAndroid builds the Android library for the specified architecture.
func buildAndroid(aOutPath, goArch, apiLevel, archName string) error {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)
	log.Printf("Building for Android %s...\n", archName)

	ndkPath := os.Getenv("ANDROID_NDK_HOME")
	osSuffix := ""
	if runtime.GOOS == "windows" {
		osSuffix = ".cmd" //
	}

	ccBasePath := ndkPath + "/toolchains/llvm/prebuilt/" + runtime.GOOS + "-x86_64/bin/"

	var cc string
	switch goArch {
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
		"GOARCH=" + goArch,
		"CC=" + cc,
	}
	buildTags := "android"
	cmd := exec.Command(
		"go",
		"build",
		"-buildmode=c-shared",
		"-trimpath",
		"-ldflags=-s -w",
		"-tags", buildTags,
		"-o", filepath.Join(aOutPath, archName, strings.Join([]string{soName, "so"}, ".")), ".")

	cmd.Dir = goSrc
	cmd.Env = append(os.Environ(), env...)

	connectStd(cmd)
	return cmd.Run()
}

func BuildHarmonyOS_API9() error {
	buildFunc := func(outPath string, arch string, archName string) error {
		log.SetOutput(os.Stdout)
		log.SetFlags(log.Lshortfile)
		log.Printf("Building for HarmonyOS API9 : %s...\n", archName)

		ndkPath := os.Getenv("HARMONY_NDK_API9")
		if len(ndkPath) == 0 {
			ndkPath = "D:/OpenHarmonySDK/9/native"
			log.Println("not find env variable: HARMONY_NDK_API9. Use Default:", ndkPath)
		}

		goBinFullPath := os.Getenv("HARMONY_GO_BIN")
		if len(goBinFullPath) == 0 {
			goBinFullPath = "go" // Default Go binary path
			log.Println("not find env variable: HARMONY_GO_BIN ,Use Default:", goBinFullPath)
		}
		GOOS := "android"
		baseFlag := fmt.Sprintf("--sysroot=%s/sysroot -fdata-sections -ffunction-sections -funwind-tables -fstack-protector-strong -no-canonical-prefixes -fno-addrsig -Wformat -Werror=format-security  -D__MUSL__ -fPIC -MD -MT -MF", ndkPath)
		var cc string
		var cFlag string
		switch arch {
		case "arm64":
			cFlag = fmt.Sprintf("--target=aarch64-linux-ohos %s", baseFlag)
		case "amd64":
			cFlag = fmt.Sprintf("--target=x86_64-linux-ohos %s", baseFlag)
		}
		cc = fmt.Sprintf("%s/llvm/bin/clang %s", ndkPath, cFlag)
		env := []string{
			"CGO_ENABLED=1",
			fmt.Sprintf("GOOS=%s", GOOS),
			fmt.Sprintf("GOARCH=%s", arch),
			fmt.Sprintf("NM=%s/llvm/bin/llvm-nm", ndkPath),
			fmt.Sprintf("AR=%s/llvm/bin/llvm-ar", ndkPath),
			fmt.Sprintf("LD=%s/llvm/bin/ld.lld", ndkPath),
			fmt.Sprintf("CC=%s", cc),
		}
		buildTags := "harmony"
		cmd := exec.Command(
			goBinFullPath,
			"build",
			"-buildmode=c-shared",
			"-trimpath",
			"-ldflags=-s -w",
			"-tags", buildTags,
			"-o", filepath.Join(outPath, archName, strings.Join([]string{soName, "so"}, ".")), ".")

		cmd.Dir = goSrc
		cmd.Env = append(os.Environ(), env...)

		connectStd(cmd)
		return cmd.Run()
	}
	architectures := []struct {
		GoArch, OutArch string
	}{
		{"arm64", "arm64-v8a"},
		{"amd64", "x86_64"},
	}
	output := filepath.Join(outPath, "HarmonyOS_API9")
	for _, arch := range architectures {
		if err := os.MkdirAll(filepath.Join(goSrc, output, arch.OutArch), 0755); err != nil {
			return err
		}

		if err := buildFunc(output, arch.GoArch, arch.OutArch); err != nil {
			log.Fatalf("Failed to build for  HarmanyOS API9 %s: %v\n", arch.OutArch, err)
		} else {
			log.Printf("Success to build for  HarmanyOS API9 %s\n", arch.OutArch)
		}
	}
	return nil
}

func BuildHarmonyOS_API12() error {
	buildFunc := func(outPath string, arch string, archName string) error {
		log.SetOutput(os.Stdout)
		log.SetFlags(log.Lshortfile)
		log.Printf("Building for HarmonyOS API12 : %s...\n", archName)

		ndkPath := os.Getenv("HARMONY_NDK_API12")
		if len(ndkPath) == 0 {
			ndkPath = "D:/OpenHarmonySDK/12/native"
			log.Println("not find env variable: HARMONY_NDK_API12 ,Use Default:", ndkPath)
		} else {
			log.Println("find env variable: HARMONY_NDK_API12 ,Use:", ndkPath)
		}
		ndkPath = "D:/Harmony/harmonyDevEcoStudio/DevEcoStudio/sdk/default/openharmony/native"

		goBinFullPath := os.Getenv("HARMONY_GO_BIN")
		if len(goBinFullPath) == 0 {
			goBinFullPath = "go" // Default Go binary path
			log.Println("not find env variable: HARMONY_GO_BIN ,Use Default:", goBinFullPath)
		} else {
			log.Println("find env variable: HARMONY_GO_BIN ,Use:", goBinFullPath)
		}
		baseFlag := fmt.Sprintf("--sysroot=%s/sysroot -fdata-sections -ffunction-sections -funwind-tables -fstack-protector-strong -no-canonical-prefixes -fno-addrsig -Wformat -Werror=format-security  -D__MUSL__ -fPIC -MD -MT -MF", ndkPath)
		var GOOS string
		var cc string
		var cFlag string
		var cmd *exec.Cmd
		buildTags := "harmony"
		switch arch {
		case "arm64":
			GOOS = "linux"
			cFlag = fmt.Sprintf("--target=aarch64-linux-ohos %s", baseFlag)
			cmd = exec.Command("D:/Harmony/su-go/go/bin/go",
				"build",
				"-buildmode=c-shared",
				"-tlsmodegd",
				"-trimpath",
				"-ldflags=-s -w",
				"-tags", buildTags,
				"-o", filepath.Join(outPath, archName, strings.Join([]string{soName, "so"}, ".")), ".")
		case "amd64":
			GOOS = "android"
			cFlag = fmt.Sprintf("--target=x86_64-linux-ohos %s", baseFlag)
			cmd = exec.Command(
				goBinFullPath,
				"build",
				"-buildmode=c-shared",
				"-trimpath",
				"-ldflags=-s -w",
				"-tags", buildTags,
				"-o", filepath.Join(outPath, archName, strings.Join([]string{soName, "so"}, ".")), ".")
		}
		cc = fmt.Sprintf("%s/llvm/bin/clang %s", ndkPath, cFlag)
		env := []string{
			"CGO_ENABLED=1",
			fmt.Sprintf("GOOS=%s", GOOS),
			fmt.Sprintf("GOARCH=%s", arch),
			fmt.Sprintf("NM=%s/llvm/bin/llvm-nm", ndkPath),
			fmt.Sprintf("AR=%s/llvm/bin/llvm-ar", ndkPath),
			fmt.Sprintf("LD=%s/llvm/bin/ld.lld", ndkPath),
			fmt.Sprintf("CC=%s", cc),
		}

		cmd.Dir = goSrc
		cmd.Env = append(os.Environ(), env...)
		connectStd(cmd)
		return cmd.Run()
	}
	architectures := []struct {
		GoArch, OutArch string
	}{
		{"arm64", "arm64-v8a"},
		{"amd64", "x86_64"},
	}
	output := filepath.Join(outPath, "HarmanyOS_API12")
	for _, arch := range architectures {
		if err := os.MkdirAll(filepath.Join(goSrc, output, arch.OutArch), 0755); err != nil {
			return err
		}
		if err := buildFunc(output, arch.GoArch, arch.OutArch); err != nil {
			log.Fatalf("Failed to build for  HarmanyOS API12 %s: %v\n", arch.OutArch, err)
		} else {
			log.Printf("Success to build for  HarmanyOS API12 %s\n", arch.OutArch)
		}
	}
	return nil
}

/*  Dependencies func */

func getWorkDirToolPath(name string) string {
	toolPath := ""
	workDir, err := os.Getwd()
	if err != nil {
		log.Println("Error", err)
		return toolPath
	}
	toolsPath := filepath.Join(workDir, "tools")
	filepath.Walk(toolsPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.TrimSuffix(info.Name(), filepath.Ext(info.Name())) == name {
			toolPath = path
		}
		return nil
	})
	return toolPath
}

func getToolPath(name string) (string, error) {
	toolPath := getWorkDirToolPath(name)
	if toolPath != "" {
		return toolPath, nil
	}
	if p, err := exec.LookPath(name); err == nil {
		return p, nil
	}
	// check under gopath
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	p := filepath.Join(gopath, "bin", name)

	if _, err := os.Stat(p); err != nil {
		return "", err
	}
	return p, nil
}

func getRustToolPath(name string) (string, error) {
	toolPath := getWorkDirToolPath(name)
	if toolPath != "" {
		return toolPath, nil
	}

	if p, err := exec.LookPath(name); err == nil {
		return p, nil
	}

	cargoHome := os.Getenv("CARGO_HOME")
	if cargoHome == "" {
		cargoHome = filepath.Join(os.Getenv("HOME"), ".cargo")
	}
	p := filepath.Join(cargoHome, "bin", name)

	if _, err := os.Stat(p); err != nil {
		return "", err
	}
	return p, nil
}

func connectStd(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}

func removeOmitemptyTags() error {
	protoGoDir := filepath.Join(protoDir, GO) // "./proto/go"

	re := regexp.MustCompile(`,\s*omitempty`)

	return filepath.Walk(protoGoDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("access path error:", err)
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".pb.go") {
			input, err := os.ReadFile(path)
			if err != nil {
				fmt.Println("ReadFile error. Path: %s, Error %v", path, err)
				return err
			}

			output := re.ReplaceAllString(string(input), "")

			// check replace is happened
			if string(input) != output {
				err = os.WriteFile(path, []byte(output), info.Mode())
				if err != nil {
					fmt.Printf("Error writing file: %s, error: %v\n", path, err)
					return err
				}
				// fmt.Println("Modified file:", path)
			}
		}

		return nil
	})
}
