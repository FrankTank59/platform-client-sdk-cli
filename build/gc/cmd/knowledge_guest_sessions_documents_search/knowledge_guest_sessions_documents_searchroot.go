package knowledge_guest_sessions_documents_search

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_guest_sessions_documents_search_suggestions"
)

func init() {
	knowledge_guest_sessions_documents_searchCmd.AddCommand(knowledge_guest_sessions_documents_search_suggestions.Cmdknowledge_guest_sessions_documents_search_suggestions())
	knowledge_guest_sessions_documents_searchCmd.Short = utils.GenerateCustomDescription(knowledge_guest_sessions_documents_searchCmd.Short, knowledge_guest_sessions_documents_search_suggestions.Description, )
	knowledge_guest_sessions_documents_searchCmd.Long = knowledge_guest_sessions_documents_searchCmd.Short
}
