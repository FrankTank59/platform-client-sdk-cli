package conversations_videos

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_videos_recordingstate"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_videos_participants"
)

func init() {
	conversations_videosCmd.AddCommand(conversations_videos_recordingstate.Cmdconversations_videos_recordingstate())
	conversations_videosCmd.AddCommand(conversations_videos_participants.Cmdconversations_videos_participants())
	conversations_videosCmd.Short = utils.GenerateCustomDescription(conversations_videosCmd.Short, conversations_videos_recordingstate.Description, conversations_videos_participants.Description, )
	conversations_videosCmd.Long = conversations_videosCmd.Short
}
