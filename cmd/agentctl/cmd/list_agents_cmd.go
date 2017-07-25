package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ligato/vpp-agent/cmd/agentctl/utils"
	"github.com/ligato/cn-infra/servicelabel"
	"github.com/ligato/cn-infra/statuscheck/model/status"
)

var listAgents = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "li"},
	Short:   "List agents recorded in Etcd (active or inactive)",
	Long: `
List all Agents for which records exist in Etcd. A record may be
configuration for an agent that is stored in Etcd by a 3rd party
entity, or state stored in Etcd by an agent.

Agents for which only configuration records exist (i.e. they do
not push status records into Etcd) are listed as 'INACTIVE'.`,
	Run: listAgentsFunc,
}

func init() {
	RootCmd.AddCommand(listAgents)
	listAgents.Flags().BoolVar(&showEtcd, "etcd", false,
		"Show Etcd Metadata (revision, key))")
}

func listAgentsFunc(cmd *cobra.Command, args []string) {
	db, err := utils.GetDbForAllAgents(globalFlags.Endpoints)
	if err != nil {
		utils.ExitWithError(utils.ExitError, errors.New("Failed connect to Etcd - "+err.Error()))
	}

	keyIter, err := db.ListKeys(servicelabel.GetAllAgentsPrefix())
	if err != nil {
		utils.ExitWithError(utils.ExitError, errors.New("Failed to get keys - "+err.Error()))
	}

	ed := utils.NewEtcdDump()
	for {
		if key, _, done := keyIter.GetNext(); !done {
			found, err := ed.ReadDataFromDb(db, key, nil, []string{status.StatusPrefix})
			if err != nil {
				utils.ExitWithError(utils.ExitError, err)
			} else if !found {
				label, _, _, _ := utils.ParseKey(key)
				if _, ok := ed[label]; !ok {
					ed.CreateEmptyRecord(key)
				}
			}
			continue
		}
		break
	}

	if len(ed) > 0 {
		buffer := ed.PrintDataAsText(showEtcd, false)
		fmt.Print(buffer.String())
	} else {
		fmt.Print("No data found.\n")
	}
}
