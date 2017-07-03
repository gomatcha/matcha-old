package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gomatcha.io/matcha/cmd"
)

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var RootCmd = &cobra.Command{
	Use:   "matcha",
	Short: "Matcha is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var (
	buildN       bool   // -n
	buildX       bool   // -x
	buildV       bool   // -v
	buildWork    bool   // -work
	buildGcflags string // -gcflags
	buildLdflags string // -ldflags
	buildO       string // -o
)

func init() {
	flags := InitCmd.Flags()
	flags.BoolVar(&buildN, "n", false, "print the commands but do not run them.")
	flags.BoolVar(&buildX, "x", false, "print the commands.")
	flags.BoolVar(&buildV, "v", false, "print the names of packages as they are compiled.")
	flags.BoolVar(&buildWork, "work", false, "print the name of the temporary work directory and do not delete it when exiting.")
	flags.StringVar(&buildGcflags, "gcflags", "", "arguments to pass on each go tool compile invocation.")
	flags.StringVar(&buildLdflags, "ldflags", "", "arguments to pass on each go tool link invocation.")

	RootCmd.AddCommand(InitCmd)
}

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "install mobile compiler toolchain",
	Long: `Init builds copies of the Go standard library for mobile devices.
It uses Xcode, if available, to build for iOS and uses the Android
NDK from the ndk-bundle SDK package or from the -ndk flag, to build
for Android.`,
	Run: func(command *cobra.Command, args []string) {
		flags := &cmd.Flags{
			BuildN:       buildN,
			BuildX:       buildX,
			BuildV:       buildV,
			BuildWork:    buildWork,
			BuildGcflags: buildGcflags,
			BuildLdflags: buildLdflags,
		}
		if err := cmd.Init(flags); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	flags := DevCmd.Flags()
	flags.BoolVar(&buildN, "n", false, "print the commands but do not run them.")
	flags.BoolVar(&buildX, "x", false, "print the commands.")
	flags.BoolVar(&buildV, "v", false, "print the names of packages as they are compiled.")
	flags.BoolVar(&buildWork, "work", false, "print the name of the temporary work directory and do not delete it when exiting.")
	flags.StringVar(&buildGcflags, "gcflags", "", "arguments to pass on each go tool compile invocation.")
	flags.StringVar(&buildLdflags, "ldflags", "", "arguments to pass on each go tool link invocation.")
	flags.StringVar(&buildO, "output", "", "forces build to write the resulting object to the named output file.")

	RootCmd.AddCommand(DevCmd)
}

var DevCmd = &cobra.Command{
	Use:   "dev",
	Short: "internal dev command",
	Long:  ``,
	Run: func(command *cobra.Command, args []string) {
		flags := &cmd.Flags{
			BuildN:       buildN,
			BuildX:       buildX,
			BuildV:       buildV,
			BuildWork:    buildWork,
			BuildGcflags: buildGcflags,
			BuildLdflags: buildLdflags,
			BuildO:       buildO,
		}
		if err := cmd.Bind(flags, args, true); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	flags := BuildCmd.Flags()
	flags.BoolVar(&buildN, "n", false, "print the commands but do not run them.")
	flags.BoolVar(&buildX, "x", false, "print the commands.")
	flags.BoolVar(&buildV, "v", false, "print the names of packages as they are compiled.")
	flags.BoolVar(&buildWork, "work", false, "print the name of the temporary work directory and do not delete it when exiting.")
	flags.StringVar(&buildGcflags, "gcflags", "", "arguments to pass on each go tool compile invocation.")
	flags.StringVar(&buildLdflags, "ldflags", "", "arguments to pass on each go tool link invocation.")
	flags.StringVar(&buildO, "output", "", "forces build to write the resulting object to the named output file.")

	RootCmd.AddCommand(BuildCmd)
}

var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "install mobile compiler toolchain",
	Long:  ``,
	Run: func(command *cobra.Command, args []string) {
		flags := &cmd.Flags{
			BuildN:       buildN,
			BuildX:       buildX,
			BuildV:       buildV,
			BuildWork:    buildWork,
			BuildGcflags: buildGcflags,
			BuildLdflags: buildLdflags,
			BuildO:       buildO,
		}
		if err := cmd.Bind(flags, args, false); err != nil {
			fmt.Println(err)
		}
	},
}
