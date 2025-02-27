package conversations_messages_participants

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messages_participants_wrapupcodes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messages_participants_attributes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messages_participants_wrapup"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messages_participants_communications"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messages_participants_replace"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messages_participants_monitor"
)

func init() {
	conversations_messages_participantsCmd.AddCommand(conversations_messages_participants_wrapupcodes.Cmdconversations_messages_participants_wrapupcodes())
	conversations_messages_participantsCmd.AddCommand(conversations_messages_participants_attributes.Cmdconversations_messages_participants_attributes())
	conversations_messages_participantsCmd.AddCommand(conversations_messages_participants_wrapup.Cmdconversations_messages_participants_wrapup())
	conversations_messages_participantsCmd.AddCommand(conversations_messages_participants_communications.Cmdconversations_messages_participants_communications())
	conversations_messages_participantsCmd.AddCommand(conversations_messages_participants_replace.Cmdconversations_messages_participants_replace())
	conversations_messages_participantsCmd.AddCommand(conversations_messages_participants_monitor.Cmdconversations_messages_participants_monitor())
	conversations_messages_participantsCmd.Short = utils.GenerateCustomDescription(conversations_messages_participantsCmd.Short, conversations_messages_participants_wrapupcodes.Description, conversations_messages_participants_attributes.Description, conversations_messages_participants_wrapup.Description, conversations_messages_participants_communications.Description, conversations_messages_participants_replace.Description, conversations_messages_participants_monitor.Description, )
	conversations_messages_participantsCmd.Long = conversations_messages_participantsCmd.Short
}
