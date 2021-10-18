package services

import (
	"errors"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/valp0/academy-go-q32021/common"
	"github.com/valp0/academy-go-q32021/repo"
)

var characters = []common.Element{
	{Key: 36, Value: "clefable"},
	{Key: 73, Value: "tentacruel"},
	{Key: 125, Value: "electabuzz"},
	{Key: 142, Value: "aerodactyl"},
	{Key: 150, Value: "mewtwo"},
	{Key: 460, Value: "abomasnow"},
	{Key: 498, Value: "tepig"},
	{Key: 794, Value: "buzzwole"},
}

var tentacruel = []common.Element{
	{Key: 73, Value: "tentacruel"},
}

var mewtwo = []common.Element{
	{Key: 150, Value: "mewtwo"},
}

func TestReadSvc(t *testing.T) {
	testCases := []struct {
		name           string
		path           string
		params         map[string][]string
		expectedLength int
		response       []common.Element
		err            error
	}{
		{
			"works fine, displays whole list",
			"../files/pokemock.csv",
			map[string][]string{},
			8,
			characters,
			nil,
		},
		{
			"works fine, displays tentacruel",
			"../files/pokemock.csv",
			map[string][]string{"id": {"73"}},
			1,
			tentacruel,
			nil,
		},
		{
			"wrong id error",
			"../files/wrongId.csv",
			map[string][]string{},
			0,
			[]common.Element{},
			errors.New("parsing error in line 2: \"Hola\" is not a valid id, please use integers for id field"),
		},
		{
			"wrong path error",
			"../wrongPath/someFile.csv",
			map[string][]string{},
			0,
			[]common.Element{},
			errors.New("open ../wrongPath/someFile.csv: El sistema no puede encontrar la ruta especificada."),
		},
		{
			"unexisting file, creates it",
			"../files/deleteThis.csv",
			map[string][]string{},
			0,
			[]common.Element{},
			nil,
		},
	}

	r := repo.NewLocalRepo()
	rSvc := NewReadSvc(r)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			params := tc.params
			res, err := rSvc.Read(params, tc.path)
			if err == nil && tc.err != nil || err != nil && tc.err == nil {
				t.Fatalf("error is not as expected:\ngot: %v\nwant: %v\n", err, tc.err)
			}

			if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Fatalf("error is not as expected:\ngot: %v\nwant: %v\n", err, tc.err)
			}

			if err == nil && err != tc.err {
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

	err := os.Remove("../files/deleteThis.csv")
	if err != nil {
		log.Print("Couldn't remove ../files/deleteThis.csv")
	}
}
