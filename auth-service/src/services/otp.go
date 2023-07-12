package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
)

const (
	OTP_SERVER_URL = "http://host.docker.internal:9001/api/v1/otp"
)

func ValidateOtp(mobileNo string, otp int) (bool, error) {
	var err error

	url := fmt.Sprintf("%s/%s", OTP_SERVER_URL, "validate")

	reqData := map[string]interface{}{
		"mobileNo": mobileNo,
		"otp":      otp,
	}

	jsonData, _ := json.Marshal(reqData)

	log.Info().
		Str("mobile", mobileNo).
		Int("otp", otp).
		Msg("Validating Otp")

	resp, err := http.Post(url,
		"application/json",
		bytes.NewBuffer(jsonData))

	if err != nil {
		log.Err(err).
			Str("mobile", mobileNo).
			Int("otp", otp).
			Msg("Otp validation request failed")
		return false, err
	}

	defer resp.Body.Close()

	r := make(map[string]interface{})

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Err(err).
			Str("mobile", mobileNo).
			Int("otp", otp).
			Msg("Otp validation response parse error")
		return false, err

	}

	if err = json.Unmarshal(body, &r); err != nil {
		return false, logErrorAndReturn(err, "Error while parsing profile response")
	}

	if !r["success"].(bool) {
		return false, errors.New(r["error"].(string))
	}

	return true, nil

}
