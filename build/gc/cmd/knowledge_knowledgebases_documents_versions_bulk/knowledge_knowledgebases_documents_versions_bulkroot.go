package knowledge_knowledgebases_documents_versions_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_knowledgebases_documents_versions_bulk_add"
)

func init() {
	knowledge_knowledgebases_documents_versions_bulkCmd.AddCommand(knowledge_knowledgebases_documents_versions_bulk_add.Cmdknowledge_knowledgebases_documents_versions_bulk_add())
	knowledge_knowledgebases_documents_versions_bulkCmd.Short = utils.GenerateCustomDescription(knowledge_knowledgebases_documents_versions_bulkCmd.Short, knowledge_knowledgebases_documents_versions_bulk_add.Description, )
	knowledge_knowledgebases_documents_versions_bulkCmd.Long = knowledge_knowledgebases_documents_versions_bulkCmd.Short
}
