package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/protobuf/proto"
	"bytes"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"gin-proto-test/example"
)

// protoc build code =>
// protoc HouseContractListInput.proto HouseContractListOutput.proto --go_out=../

func TestProtoMarshal(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	body := &example.HouseContractListInput{KeyWordName: "搜索"}
	inData, _ := proto.Marshal(body)

	req, _ := http.NewRequest("POST", "/test", bytes.NewReader(inData))
	req.Header.Set("Content-Type", "application/x-protobuf")
	router.ServeHTTP(w, req)

	var output example.HouseContractListOutput
	_ = proto.Unmarshal(w.Body.Bytes(), &output)

	assert.Equal(t, 2, int(output.TotalRow))
}

func TestJsonMarshal(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	body := &example.HouseContractListInput{KeyWordName: "搜索"}
	inData, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/test", bytes.NewReader(inData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var output example.HouseContractListOutput
	_ = json.Unmarshal(w.Body.Bytes(), &output)

	assert.Equal(t, 2, int(output.TotalRow))
}
