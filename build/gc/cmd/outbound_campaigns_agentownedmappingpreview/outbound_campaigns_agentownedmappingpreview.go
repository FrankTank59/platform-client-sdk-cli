package outbound_campaigns_agentownedmappingpreview

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
	Description = utils.FormatUsageDescription("outbound_campaigns_agentownedmappingpreview", "SWAGGER_OVERRIDE_/api/v2/outbound/campaigns/{campaignId}/agentownedmappingpreview", )
	outbound_campaigns_agentownedmappingpreviewCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_campaigns_agentownedmappingpreview"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_campaigns_agentownedmappingpreviewCmd)
}

func Cmdoutbound_campaigns_agentownedmappingpreview() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/outbound/campaigns/{campaignId}/agentownedmappingpreview", utils.FormatPermissions([]string{ "outbound:campaign:view", "outbound:contact:view", "directory:user:view",  })))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Empty&quot;
  }
}`)
	outbound_campaigns_agentownedmappingpreviewCmd.AddCommand(createCmd)
	
	return outbound_campaigns_agentownedmappingpreviewCmd
}

var createCmd = &cobra.Command{
	Use:   "create [campaignId]",
	Short: "Initiate request for a preview of how agents will be mapped to this campaign`s contact list.",
	Long:  "Initiate request for a preview of how agents will be mapped to this campaign`s contact list.",
	Args:  utils.DetermineArgs([]string{ "campaignId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/outbound/campaigns/{campaignId}/agentownedmappingpreview"
		campaignId, args := args[0], args[1:]
		path = strings.Replace(path, "{campaignId}", fmt.Sprintf("%v", campaignId), -1)

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
