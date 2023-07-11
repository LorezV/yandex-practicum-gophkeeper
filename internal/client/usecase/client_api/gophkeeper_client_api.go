package clientapi

type GophKeeperClientAPI struct {
	serverURL string
}

func New(url string) *GophKeeperClientAPI {
	return &GophKeeperClientAPI{
		serverURL: url,
	}
}
