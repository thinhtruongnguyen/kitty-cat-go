package aiozaiplatformgosdk

type GetCommitHistoryRequest struct {
	OwnerUsername  string
	RepositoryName string
	Path           string
	Sha            string
	RepoType       string
	Page           int32
	PageSize       int32
}
