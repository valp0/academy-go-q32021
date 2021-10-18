package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/valp0/academy-go-q32021/common"
	"github.com/valp0/academy-go-q32021/repo"
)

var charactersMap = map[string]common.Element{
	"36":  {Key: 36, Value: "clefable"},
	"73":  {Key: 73, Value: "tentacruel"},
	"125": {Key: 125, Value: "electabuzz"},
	"142": {Key: 142, Value: "aerodactyl"},
	"150": {Key: 150, Value: "mewtwo"},
	"460": {Key: 460, Value: "abomasnow"},
	"498": {Key: 498, Value: "tepig"},
	"794": {Key: 794, Value: "buzzwole"},
}

func TestFetchSvc(t *testing.T) {
	testCases := []struct {
		name           string
		path           string
		params         map[string][]string
		expectedLength int
		response       []common.Element
		err            error
	}{
		{
			"unexisting file, tentacruel",
			"../files/deleteThisToo.csv",
			map[string][]string{"id": {"73"}},
			1,
			tentacruel,
			nil,
		},
		{
			"already added",
			"../files/deleteThisToo.csv",
			map[string][]string{"id": {"73"}},
			0,
			[]common.Element{},
			errors.New("pokémon with id 73 was already stored in csv"),
		},
		{
			"add random pokémon (mocked mewtwo)",
			"../files/deleteThisToo.csv",
			map[string][]string{},
			1,
			mewtwo,
			nil,
		},
	}

	r := repo.NewApiRepo()
	fSvc := NewFetchSvc(r)
	common.RandInt = func(int) int { return 150 } // Mocking RandInt function
	repo.CallApi = mockedGet                      // Mocking get HTTP request

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			params := tc.params
			res, err := fSvc.Fetch(params, tc.path)
			if err == nil && tc.err != nil || err != nil && tc.err == nil {
				t.Fatalf("error is not as expected:\ngot: %v\nwant: %v\n", err, tc.err)
			}

			if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Fatalf("error is not as expected:\ngot: %v\nwant: %v\n", err, tc.err)
			}

			if len(res) != tc.expectedLength {
				t.Fatalf("response has different length:\ngot: %v\nwant: %v\n", len(res), tc.expectedLength)
			}

			if !reflect.DeepEqual(res, tc.response) {
				t.Fatalf("response is not as expected:\ngot: %v\nwant: %v\n", res, tc.response)
			}
		})
	}

	err := os.Remove("../files/deleteThisToo.csv")
	if err != nil {
		log.Print("Couldn't remove ../files/deleteThisToo.csv")
	}
}

var mockedGet = func(url string) (*http.Response, error) {
	id := strings.Split(url, "/")[6]
	pokemon := charactersMap[id]
	jRes, _ := json.Marshal(pokemon)
	r := ioutil.NopCloser(bytes.NewReader(jRes))

	return &http.Response{
		StatusCode: 200,
		Body:       r,
	}, nil
}
