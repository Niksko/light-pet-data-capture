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
	"io"
	"errors"
)

// Type for testing that allows us to test error scenarios during reading response bodies
type ErrorReader struct {
	io.Reader
}

func (ErrorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("Test")
}

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
			Body: ioutil.NopCloser(strings.NewReader("061af7")),
		}

		dummyUnmarshaler := func (b []byte, pb proto.Message) error {
			return nil;
		}

		http_handlers.RootHandler(mockResponseWriter, &dummyRequest, dummyUnmarshaler)

		mockCtrl.Finish()
	}
}

func TestRootHandler_RequestBodyCannotBeRead_ReturnsInternalServerError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	mockResponseWriter := mock_http.NewMockResponseWriter(mockCtrl)
	mockResponseWriter.EXPECT().WriteHeader(http.StatusInternalServerError)
	mockResponseWriter.EXPECT().Header().Times(3).Return(make(map[string][]string))

	dummyRequest := http.Request {
		Method: http.MethodPost,
		Body: ioutil.NopCloser( ErrorReader { strings.NewReader("Test") } ),
	}

	dummyUnmarshaler := func (b []byte, pb proto.Message) error {
		return nil;
	}

	http_handlers.RootHandler(mockResponseWriter, &dummyRequest, dummyUnmarshaler)

	mockCtrl.Finish()
}

func TestRootHandler_UnmarshallingBodyReturnsError_ReturnsInternalServerError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	mockResponseWriter := mock_http.NewMockResponseWriter(mockCtrl)
	mockResponseWriter.EXPECT().WriteHeader(http.StatusInternalServerError)
	mockResponseWriter.EXPECT().Header().Times(3).Return(make(map[string][]string))

	dummyRequest := http.Request {
		Method: http.MethodPost,
		Body: ioutil.NopCloser( strings.NewReader("Test") ),
	}

	dummyUnmarshaler := func (b []byte , pb proto.Message) error {
		return errors.New("Error message");
	}

	http_handlers.RootHandler(mockResponseWriter, &dummyRequest, dummyUnmarshaler)

	mockCtrl.Finish()
}
