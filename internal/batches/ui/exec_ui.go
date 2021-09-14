package ui

import (
	"github.com/sourcegraph/src-cli/internal/batches"
	"github.com/sourcegraph/src-cli/internal/batches/executor"
	"github.com/sourcegraph/src-cli/internal/batches/graphql"
	"github.com/sourcegraph/src-cli/internal/batches/workspace"
)

type ExecUI interface {
	ParsingBatchSpec()
	ParsingBatchSpecSuccess()
	ParsingBatchSpecFailure(error)

	ResolvingNamespace()
	ResolvingNamespaceSuccess(namespace string)

	PreparingContainerImages()
	PreparingContainerImagesProgress(done, total int)
	PreparingContainerImagesSuccess()

	DeterminingWorkspaceCreatorType()
	DeterminingWorkspaceCreatorTypeSuccess(wt workspace.CreatorType)

	ResolvingRepositories()
	ResolvingRepositoriesDone(repos []*graphql.Repository, unsupported batches.UnsupportedRepoSet, ignored batches.IgnoredRepoSet)

	DeterminingWorkspaces()
	DeterminingWorkspacesSuccess(num int)

	CheckingCache()
	CheckingCacheSuccess(cachedSpecsFound int, tasksToExecute int)

	ExecutingTasks(verbose bool, parallelism int) executor.TaskExecutionUI
	ExecutingTasksSkippingErrors(err error)

	LogFilesKept(files []string)

	NoChangesetSpecs()
	UploadingChangesetSpecs(num int)
	UploadingChangesetSpecsProgress(done, total int)
	UploadingChangesetSpecsSuccess(ids []graphql.ChangesetSpecID)

	CreatingBatchSpec()
	CreatingBatchSpecSuccess()
	CreatingBatchSpecError(err error) error

	PreviewBatchSpec(previewURL string)

	ApplyingBatchSpec()
	ApplyingBatchSpecSuccess(batchChangeURL string)

	ExecutionError(error)

	WriteCache(key string, data []byte)
}
