package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type GCPInstance struct {
	Name string `json:"name"`
	Zone string `json:"zone"`
}

func ssh(project string) error {
	instanceCmd := exec.Command("gcloud", "compute", "--project", project, "instances", "list", "--format", "json")
	instanceOut, err := instanceCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("gcloud compute instances list: %v, out: %s", err, string(instanceOut))
	}

	instances := []GCPInstance{}
	err = json.NewDecoder(bytes.NewReader(instanceOut)).Decode(&instances)
	if err != nil {
		return err
	}

	names := make([]string, 0, len(instances))
	for _, instance := range instances {
		names = append(names, instance.Name)
	}

	name := fzf(names)
	fmt.Println(name)
	for _, instance := range instances {
		fmt.Println(instance.Name)
		if instance.Name == name {
			zoneParts := strings.Split(instance.Zone, "/")
			zone := zoneParts[len(zoneParts)-1]
			cmd := exec.Command("gcloud", "compute", "ssh", "--tunnel-through-iap", "--project", project, "--zone", zone, name)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			if err != nil {
				return fmt.Errorf("gcloud compute ssh: %v", err)
			}
			return nil
		}
	}

	return fmt.Errorf("BUG: selected instance not found")
}
