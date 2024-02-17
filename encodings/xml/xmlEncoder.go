package xml

import (
	"encoding/xml"
	"net/http"
)

func WriteXml(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(status)

	return xml.NewEncoder(w).Encode(v)
}
