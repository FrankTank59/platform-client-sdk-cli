package authorization_roles_users_add

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
	Description = utils.FormatUsageDescription("authorization_roles_users_add", "SWAGGER_OVERRIDE_/api/v2/authorization/roles/{roleId}/users/add", )
	authorization_roles_users_addCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("authorization_roles_users_add"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(authorization_roles_users_addCmd)
}

func Cmdauthorization_roles_users_add() *cobra.Command { 
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/authorization/roles/{roleId}/users/add", utils.FormatPermissions([]string{ "authorization:grant:add",  }), utils.GenerateDevCentreLink("PUT", "Authorization", "/api/v2/authorization/roles/{roleId}/users/add")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "in" : "body",
  "name" : "body",
  "description" : "List of user IDs",
  "required" : true,
  "schema" : {
    "type" : "array",
    "items" : {
      "type" : "string"
    }
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "schema" : {
    "type" : "array",
    "items" : {
      "type" : "string"
    }
  }
}`)
	authorization_roles_users_addCmd.AddCommand(updateCmd)
	
	return authorization_roles_users_addCmd
}

var updateCmd = &cobra.Command{
	Use:   "update [roleId]",
	Short: "Sets the users for the role",
	Long:  "Sets the users for the role",
	Args:  utils.DetermineArgs([]string{ "roleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/authorization/roles/{roleId}/users/add"
		roleId, args := args[0], args[1:]
		path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)


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
