package conversations_screenshares_participants_communications

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_screenshares_participants_communications", "SWAGGER_OVERRIDE_/api/v2/conversations/screenshares/{conversationId}/participants/{participantId}/communications")
	conversations_screenshares_participants_communicationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_screenshares_participants_communications"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_screenshares_participants_communicationsCmd)
}

func Cmdconversations_screenshares_participants_communications() *cobra.Command {
	return conversations_screenshares_participants_communicationsCmd
}
