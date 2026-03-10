// An Intro module for Dagger
package main

import (
	"context"
	"dagger/intro/internal/dagger"
	"encoding/json"
	"errors"
	"fmt"
)

type Intro struct{}

// Install your first Toolchain
// +check
func (m *Intro) InstallIntro() error {
	// You're running this, so toolchain installed!
	return nil
}

// Login to Dagger Cloud with dagger login
// +check
func (m *Intro) LoginToCloud(ctx context.Context) error {
	cloudUrl, err := dag.Cloud().TraceURL(ctx)
	if err != nil || cloudUrl == "" {
		return errors.New("Run 'dagger login' to login to Dagger Cloud")
	}

	return nil
}

// Install your next toolchain
// +check
func (m *Intro) InstallMoreToolchains(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) error {
	message := ""
	hasToolchain := false

	// List of known toolchains.
	knownToolchains := []struct {
		indicator string
		name      string
		repo      string
	}{
		{
			".prettierrc",
			"prettier",
			"github.com/dagger/prettier",
		},
		{
			"playwright.config.js",
			"playwright",
			"github.com/dagger/playwright",
		},
		{
			"jest.config.js",
			"jest",
			"github.com/dagger/jest",
		},
		{
			".eslintrc.js",
			"eslint",
			"github.com/dagger/eslint",
		},
		{
			"go.mod",
			"go",
			"github.com/dagger/dagger/toolchains/go",
		},
	}

	// See if known toolchains need to be installed
	for _, tc := range knownToolchains {
		needsTc, hasTc, err := toolchainIndicator(ctx, tc.indicator, tc.name, source)
		if err == nil && needsTc && !hasTc {
			message += fmt.Sprintf("Install %s: dagger toolchain install %s\n", tc.name, tc.repo)
		}
		hasToolchain = hasToolchain || hasTc
	}

	if message == "" && hasToolchain {
		// You have installed all of the toolchains we know to look for
		return nil
	}

	if message == "" {
		// You need toolchains but we didnt find any indicators that we know about
		message += "Find some toolchains for your project.\nAsk in the Dagger discord for some relevant ones.\n"
	}

	return errors.New(message)
}

// See if a project needs a known toolchain based on project indicators
func toolchainIndicator(
	ctx context.Context,
	indicator, toolchain string,
	source *dagger.Directory,
) (bool, bool, error) {
	hasIndicator, err := source.Exists(ctx, indicator)
	if err != nil {
		return false, false, err
	}
	hasToolchain, err := modContains(ctx, source, toolchain)
	return hasIndicator, hasToolchain, err
}

// Find an object name in the top level of a module
func modContains(ctx context.Context, source *dagger.Directory, name string) (bool, error) {
	mod := struct {
		Toolchains []struct {
			Name   string
			Source string
		}
	}{}

	daggerJson, err := source.File("dagger.json").Contents(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to read dagger.json: %w", err)
	}

	err = json.Unmarshal([]byte(daggerJson), &mod)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal dagger.json: %w", err)
	}

	for _, tc := range mod.Toolchains {
		if tc.Name == name {
			return true, nil
		}
	}
	return false, nil
}
