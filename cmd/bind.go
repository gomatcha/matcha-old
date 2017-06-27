package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"go/build"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func Bind(flags *Flags, args []string) error {
	// Make $WORK.
	tempdir, err := NewTmpDir(flags, "")
	if err != nil {
		return err
	}
	defer RemoveAll(flags, tempdir)

	// Get $GOPATH/pkg/gomobile.
	gomobilepath, err := GoMobilePath()
	if err != nil {
		return err
	}

	// Get toolchain version.
	installedVersion, err := ReadFile(flags, filepath.Join(gomobilepath, "version"))
	if err != nil {
		return errors.New("toolchain partially installed, run `gomobile init`")
	}

	// Get go version.
	goVersion, err := GoVersion(flags)
	if err != nil {
		return err
	}

	// Check toolchain matches go version.
	if !bytes.Equal(installedVersion, goVersion) {
		return errors.New("toolchain out of date, run `gomobile init`")
	}

	// Get current working directory.
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// Create a build context.
	ctx := build.Default
	ctx.GOARCH = "arm"
	ctx.GOOS = "darwin"
	ctx.BuildTags = append(ctx.BuildTags, "ios")

	// Get packages to be built.
	pkgs := []*build.Package{}
	if len(args) == 0 {
		pkg, err := ctx.ImportDir(cwd, build.ImportComment)
		if err != nil {
			return err
		}
		pkgs = append(pkgs, pkg)
	} else {
		for _, a := range args {
			a = path.Clean(a)
			pkg, err := ctx.Import(a, cwd, build.ImportComment)
			if err != nil {
				return err
			}
			pkgs = append(pkgs, pkg)
		}
	}

	// Check if any of the package is main.
	for _, pkg := range pkgs {
		if pkg.Name == "main" {
			return fmt.Errorf("binding 'main' package (%s) is not supported", pkg.ImportComment)
		}
	}

	title := "Matcha"
	genDir := filepath.Join(tempdir, "gen")
	frameworkDir := title + ".framework"
	// frameworkDir := flags.BuildO
	// if frameworkDir != "" && !strings.HasSuffix(frameworkDir, ".framework") {
	// 	return fmt.Errorf("static framework name %q missing .framework suffix", frameworkDir)
	// }
	// if frameworkDir == "" {
	// 	frameworkDir = title + ".framework"
	// }

	// Build the "matcha/bridge" dir
	bridgeDir := filepath.Join(genDir, "src", "github.com", "overcyn", "matchabridge")
	if err := Mkdir(flags, bridgeDir); err != nil {
		return err
	}

	// Create the "main" go package, that references the other go packages
	mainPath := filepath.Join(tempdir, "src", "iosbin", "main.go")
	err = WriteFile(flags, mainPath, func(w io.Writer) error {
		format := fmt.Sprintf(BindFile, args[0]) // TODO(KD): Should this be args[0] or should it use the logic to generate pkgs
		_, err := w.Write([]byte(format))
		return err
	})
	if err != nil {
		return fmt.Errorf("failed to create the binding package for iOS: %v", err)
	}

	// Get the supporting files
	objcPkg, err := ctx.Import("github.com/overcyn/matcha/cmd", "", build.FindOnly)
	if err != nil {
		return err
	}
	if err := CopyFile(flags, filepath.Join(bridgeDir, "matchaobjc.h"), filepath.Join(objcPkg.Dir, "matchaobjc.h.support")); err != nil {
		return err
	}
	if err := CopyFile(flags, filepath.Join(bridgeDir, "matchaobjc.m"), filepath.Join(objcPkg.Dir, "matchaobjc.m.support")); err != nil {
		return err
	}
	if err := CopyFile(flags, filepath.Join(bridgeDir, "matchaobjc.go"), filepath.Join(objcPkg.Dir, "matchaobjc.go.support")); err != nil {
		return err
	}
	if err := CopyFile(flags, filepath.Join(bridgeDir, "matchago.h"), filepath.Join(objcPkg.Dir, "matchago.h.support")); err != nil {
		return err
	}
	if err := CopyFile(flags, filepath.Join(bridgeDir, "matchago.m"), filepath.Join(objcPkg.Dir, "matchago.m.support")); err != nil {
		return err
	}
	if err := CopyFile(flags, filepath.Join(bridgeDir, "matchago.go"), filepath.Join(objcPkg.Dir, "matchago.go.support")); err != nil {
		return err
	}

	// Build static framework output directory.
	if err := RemoveAll(flags, frameworkDir); err != nil {
		return err
	}

	// Build framework directory structure.
	headersDir := filepath.Join(frameworkDir, "Versions", "A", "Headers")
	resourcesDir := filepath.Join(frameworkDir, "Versions", "A", "Resources")
	modulesDir := filepath.Join(frameworkDir, "Versions", "A", "Modules")
	binaryPath := filepath.Join(frameworkDir, "Versions", "A", title)
	if err := Mkdir(flags, headersDir); err != nil {
		return err
	}
	if err := Mkdir(flags, resourcesDir); err != nil {
		return err
	}
	if err := Mkdir(flags, modulesDir); err != nil {
		return err
	}
	if err := Symlink(flags, "A", filepath.Join(frameworkDir, "Versions", "Current")); err != nil {
		return err
	}
	if err := Symlink(flags, filepath.Join("Versions", "Current", "Headers"), filepath.Join(frameworkDir, "Headers")); err != nil {
		return err
	}
	if err := Symlink(flags, filepath.Join("Versions", "Current", "Resources"), filepath.Join(frameworkDir, "Resources")); err != nil {
		return err
	}
	if err := Symlink(flags, filepath.Join("Versions", "Current", "Modules"), filepath.Join(frameworkDir, "Modules")); err != nil {
		return err
	}
	if err := Symlink(flags, filepath.Join("Versions", "Current", title), filepath.Join(frameworkDir, title)); err != nil {
		return err
	}

	// Copy in headers.
	if err = CopyFile(flags, filepath.Join(headersDir, "matchaobjc.h"), filepath.Join(bridgeDir, "matchaobjc.h")); err != nil {
		return err
	}
	if err = CopyFile(flags, filepath.Join(headersDir, "matchago.h"), filepath.Join(bridgeDir, "matchago.h")); err != nil {
		return err
	}

	// Copy in resources.
	if err := ioutil.WriteFile(filepath.Join(resourcesDir, "Info.plist"), []byte(InfoPlist), 0666); err != nil {
		return err
	}

	// Write modulemap.
	err = WriteModuleMap(flags, filepath.Join(modulesDir, "module.modulemap"), title)
	if err != nil {
		return err
	}

	// Build platform binaries concurrently.
	matchaDarwinArmEnv, err := DarwinArmEnv(flags)
	if err != nil {
		return err
	}

	matchaDarwinArm64Env, err := DarwinArm64Env(flags)
	if err != nil {
		return err
	}

	matchaDarwinAmd64Env, err := DarwinAmd64Env(flags)
	if err != nil {
		return err
	}

	type archPath struct {
		arch string
		path string
		err  error
	}
	archChan := make(chan archPath)
	for _, i := range [][]string{matchaDarwinArmEnv, matchaDarwinArm64Env, matchaDarwinAmd64Env} {
		go func(env []string) {
			arch := Getenv(env, "GOARCH")
			env = append(env, "GOPATH="+genDir+string(filepath.ListSeparator)+os.Getenv("GOPATH"))
			path := filepath.Join(tempdir, "matcha-"+arch+".a")
			err := GoBuild(flags, mainPath, env, ctx, tempdir, "-buildmode=c-archive", "-o", path)
			archChan <- archPath{arch, path, err}
		}(i)
	}
	archs := []archPath{}
	for i := 0; i < 3; i++ {
		arch := <-archChan
		if arch.err != nil {
			return arch.err
		}
		archs = append(archs, arch)
	}

	// Lipo to build fat binary.
	cmd := exec.Command("xcrun", "lipo", "-create")
	for _, i := range archs {
		cmd.Args = append(cmd.Args, "-arch", ArchClang(i.arch), i.path)
	}
	cmd.Args = append(cmd.Args, "-o", binaryPath)
	return RunCmd(flags, tempdir, cmd)
}

var InfoPlist = `<?xml version="1.0" encoding="UTF-8"?>
    <!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
    <plist version="1.0">
      <dict>
      </dict>
    </plist>
`

var BindFile = `
package main

import (
    _ "github.com/overcyn/matchabridge"
    _ "%s"
)

import "C"

func main() {}
`

var ModuleMapTemplate = template.Must(template.New("iosmmap").Parse(`framework module "{{.Module}}" {
    // header "ref.h"
{{range .Headers}}    header "{{.}}"
{{end}}
    export *
}`))

func WriteModuleMap(flags *Flags, filename string, title string) error {
	// Write modulemap.
	var mmVals = struct {
		Module  string
		Headers []string
	}{
		Module:  title,
		Headers: []string{"matchaobjc.h", "matchago.h"},
	}
	err := WriteFile(flags, filename, func(w io.Writer) error {
		return ModuleMapTemplate.Execute(w, mmVals)
	})
	if err != nil {
		return err
	}
	return nil
}
