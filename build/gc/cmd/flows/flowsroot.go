package flows

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_milestones"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_outcomes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_datatables"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_history"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_latestconfiguration"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_divisionviews"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_executions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_actions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_versions"
)

func init() {
	flowsCmd.AddCommand(flows_milestones.Cmdflows_milestones())
	flowsCmd.AddCommand(flows_outcomes.Cmdflows_outcomes())
	flowsCmd.AddCommand(flows_datatables.Cmdflows_datatables())
	flowsCmd.AddCommand(flows_history.Cmdflows_history())
	flowsCmd.AddCommand(flows_latestconfiguration.Cmdflows_latestconfiguration())
	flowsCmd.AddCommand(flows_divisionviews.Cmdflows_divisionviews())
	flowsCmd.AddCommand(flows_executions.Cmdflows_executions())
	flowsCmd.AddCommand(flows_actions.Cmdflows_actions())
	flowsCmd.AddCommand(flows_versions.Cmdflows_versions())
	flowsCmd.Short = utils.GenerateCustomDescription(flowsCmd.Short, flows_milestones.Description, flows_outcomes.Description, flows_datatables.Description, flows_history.Description, flows_latestconfiguration.Description, flows_divisionviews.Description, flows_executions.Description, flows_actions.Description, flows_versions.Description, )
	flowsCmd.Long = flowsCmd.Short
}
