package workforcemanagement_managementunits_workplans

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_workplans_validate"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_workplans_copy"
)

func init() {
	workforcemanagement_managementunits_workplansCmd.AddCommand(workforcemanagement_managementunits_workplans_validate.Cmdworkforcemanagement_managementunits_workplans_validate())
	workforcemanagement_managementunits_workplansCmd.AddCommand(workforcemanagement_managementunits_workplans_copy.Cmdworkforcemanagement_managementunits_workplans_copy())
	workforcemanagement_managementunits_workplansCmd.Short = utils.GenerateCustomDescription(workforcemanagement_managementunits_workplansCmd.Short, workforcemanagement_managementunits_workplans_validate.Description, workforcemanagement_managementunits_workplans_copy.Description, )
	workforcemanagement_managementunits_workplansCmd.Long = workforcemanagement_managementunits_workplansCmd.Short
}
