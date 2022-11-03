package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/databricks-sdk-go/workspaces"
)

func main() {
	w := workspaces.New(&databricks.Config{
		Host:              askFor("Host:"),
		AzureResourceID:   askFor("Azure Resource ID:"),
		AzureTenantID:     askFor("AAD Tenant ID:"),
		AzureClientID:     askFor("AAD Client ID:"),
		AzureClientSecret: askFor("AAD Client Secret:"),
		Credentials:       databricks.AzureClientSecretCredentials{},
	})
	all, err := w.Clusters.ListAll(context.Background(), clusters.ListRequest{})
	if err != nil {
		panic(err)
	}
	for _, c := range all {
		println(c.ClusterName)
	}
}

func askFor(prompt string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, prompt+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}