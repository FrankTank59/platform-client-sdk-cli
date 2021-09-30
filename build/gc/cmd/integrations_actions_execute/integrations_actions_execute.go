package integrations_actions_execute

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("integrations_actions_execute", "SWAGGER_OVERRIDE_/api/v2/integrations/actions/{actionId}/execute", )
	integrations_actions_executeCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_actions_execute"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_actions_executeCmd)
}

func Cmdintegrations_actions_execute() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/integrations/actions/{actionId}/execute", utils.FormatPermissions([]string{ "integrations:action:execute", "bridge:actions:execute",  }), utils.GenerateDevCentreLink("POST", "Integrations", "/api/v2/integrations/actions/{actionId}/execute")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "Map of parameters used for variable substitution.",
  "required" : true,
  "schema" : {
    "type" : "object",
    "additionalProperties" : {
      "type" : "object",
      "properties" : { }
    }
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "type" : "object",
    "properties" : { }
  }
}`)
	integrations_actions_executeCmd.AddCommand(createCmd)
	
	return integrations_actions_executeCmd
}

var createCmd = &cobra.Command{
	Use:   "create [actionId]",
	Short: "Execute Action and return response from 3rd party.  Responses will follow the schemas defined on the Action for success and error.",
	Long:  "Execute Action and return response from 3rd party.  Responses will follow the schemas defined on the Action for success and error.",
	Args:  utils.DetermineArgs([]string{ "actionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Interface{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/integrations/actions/{actionId}/execute"
		actionId, args := args[0], args[1:]
		path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)


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
