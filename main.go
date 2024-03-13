package main

import (
	"context"
	"log"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

func main() {
	playbookOpts := &playbook.AnsiblePlaybookOptions{
		AskBecomePass: true,
		Connection:    "local",
	}

	playbookCmd := &playbook.AnsiblePlaybookCmd{
		Playbooks:       []string{"playbooks/playbook.yaml"},
		PlaybookOptions: playbookOpts,
	}

	exec := execute.NewDefaultExecute(execute.WithCmd(playbookCmd))

	if err := exec.Execute(context.Background()); err != nil {
		log.Println(err)
	}
}
