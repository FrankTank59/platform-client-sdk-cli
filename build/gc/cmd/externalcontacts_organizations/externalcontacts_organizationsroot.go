package externalcontacts_organizations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_organizations_trustor"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_organizations_schemas"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_organizations_notes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_organizations_contacts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_organizations_relationships"
)

func init() {
	externalcontacts_organizationsCmd.AddCommand(externalcontacts_organizations_trustor.Cmdexternalcontacts_organizations_trustor())
	externalcontacts_organizationsCmd.AddCommand(externalcontacts_organizations_schemas.Cmdexternalcontacts_organizations_schemas())
	externalcontacts_organizationsCmd.AddCommand(externalcontacts_organizations_notes.Cmdexternalcontacts_organizations_notes())
	externalcontacts_organizationsCmd.AddCommand(externalcontacts_organizations_contacts.Cmdexternalcontacts_organizations_contacts())
	externalcontacts_organizationsCmd.AddCommand(externalcontacts_organizations_relationships.Cmdexternalcontacts_organizations_relationships())
	externalcontacts_organizationsCmd.Short = utils.GenerateCustomDescription(externalcontacts_organizationsCmd.Short, externalcontacts_organizations_trustor.Description, externalcontacts_organizations_schemas.Description, externalcontacts_organizations_notes.Description, externalcontacts_organizations_contacts.Description, externalcontacts_organizations_relationships.Description, )
	externalcontacts_organizationsCmd.Long = externalcontacts_organizationsCmd.Short
}
