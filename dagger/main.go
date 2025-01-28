// A generated module for DaggerGolangCalc functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/dagger-golang-calc/internal/dagger"
)

type DaggerGolangCalc struct{}

// Test runs unit tests for the Go application
func (m *DaggerGolangCalc) Test(ctx context.Context, source *dagger.Directory) (string, error) {
	return m.BuildEnv(source).
		WithExec([]string{"go", "test", "./..."}).
		Stdout(ctx)
}

// BuildEnv sets up a development environment for the Go application
func (m *DaggerGolangCalc) BuildEnv(source *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("golang:latest").
		WithDirectory("/src", source).
		WithWorkdir("/src").
		WithExec([]string{"go", "mod", "download"})
}
