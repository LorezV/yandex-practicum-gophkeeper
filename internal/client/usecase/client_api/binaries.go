package clientapi

import (
	"fmt"
	"log"
	"os"

	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/go-resty/resty/v2"
)

const binaryEndpoint = "api/v1/user/binary"

func (api *GophKeeperClientAPI) GetBinaries(accessToken string) (binaries []entity.Binary, err error) {
	if err := api.getEntities(&binaries, accessToken, binaryEndpoint); err != nil {
		return nil, err
	}

	return binaries, nil
}

func (api *GophKeeperClientAPI) AddBinary(accessToken string, binary *entity.Binary, tmpFilePath string) error {
	var responseBinary entity.Binary
	client := resty.New()
	client.SetAuthToken(accessToken)
	file, err := os.Open(tmpFilePath)
	if err != nil {
		return fmt.Errorf("GophKeeperClientAPI - AddBinary - %w ", err)
	}
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParam("name", binary.Name).
		SetFileReader("file", binary.FileName, file).
		SetResult(&responseBinary).
		Post(fmt.Sprintf("%s/%s", api.serverURL, binaryEndpoint))
	if err != nil {
		return fmt.Errorf("GophKeeperClientAPI - AddBinary - %w ", err)
	}

	if err := api.checkResCode(resp); err != nil {
		return errServer
	}
	binary.ID = responseBinary.ID
	metaEndpoint := fmt.Sprintf("%s/%s/meta", binaryEndpoint, binary.ID.String())
	if err := api.addEntity(&binary.Meta, accessToken, metaEndpoint); err != nil {
		return fmt.Errorf("GophKeeperClientAPI - AddBinary - AddMeta %w ", err)
	}

	return nil
}

func (api *GophKeeperClientAPI) DelBinary(accessToken, binaryID string) error {
	return api.delEntity(accessToken, binaryEndpoint, binaryID)
}

func (api *GophKeeperClientAPI) DownloadBinary(accessToken, outpuFilePath string, binary *entity.Binary) error {
	client := resty.New()
	client.SetAuthToken(accessToken)
	resp, err := client.R().
		SetOutput(outpuFilePath).
		Get(fmt.Sprintf("%s/%s/%s", api.serverURL, binaryEndpoint, binary.ID.String()))
	if err != nil {
		log.Println(err)

		return err
	}

	if err := api.checkResCode(resp); err != nil {
		return err
	}

	return nil
}
