package http_handlers


import (
	"github.com/golang/protobuf/proto"
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/niksko/light-pet-data-capture/sensor-data"
	"encoding/hex"
)

func RootHandler(response http.ResponseWriter, request *http.Request, unmarshalFunc func([]byte, proto.Message) error) {
	response.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
	response.Header().Add("X-Frame-Options", "DENY")
	response.Header().Add("X-Content-Type-Options", "nosniff")

	log.Print(fmt.Sprintf("Handled request at / from %s", request.RemoteAddr))

	if (request.Method == http.MethodPost) {
		log.Print("Request method was POST")

		// Decode sent data
		decodedData := &sensor_data.SensorData {}
		body, err := ioutil.ReadAll(request.Body)

		if (err != nil) {
			log.Print("Error getting request body: ", err)
			response.WriteHeader(http.StatusInternalServerError)
			return
		}

		bodyDataAsBytes, err := hex.DecodeString(string(body))

		if (err != nil) {
			log.Print("Error decoding hex string to bytes: ", err)
			response.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = unmarshalFunc(bodyDataAsBytes, decodedData)

		if (err != nil) {
			log.Print("Error unmarshaling request body: ", err)
			response.WriteHeader(http.StatusInternalServerError)
			return
		}

		response.WriteHeader(http.StatusOK)
		log.Print("Successfully decoded message, sent 200 OK")
		log.Print(decodedData)
	} else {
		response.WriteHeader(http.StatusMethodNotAllowed)
		log.Print(fmt.Sprintf("Request method was %s, sending 405 Method Not Allowed", request.Method))
	}
}