// Monorepo level CI/CD for packages

package main

import (
	"context"
)

type Ci struct{}

// Test all packages
func (m *Ci) Test(ctx context.Context, rootDir *Directory) (string, error) {
	return dag.Tapservicego().Test(ctx, rootDir.Directory("tapservicego"))
}

// Build all packages
func (m *Ci) Build(ctx context.Context, rootDir *Directory) (string, error) {
	return dag.Tapservicego().Build(rootDir.Directory("tapservicego"), 8080).Stdout(ctx)
}

// Publish container images
func (m *Ci) PublishImages(ctx context.Context, rootDir *Directory, username string, password *Secret, tags string) (string, error) {
	images, err := m.publishTapservice(ctx, rootDir, username, password, tags)
	if err != nil {
		return "", err
	}
	result := "Published images:\n"
	for _, image := range images {
		result += image + "\n"
	}
	return result, nil
}

// Publish Helm Charts
func (m *Ci) PublishHelmCharts(
	ctx context.Context,
	rootDir *Directory,
	username string,
	password *Secret,
	ghOrg *string,
) (string, error) {
	var result string
	output, err := m.publishTapserviceHelmChart(ctx, rootDir, username, password, ghOrg)
	if err != nil {
		return "", err
	}
	result += output + "\n"
	return result, nil
}

func (m *Ci) publishTapservice(
	ctx context.Context,
	rootDir *Directory,
	username string,
	password *Secret,
	tags string,
) ([]string, error) {
	tapOptions := TapservicegoPublishContainerOpts{
		Username: username,
		Tags:     tags,
	}
	images, err := dag.Tapservicego().PublishContainer(
		ctx,
		rootDir.Directory("tapservicego"),
		password,
		tapOptions,
	)
	if err != nil {
		return []string{}, err
	}
	return images, nil
}

func (m *Ci) publishTapserviceHelmChart(
	ctx context.Context,
	rootDir *Directory,
	username string,
	password *Secret,
	ghOrg *string,
) (string, error) {
	chartDir := rootDir.Directory("tapservicego").Directory("deployments/tapservice")
	var org string
	if ghOrg == nil {
		org = username
	}
	opts := TapservicegoPublishHelmChartOpts{
		GhOrg: org,
	}
	return dag.Tapservicego().PublishHelmChart(ctx, chartDir, username, password, opts)
}
