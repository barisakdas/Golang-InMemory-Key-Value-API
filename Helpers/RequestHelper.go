package Helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	. "ys/Log"
	. "ys/Models"
)

func ParseRequestToModel(r *http.Request) (DataModel, error) {
	var requestModel DataModel
	bodyText, err := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &requestModel)

	if err != nil {
		CreateLogJson("Error", "ParseRequestToModel", "An error occurred when converting the data from the request json to a golang struct.", err.Error())
	}
	return requestModel, nil
}
