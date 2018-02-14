package untappd

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckinParse(t *testing.T) {
	payload, err := ioutil.ReadFile("examples/timeline.json")
	assert.Nil(t, err)
	var f feedResp
	err = json.Unmarshal(payload, &f)
	assert.Nil(t, err)
}

func TestToastParse(t *testing.T) {
	payload, err := ioutil.ReadFile("examples/toast.json")
	assert.Nil(t, err)
	var f toastResponse
	err = json.Unmarshal(payload, &f)
	assert.Nil(t, err)
}
