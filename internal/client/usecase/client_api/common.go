package clientapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LorezV/gophkeeper/internal/utils/errs"
	"github.com/fatih/color"
	"github.com/go-resty/resty/v2"
	"golang.org/x/exp/slices"
)

func (api *GophKeeperClientAPI) addEntity(entity interface{}, accessToken, endpoint string) error {
	client := resty.New()
	client.SetAuthToken(accessToken)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(entity).
		SetResult(entity).
		Post(fmt.Sprintf("%s/%s", api.serverURL, endpoint))
	if err != nil {
		log.Fatalf("GophKeeperClientAPI - client.R - %v ", err)
	}
	if err := api.checkResCode(resp); err != nil {
		return errServer
	}

	return nil
}

func (api *GophKeeperClientAPI) getEntities(entity interface{}, accessToken, endpoint string) error {
	client := resty.New()
	client.SetAuthToken(accessToken)
	resp, err := client.R().
		SetResult(entity).
		Get(fmt.Sprintf("%s/%s", api.serverURL, endpoint))
	if err != nil {
		log.Println(err)

		return err
	}

	if err := api.checkResCode(resp); err != nil {
		return err
	}

	return nil
}

func (api *GophKeeperClientAPI) checkResCode(resp *resty.Response) error {
	badCodes := []int{http.StatusBadRequest, http.StatusInternalServerError, http.StatusUnauthorized}
	if slices.Contains(badCodes, resp.StatusCode()) {
		errMessage := errs.ParseServerError(resp.Body())
		color.Red("Server error: %s", errMessage)

		return errServer
	}

	return nil
}

func (api *GophKeeperClientAPI) delEntity(accessToken, endpoint, id string) error {
	client := resty.New()
	client.SetAuthToken(accessToken)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Delete(fmt.Sprintf("%s/%s/%s", api.serverURL, endpoint, id))
	if err != nil {
		log.Fatalf("GophKeeperClientAPI - client.R - %v ", err)
	}
	if err := api.checkResCode(resp); err != nil {
		return errServer
	}

	return nil
}
