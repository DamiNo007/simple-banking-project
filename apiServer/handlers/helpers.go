package handlers

import (
	"banking/apiServer/wrappers"
	"banking/customErrors"
	"banking/encodings"
	"banking/encodings/json"
	"banking/encodings/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
)

const defaultEncodingFormat = encodings.JSON

func retrieveFormatFromRequestHeaders(r *http.Request) encodings.EncodingFormat {
	format := encodings.EncodingFormat(r.Header.Get("Content-Type"))

	if format == "" {
		format = defaultEncodingFormat
	}

	return format
}

func writeResponse(w http.ResponseWriter, status int, fmt encodings.EncodingFormat, response any) (err error) {
	switch fmt {
	case encodings.JSON:
		err = json.WriteJson(w, status, response)
	case encodings.XML:
		err = xml.WriteXml(w, status, response)
	default:
		err = errors.New("unhandled format")
	}

	return
}

func writeErrorResponse(w http.ResponseWriter, format encodings.EncodingFormat, err *customErrors.AppError) error {
	return writeResponse(w, err.Code, format, err.AsMessage())
}

func defaultWriteErrorFunc(w http.ResponseWriter, err *customErrors.AppError) {
	w.WriteHeader(err.Code)
	fmt.Fprintf(w, err.Message)
}

func handleApiResponse(
	w http.ResponseWriter,
	format encodings.EncodingFormat,
	wrappedResponse *wrappers.ApiResponseWrapper,
	appErr *customErrors.AppError,
) {
	if appErr != nil {
		if err := writeErrorResponse(w, format, appErr); err != nil {
			defaultWriteErrorFunc(w, appErr)
		}
	} else {
		if err := writeResponse(w, wrappedResponse.Code, format, wrappedResponse.Body); err != nil {
			log.Printf("failed to write a response: %v", err.Error())
			if err := writeErrorResponse(
				w,
				format,
				customErrors.NewInternalServerError("ups... something went wrong"),
			); err != nil {
				defaultWriteErrorFunc(w, customErrors.NewInternalServerError("ups... something went wrong"))
			}
		}
	}
}
