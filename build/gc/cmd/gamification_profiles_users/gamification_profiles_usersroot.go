package gamification_profiles_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_profiles_users_me"
)

func init() {
	gamification_profiles_usersCmd.AddCommand(gamification_profiles_users_me.Cmdgamification_profiles_users_me())
	gamification_profiles_usersCmd.Short = utils.GenerateCustomDescription(gamification_profiles_usersCmd.Short, gamification_profiles_users_me.Description, )
	gamification_profiles_usersCmd.Long = gamification_profiles_usersCmd.Short
}
