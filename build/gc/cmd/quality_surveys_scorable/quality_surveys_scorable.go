package quality_surveys_scorable

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
	Description = utils.FormatUsageDescription("quality_surveys_scorable", "SWAGGER_OVERRIDE_/api/v2/quality/surveys/scorable", "SWAGGER_OVERRIDE_/api/v2/quality/surveys/scorable", )
	quality_surveys_scorableCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("quality_surveys_scorable"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(quality_surveys_scorableCmd)
}

func Cmdquality_surveys_scorable() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "customerSurveyUrl", "", "customerSurveyUrl - REQUIRED")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/quality/surveys/scorable", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	getCmd.MarkFlagRequired("customerSurveyUrl")
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ScorableSurvey&quot;
  }
}`)
	quality_surveys_scorableCmd.AddCommand(getCmd)
	
	utils.AddFlag(updateCmd.Flags(), "string", "customerSurveyUrl", "", "customerSurveyUrl - REQUIRED")
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/quality/surveys/scorable", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;survey&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ScorableSurvey&quot;
  }
}`)
	updateCmd.MarkFlagRequired("customerSurveyUrl")
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ScorableSurvey&quot;
  }
}`)
	quality_surveys_scorableCmd.AddCommand(updateCmd)
	
	return quality_surveys_scorableCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a survey as an end-customer, for the purposes of scoring it.",
	Long:  "Get a survey as an end-customer, for the purposes of scoring it.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/quality/surveys/scorable"

		customerSurveyUrl := utils.GetFlag(cmd.Flags(), "string", "customerSurveyUrl")
		if customerSurveyUrl != "" {
			queryParams["customerSurveyUrl"] = customerSurveyUrl
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", urlString, cmd.Flags())
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
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a survey as an end-customer, for the purposes of scoring it.",
	Long:  "Update a survey as an end-customer, for the purposes of scoring it.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/quality/surveys/scorable"

		customerSurveyUrl := utils.GetFlag(cmd.Flags(), "string", "customerSurveyUrl")
		if customerSurveyUrl != "" {
			queryParams["customerSurveyUrl"] = customerSurveyUrl
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PUT", urlString, cmd.Flags())
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
