package router

// TODO: Continue working on test as test fails

import (
	"bytes"
	"diary_api/model"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {

	var reqBuffer bytes.Buffer

	authInput := model.AuthenticationInput{
		Username: "test",
		Password: "test",
	}

	json.NewEncoder(&reqBuffer).Encode(authInput)

	fmt.Println(string(reqBuffer.Bytes()))

	req, err := http.NewRequest("POST", "/auth/login", &reqBuffer)

	// reqBody, _ := io.ReadAll(req.Body)

	// fmt.Println("Request:", string(reqBody))

	fmt.Println("Error:", err)

	require.NoError(t, err)

	resp := httptest.NewRecorder()

	router := SetupRouter()

	router.ServeHTTP(resp, req)

	// var received interface{}

	// json.NewDecoder(resp.Body).Decode(&received)

	// fmt.Println(received)

	// fmt.Println(w.Body.String())

}
