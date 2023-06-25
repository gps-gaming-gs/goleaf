package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gps-gaming-gs/goleaf/rpcx"
	"github.com/gps-gaming-gs/goleaf/tplx"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

const (
	codeFailure = 1
	dash        = "-"
	doubleDash  = "--"
	assign      = "="
)

var (
	rootCmd = &cobra.Command{
		Use:   "goleaf",
		Short: "A cli tool to generate go-zero code",
		Long: "A cli tool to generate api, zrpc, model code\n\n" +
			"GitHub: https://github.com/zeromicro/go-zero\n" +
			"Site:   https://go-zero.dev",
	}
)

func init() {
	rootCmd.AddCommand(tplx.Cmd)
	rootCmd.AddCommand(rpcx.Cmd)
	rootCmd.AddCommand(rpcx.CocosCmd)
}

func supportGoStdFlag(args []string) []string {
	copyArgs := append([]string(nil), args...)
	parentCmd, _, err := rootCmd.Traverse(args[:1])
	if err != nil { // ignore it to let cobra handle the error.
		return copyArgs
	}

	for idx, arg := range copyArgs[0:] {
		parentCmd, _, err = parentCmd.Traverse([]string{arg})
		if err != nil { // ignore it to let cobra handle the error.
			break
		}
		if !strings.HasPrefix(arg, dash) {
			continue
		}

		flagExpr := strings.TrimPrefix(arg, doubleDash)
		flagExpr = strings.TrimPrefix(flagExpr, dash)
		flagName, flagValue := flagExpr, ""
		assignIndex := strings.Index(flagExpr, assign)
		if assignIndex > 0 {
			flagName = flagExpr[:assignIndex]
			flagValue = flagExpr[assignIndex:]
		}

		if !isBuiltin(flagName) {
			// The method Flag can only match the user custom flags.
			f := parentCmd.Flag(flagName)
			if f == nil {
				continue
			}
			if f.Shorthand == flagName {
				continue
			}
		}

		goStyleFlag := doubleDash + flagName
		if assignIndex > 0 {
			goStyleFlag += flagValue
		}

		copyArgs[idx] = goStyleFlag
	}
	return copyArgs
}

func isBuiltin(name string) bool {
	return name == "version" || name == "help"
}

func main() {
	os.Args = supportGoStdFlag(os.Args)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(aurora.Red(err.Error()))
		os.Exit(codeFailure)
	}
}
