package verify

var (
	ArchiveVerification = Rules{"Path": {NotEmpty()}}
	AppVerification     = Rules{}
	AuthVerification    = Rules{}
)
