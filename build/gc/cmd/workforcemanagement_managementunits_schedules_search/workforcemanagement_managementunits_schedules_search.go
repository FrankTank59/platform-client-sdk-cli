package workforcemanagement_managementunits_schedules_search

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
	Description = utils.FormatUsageDescription("workforcemanagement_managementunits_schedules_search", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/schedules/search", )
	workforcemanagement_managementunits_schedules_searchCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_managementunits_schedules_search"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_managementunits_schedules_searchCmd)
}

func Cmdworkforcemanagement_managementunits_schedules_search() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/schedules/search", utils.FormatPermissions([]string{ "wfm:publishedSchedule:view", "wfm:schedule:view",  }), utils.GenerateDevCentreLink("POST", "Workforce Management", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/schedules/search")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "body",
  "required" : false,
  "schema" : {
    "$ref" : "#/definitions/UserListScheduleRequestBody"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/UserScheduleContainer"
  }
}`)
	workforcemanagement_managementunits_schedules_searchCmd.AddCommand(createCmd)
	
	return workforcemanagement_managementunits_schedules_searchCmd
}

var createCmd = &cobra.Command{
	Use:   "create [managementUnitId]",
	Short: "Query published schedules for given given time range for set of users",
	Long:  "Query published schedules for given given time range for set of users",
	Args:  utils.DetermineArgs([]string{ "managementUnitId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Userlistschedulerequestbody{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/managementunits/{managementUnitId}/schedules/search"
		managementUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)


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
