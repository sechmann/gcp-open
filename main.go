package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/iterator"
)

func main() {
	// all_projects="$(gcloud projects list --format=json)"
	// tenants=$( (jq -r '.[] | .labels | select(has("tenant")) | .tenant' | sort | uniq) <<<"$all_projects")
	// tenant="$(fzf --ansi -1 -q "${1}" <<<"$tenants")"
	// projects=$( (jq -r ".[] | select(.labels.tenant == \"$tenant\")|.projectId") <<<"$all_projects")
	// project="$(fzf --ansi -1 -q "${1}" <<<"$projects")"
	// wl-copy -n <<< "$project"
	// echo "$project copied to clipboard"

	if err := run(); err != nil {
		log.Errorf("run: %v", err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()
	projectsClient, err := resourcemanager.NewProjectsClient(ctx)
	if err != nil {
		return err
	}

	projectsIt := projectsClient.SearchProjects(ctx, &resourcemanagerpb.SearchProjectsRequest{})
	tenantProjects := make(map[string]map[string]*resourcemanagerpb.Project)
	for project, err := projectsIt.Next(); err != iterator.Done; project, err = projectsIt.Next() {
		if err != nil {
			return err
		}
		tenant := project.GetLabels()["tenant"]

		if tenantProjects[tenant] == nil {
			tenantProjects[tenant] = make(map[string]*resourcemanagerpb.Project)
		}

		tenantProjects[tenant][project.GetProjectId()] = project
	}

	tenants := make([]string, len(tenantProjects))
	i := 0
	for tenant := range tenantProjects {
		tenants[i] = tenant
		i++
	}

	tenant := fzf(tenants)
	project := fzf(keys(tenantProjects[tenant]))
	url := productUrls[fzf(keys(productUrls))]

	fmt.Printf("https://console.cloud.google.com%s?project=%s\n", url, project)

	return nil
}

func fzf(l []string) string {
	shell := os.Getenv("SHELL")
	if len(shell) == 0 {
		shell = "sh"
	}

	cmd := exec.Command(shell, "-c", "fzf")
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Error("fzf: get stdin pipe: %w", err)
		return ""
	}

	for _, el := range l {
		fmt.Fprintln(stdin, el)
	}
	stdin.Close()
	result, err := cmd.Output()
	if err != nil {
		log.Error("fzf: get stdout output: %w", err)
		return ""
	}
	return string(bytes.TrimSpace(result))
}

func keys[T any](m map[string]T) []string {
	keys := make([]string, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}

	return keys
}
