package integrations_speech

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_dialogflow"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_lex"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_tts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_nuance"
)

func init() {
	integrations_speechCmd.AddCommand(integrations_speech_dialogflow.Cmdintegrations_speech_dialogflow())
	integrations_speechCmd.AddCommand(integrations_speech_lex.Cmdintegrations_speech_lex())
	integrations_speechCmd.AddCommand(integrations_speech_tts.Cmdintegrations_speech_tts())
	integrations_speechCmd.AddCommand(integrations_speech_nuance.Cmdintegrations_speech_nuance())
	integrations_speechCmd.Short = utils.GenerateCustomDescription(integrations_speechCmd.Short, integrations_speech_dialogflow.Description, integrations_speech_lex.Description, integrations_speech_tts.Description, integrations_speech_nuance.Description, )
	integrations_speechCmd.Long = integrations_speechCmd.Short
}
