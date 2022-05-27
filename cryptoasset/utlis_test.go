package cryptoasset_test

import (
	"encoding/json"
	"os"
	"testing"
)

func readJSONFile(t *testing.T, fName string, data interface{}) {
	dataByte, err := os.ReadFile(fName)
	if err != nil {
		t.Fatalf("couldnt read file. error [%s]", err)
		return
	}
	err = json.Unmarshal(dataByte, data)
	if err != nil {
		t.Fatalf("couldnt unmarshal data to json. error [%s]", err)
	}
}

func writeJSONFile(t *testing.T, fName string, data interface{}, indent ...bool) {
	var dataByte []byte
	var err error
	if len(indent) > 0 && indent[0] {
		dataByte, err = json.MarshalIndent(data, "", "\t")
	} else {
		dataByte, err = json.Marshal(data)
	}

	if err != nil {
		t.Fatalf("couldnt marshal data to json. error [%s]", err)
		return
	}
	err = os.WriteFile(fName, dataByte, 0600)
	if err != nil {
		t.Fatalf("couldnt create file [%s] error [%s]", fName, err)
	}
}
