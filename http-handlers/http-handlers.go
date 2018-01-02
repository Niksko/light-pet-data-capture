package http_handlers


import (
	"github.com/golang/protobuf/proto"
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/niksko/light-pet-data-capture/sensor-data"
)

func RootHandler(response http.ResponseWriter, request *http.Request, unmarshalFunc func(string, proto.Message) error) {
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
			log.Print("Error getting request body: %s", err)
			response.WriteHeader(http.StatusInternalServerError)
		}

		err = unmarshalFunc(string(body), decodedData)

		if (err != nil) {
			log.Print("Error unmarshaling request body: %s", err)
			response.WriteHeader(http.StatusInternalServerError)
		}

		response.WriteHeader(http.StatusOK)
		log.Print("Successfully decoded message, sent 200 OK")
	} else {
		response.WriteHeader(http.StatusMethodNotAllowed)
		log.Print(fmt.Sprintf("Request method was %s, sending 405 Method Not Allowed", request.Method))
	}
}