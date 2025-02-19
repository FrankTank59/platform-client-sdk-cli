package conversations_chats_participants

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_chats_participants_wrapupcodes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_chats_participants_attributes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_chats_participants_wrapup"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_chats_participants_communications"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_chats_participants_replace"
)

func init() {
	conversations_chats_participantsCmd.AddCommand(conversations_chats_participants_wrapupcodes.Cmdconversations_chats_participants_wrapupcodes())
	conversations_chats_participantsCmd.AddCommand(conversations_chats_participants_attributes.Cmdconversations_chats_participants_attributes())
	conversations_chats_participantsCmd.AddCommand(conversations_chats_participants_wrapup.Cmdconversations_chats_participants_wrapup())
	conversations_chats_participantsCmd.AddCommand(conversations_chats_participants_communications.Cmdconversations_chats_participants_communications())
	conversations_chats_participantsCmd.AddCommand(conversations_chats_participants_replace.Cmdconversations_chats_participants_replace())
	conversations_chats_participantsCmd.Short = utils.GenerateCustomDescription(conversations_chats_participantsCmd.Short, conversations_chats_participants_wrapupcodes.Description, conversations_chats_participants_attributes.Description, conversations_chats_participants_wrapup.Description, conversations_chats_participants_communications.Description, conversations_chats_participants_replace.Description, )
	conversations_chats_participantsCmd.Long = conversations_chats_participantsCmd.Short
}
