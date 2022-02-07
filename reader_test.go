package pbm_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/slashformotion/pbm"
)

func TestDecodeConfigErrorNil(t *testing.T) {
	testString := `P4
24 24
`
	r := strings.NewReader(testString)
	_, err := pbm.DecodeConfig(r)
	if err != nil {
		t.Error("Decode Config should not return an error on valid string")
	}
}

func TestDecodeConfigErrorNotNil(t *testing.T) {
	testString := `P2
24 24

`
	r := strings.NewReader(testString)
	_, err := pbm.DecodeConfig(r)
	if err == nil {
		t.Error("Decode Config should  return an error on invalid string")
	}
}
func TestDecodeConfigNegativeDims(t *testing.T) {
	testString := `P4
24 -24

`
	r := strings.NewReader(testString)
	_, err := pbm.DecodeConfig(r)
	if err == nil {
		t.Error("Decode Config should  return an error on invalid dimension")
	}
}

func TestDecode(t *testing.T) {
	testString := `P4
4 1
`
	testBytes := []byte(testString)
	testBytes = append(testBytes, byte(0x0f))

	r := bytes.NewReader(testBytes)
	_, err := pbm.Decode(r)
	if err != nil {
		t.Error("Decode should not return an error on valid bytes")
	}
}

func TestDecodeFiles(t *testing.T) {
	filepaths := []string{"fixtures/sample_1280_853.pbm", "fixtures/sample_1920_1280.pbm", "fixtures/sample_5184_3456.pbm", "fixtures/sample_640_426.pbm"}
	for _, path := range filepaths {
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer func() {
			err := f.Close()
			if err != nil {
				panic(err)
			}
		}()
		img, err := pbm.Decode(f)
		if err != nil {
			t.Errorf("Decode should not return an error on valid file, got=%v (file=%v)", err, path)
		}
		if img == nil {
			t.Error("Decode should not return a nil img")
		}
	}
}
