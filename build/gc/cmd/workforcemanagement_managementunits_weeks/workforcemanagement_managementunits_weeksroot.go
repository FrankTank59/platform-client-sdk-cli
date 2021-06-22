package workforcemanagement_managementunits_weeks

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_weeks_schedules"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_weeks_shifttrades"
)

func init() {
	workforcemanagement_managementunits_weeksCmd.AddCommand(workforcemanagement_managementunits_weeks_schedules.Cmdworkforcemanagement_managementunits_weeks_schedules())
	workforcemanagement_managementunits_weeksCmd.AddCommand(workforcemanagement_managementunits_weeks_shifttrades.Cmdworkforcemanagement_managementunits_weeks_shifttrades())
	workforcemanagement_managementunits_weeksCmd.Short = utils.GenerateCustomDescription(workforcemanagement_managementunits_weeksCmd.Short, workforcemanagement_managementunits_weeks_schedules.Description, workforcemanagement_managementunits_weeks_shifttrades.Description, )
	workforcemanagement_managementunits_weeksCmd.Long = workforcemanagement_managementunits_weeksCmd.Short
}
