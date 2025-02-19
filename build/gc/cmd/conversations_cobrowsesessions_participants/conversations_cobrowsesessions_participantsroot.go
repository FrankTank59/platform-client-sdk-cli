package conversations_cobrowsesessions_participants

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_cobrowsesessions_participants_wrapupcodes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_cobrowsesessions_participants_attributes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_cobrowsesessions_participants_wrapup"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_cobrowsesessions_participants_communications"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_cobrowsesessions_participants_replace"
)

func init() {
	conversations_cobrowsesessions_participantsCmd.AddCommand(conversations_cobrowsesessions_participants_wrapupcodes.Cmdconversations_cobrowsesessions_participants_wrapupcodes())
	conversations_cobrowsesessions_participantsCmd.AddCommand(conversations_cobrowsesessions_participants_attributes.Cmdconversations_cobrowsesessions_participants_attributes())
	conversations_cobrowsesessions_participantsCmd.AddCommand(conversations_cobrowsesessions_participants_wrapup.Cmdconversations_cobrowsesessions_participants_wrapup())
	conversations_cobrowsesessions_participantsCmd.AddCommand(conversations_cobrowsesessions_participants_communications.Cmdconversations_cobrowsesessions_participants_communications())
	conversations_cobrowsesessions_participantsCmd.AddCommand(conversations_cobrowsesessions_participants_replace.Cmdconversations_cobrowsesessions_participants_replace())
	conversations_cobrowsesessions_participantsCmd.Short = utils.GenerateCustomDescription(conversations_cobrowsesessions_participantsCmd.Short, conversations_cobrowsesessions_participants_wrapupcodes.Description, conversations_cobrowsesessions_participants_attributes.Description, conversations_cobrowsesessions_participants_wrapup.Description, conversations_cobrowsesessions_participants_communications.Description, conversations_cobrowsesessions_participants_replace.Description, )
	conversations_cobrowsesessions_participantsCmd.Long = conversations_cobrowsesessions_participantsCmd.Short
}
