package quality_forms_surveys_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_forms_surveys_bulk_contexts"
)

func init() {
	quality_forms_surveys_bulkCmd.AddCommand(quality_forms_surveys_bulk_contexts.Cmdquality_forms_surveys_bulk_contexts())
	quality_forms_surveys_bulkCmd.Short = utils.GenerateCustomDescription(quality_forms_surveys_bulkCmd.Short, quality_forms_surveys_bulk_contexts.Description, )
	quality_forms_surveys_bulkCmd.Long = quality_forms_surveys_bulkCmd.Short
}
