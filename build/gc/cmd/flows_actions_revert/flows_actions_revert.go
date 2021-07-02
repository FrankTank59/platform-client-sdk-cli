package flows_actions_revert

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("flows_actions_revert", "SWAGGER_OVERRIDE_/api/v2/flows/actions/revert", )
	flows_actions_revertCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_actions_revert"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_actions_revertCmd)
}

func Cmdflows_actions_revert() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "string", "flow", "", "Flow ID - REQUIRED")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/flows/actions/revert", utils.FormatPermissions([]string{ "architect:flow:edit",  }), utils.GenerateDevCentreLink("POST", "Architect", "/api/v2/flows/actions/revert")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	createCmd.MarkFlagRequired("flow")
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Flow&quot;
  }
}`)
	flows_actions_revertCmd.AddCommand(createCmd)
	
	return flows_actions_revertCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Revert flow",
	Long:  "Revert flow",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/flows/actions/revert"

		flow := utils.GetFlag(cmd.Flags(), "string", "flow")
		if flow != "" {
			queryParams["flow"] = flow
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("POST", urlString, cmd.Flags())
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
