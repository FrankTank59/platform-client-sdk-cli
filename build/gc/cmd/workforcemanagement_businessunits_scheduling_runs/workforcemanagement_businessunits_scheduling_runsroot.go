package workforcemanagement_businessunits_scheduling_runs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_businessunits_scheduling_runs_result"
)

func init() {
	workforcemanagement_businessunits_scheduling_runsCmd.AddCommand(workforcemanagement_businessunits_scheduling_runs_result.Cmdworkforcemanagement_businessunits_scheduling_runs_result())
	workforcemanagement_businessunits_scheduling_runsCmd.Short = utils.GenerateCustomDescription(workforcemanagement_businessunits_scheduling_runsCmd.Short, workforcemanagement_businessunits_scheduling_runs_result.Description, )
	workforcemanagement_businessunits_scheduling_runsCmd.Long = workforcemanagement_businessunits_scheduling_runsCmd.Short
}
