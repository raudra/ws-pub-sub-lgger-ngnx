package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"session-service/src/models"

	"github.com/rs/zerolog/log"
)

func FetchProfile(mobileNo string) (*models.User, error) {

	url := fmt.Sprintf("http://host.docker.internal:9000/api/v1/users/findByMobile?mobileNo=%s", mobileNo)
	resp, err := http.Get(url)

	if err != nil {
		return nil, logErrorAndReturn(err, "Error while fetching profile")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, logErrorAndReturn(err, "Error while parsing profile response")
	}

	r := make(map[string]interface{})
	u := new(models.User)

	if err = json.Unmarshal(body, &r); err != nil {
		return nil, logErrorAndReturn(err, "Error while parsing profile response")
	}

	if r["success"].(bool) {
		data := r["data"].(map[string]interface{})
		user := data["user"].(map[string]interface{})
		u.ProfileId = int(user["id"].(float64))
		u.Name = user["name"].(string)
		u.Number = user["number"].(string)
	} else {
		return nil, errors.New(r["error"].(string))
	}

	return u, nil
}

func logErrorAndReturn(err error, msg string) error {
	log.Err(err).
		Msg(msg)

	return err

}
