package architect_dependencytracking

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_dependencytracking_types"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_dependencytracking_consumedresources"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_dependencytracking_object"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_dependencytracking_consumingresources"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_dependencytracking_deletedresourceconsumers"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_dependencytracking_build"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_dependencytracking_updatedresourceconsumers"
)

func init() {
	architect_dependencytrackingCmd.AddCommand(architect_dependencytracking_types.Cmdarchitect_dependencytracking_types())
	architect_dependencytrackingCmd.AddCommand(architect_dependencytracking_consumedresources.Cmdarchitect_dependencytracking_consumedresources())
	architect_dependencytrackingCmd.AddCommand(architect_dependencytracking_object.Cmdarchitect_dependencytracking_object())
	architect_dependencytrackingCmd.AddCommand(architect_dependencytracking_consumingresources.Cmdarchitect_dependencytracking_consumingresources())
	architect_dependencytrackingCmd.AddCommand(architect_dependencytracking_deletedresourceconsumers.Cmdarchitect_dependencytracking_deletedresourceconsumers())
	architect_dependencytrackingCmd.AddCommand(architect_dependencytracking_build.Cmdarchitect_dependencytracking_build())
	architect_dependencytrackingCmd.AddCommand(architect_dependencytracking_updatedresourceconsumers.Cmdarchitect_dependencytracking_updatedresourceconsumers())
	architect_dependencytrackingCmd.Short = utils.GenerateCustomDescription(architect_dependencytrackingCmd.Short, architect_dependencytracking_types.Description, architect_dependencytracking_consumedresources.Description, architect_dependencytracking_object.Description, architect_dependencytracking_consumingresources.Description, architect_dependencytracking_deletedresourceconsumers.Description, architect_dependencytracking_build.Description, architect_dependencytracking_updatedresourceconsumers.Description, )
	architect_dependencytrackingCmd.Long = architect_dependencytrackingCmd.Short
}
