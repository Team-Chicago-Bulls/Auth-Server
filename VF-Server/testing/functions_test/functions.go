package functions_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

type RequestData struct {
	Correo     string `json:"correo"`
	Contrasena string `json:"contrasena"`
}

func RegisterUser(t *testing.T, email string, password string) {
	godotenv.Load(".env")
	ip := os.Getenv("DB_HOST")
	puerto := os.Getenv("PORT")
	apiUrl := "http://" + ip + ":" + puerto + "/user/register_user"

	requestData := RequestData{
		Correo:     email,
		Contrasena: password,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	jsonResponse := string(body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba un código de estado 200 OK, pero se obtuvo %d", resp.StatusCode, "Con el error de", jsonResponse)
	}

}

func Validation_token(t *testing.T, tok string) {
	godotenv.Load(".env")
	ip := os.Getenv("DB_HOST")
	puerto := os.Getenv("PORT")
	apiUrl := "http://" + ip + ":" + puerto + "/user/log_user_token/" + tok
	print(apiUrl)
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	jsonResponse := string(body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba un código de estado 200 OK, pero se obtuvo %d", resp.StatusCode, "Con el error de", jsonResponse)
	}
}
