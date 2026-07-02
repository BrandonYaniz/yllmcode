package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/BrandonYaniz/yllmcode/internal/service"
)

func main() {
	if err := run(context.Background(), os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string) error {
	if len(args) == 0 {
		printUsage(os.Stderr)
		return fmt.Errorf("missing command")
	}

	switch args[0] {
	case "init":
		return runInit(ctx, args[1:])
	case "help", "-h", "--help":
		printUsage(os.Stdout)
		return nil
	default:
		printUsage(os.Stderr)
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func runInit(ctx context.Context, args []string) error {
	flags := flag.NewFlagSet("init", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)

	var root string
	var force bool
	flags.StringVar(&root, "root", "", "project root to initialize")
	flags.BoolVar(&force, "force", false, "overwrite default yllmcode files")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("init accepts flags only")
	}

	svc := service.New()
	result, err := svc.InitProject(ctx, service.InitProjectRequest{
		Root:  root,
		Force: force,
	})
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "initialized %s\n", result.Root)
	for _, dir := range result.CreatedDirs {
		fmt.Fprintf(os.Stdout, "created dir  %s\n", dir)
	}
	for _, file := range result.CreatedFiles {
		fmt.Fprintf(os.Stdout, "created file %s\n", file)
	}
	for _, file := range result.SkippedFiles {
		fmt.Fprintf(os.Stdout, "skipped file %s\n", file)
	}

	return nil
}

func printUsage(out *os.File) {
	fmt.Fprintln(out, "usage:")
	fmt.Fprintln(out, "  yllmcode init [--root path] [--force]")
}
