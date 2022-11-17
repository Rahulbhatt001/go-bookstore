// its main use is to umarshal the data
// came from user as it would be in json , so in order to use by controller we need this func
package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// this func is required to parse the body especially in the create func
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
