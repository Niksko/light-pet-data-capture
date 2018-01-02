package test

import (
	"testing"
	"net/http"
	"github.com/golang/mock/gomock"
	"github.com/niksko/light-pet-data-capture/mocks"
	"github.com/niksko/light-pet-data-capture/http-handlers"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"strings"
)

func TestRootHandler_WhenRequestIsCalledWithMethod_AppropriateStatusIsReturnedInResponse(t *testing.T) {
	testTable := []struct {
		httpMethod string
		expectedResponse int
	}{
		{http.MethodGet, http.StatusMethodNotAllowed},
		{http.MethodPost, http.StatusOK},
		{http.MethodPatch, http.StatusMethodNotAllowed},
	}


	for _, tableData := range (testTable) {
		mockCtrl := gomock.NewController(t)

		mockResponseWriter := mock_http.NewMockResponseWriter(mockCtrl)
		mockResponseWriter.EXPECT().WriteHeader(tableData.expectedResponse)
		mockResponseWriter.EXPECT().Header().Times(3).Return(make(map[string][]string))

		dummyRequest := http.Request {
			Method: tableData.httpMethod,
			Body: ioutil.NopCloser(strings.NewReader("foobar")),
		}

		dummyUnmarshaler := func (s string, pb proto.Message) error {
			return nil;
		}

		http_handlers.RootHandler(mockResponseWriter, &dummyRequest, dummyUnmarshaler)

		mockCtrl.Finish()
	}
}
