// Generated by github.com/hypermusk/hypermusk
// DO NOT EDIT

package apihttpimpl

import (
	"time"
	"io"
	"strings"
	"compress/gzip"
	"encoding/json"
	"encoding/base64"
	"github.com/hypermusk/hypermusk/tests/api"
	"github.com/hypermusk/hypermusk/tests/apiimpl"
	"net/http"
	"log"
)

var _ = time.Sunday
var _ = base64.StdEncoding

const StreamHTTPHeaderFieldName = "X-HyperMuskStreamParams"

type CodeError interface {
	Code() string
}

type SerializableError struct {
	Code    string
	Message string
	Reason  error
}

func (s *SerializableError) Error() string {
	return s.Message
}

func NewError(err error) (r error) {
	se := &SerializableError{Message:err.Error()}
	ce, yes := err.(CodeError)
	if yes {
		se.Code = ce.Code()
	}
	se.Reason = err
	r = se
	return
}

func AddToMux(prefix string, mux *http.ServeMux) {
	
	mux.HandleFunc(prefix+"/Service/Authorize.json", Service_Authorize)
	mux.HandleFunc(prefix+"/Service/PermiessionDenied.json", Service_PermiessionDenied)
	mux.HandleFunc(prefix+"/Service/GetReservedKeywordsForObjC.json", Service_GetReservedKeywordsForObjC)
	return
}


var service api.Service = apiimpl.DefaultService

type ServiceData struct {
}


type Service_Authorize_Params struct {
	Params struct {
		Name string
	}
}

type Service_Authorize_Results struct {
	UseMap api.UseMap
	Err error

}

func Service_Authorize(w http.ResponseWriter, r *http.Request) {
	startAt := time.Now()
	log.Printf("Started %s \"%s\" for %s at %s\n", r.Method, r.RequestURI, r.RemoteAddr, startAt.String())
	defer func() {
		spent := time.Now().Sub(startAt) / time.Millisecond
		log.Printf("Completed \"%s\" in %d ms\n\n", r.RequestURI, spent)
	}()

	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	gzipOk := strings.Contains(r.Header.Get("Accept-Encoding"), "gzip")
	if gzipOk {
		w.Header().Set("Content-Encoding", "gzip")
	}

	var body io.Reader

	jsonpCallback := r.URL.Query().Get("callback")
	if  jsonpCallback != "" {
		body = strings.NewReader(r.URL.Query().Get("data"))
		io.WriteString(w, jsonpCallback + "(")
		defer io.WriteString(w, ")")
	} else {
		if r.Body == nil {
			panic("no body")
		}
		body = r.Body
	}

	defer r.Body.Close()

	var p Service_Authorize_Params

	dec := json.NewDecoder(body)
	err := dec.Decode(&p)
	var result Service_Authorize_Results
	var mayGzipWriter io.Writer = w
	if gzipOk {
		gzipWriter := gzip.NewWriter(w)
		defer gzipWriter.Close()
		mayGzipWriter = gzipWriter
	}
	enc := json.NewEncoder(mayGzipWriter)
	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}

	s := service

	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}
	result.UseMap, result.Err = s.Authorize(p.Params.Name)
	if result.Err != nil {
		result.Err = NewError(result.Err)
	}
	err = enc.Encode(result)
	if err != nil {
		panic(err)
	}
	return
}

type Service_PermiessionDenied_Params struct {
	Params struct {
	}
}

type Service_PermiessionDenied_Results struct {
	Err error

}

func Service_PermiessionDenied(w http.ResponseWriter, r *http.Request) {
	startAt := time.Now()
	log.Printf("Started %s \"%s\" for %s at %s\n", r.Method, r.RequestURI, r.RemoteAddr, startAt.String())
	defer func() {
		spent := time.Now().Sub(startAt) / time.Millisecond
		log.Printf("Completed \"%s\" in %d ms\n\n", r.RequestURI, spent)
	}()

	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	gzipOk := strings.Contains(r.Header.Get("Accept-Encoding"), "gzip")
	if gzipOk {
		w.Header().Set("Content-Encoding", "gzip")
	}

	var body io.Reader

	jsonpCallback := r.URL.Query().Get("callback")
	if  jsonpCallback != "" {
		body = strings.NewReader(r.URL.Query().Get("data"))
		io.WriteString(w, jsonpCallback + "(")
		defer io.WriteString(w, ")")
	} else {
		if r.Body == nil {
			panic("no body")
		}
		body = r.Body
	}

	defer r.Body.Close()

	var p Service_PermiessionDenied_Params

	dec := json.NewDecoder(body)
	err := dec.Decode(&p)
	var result Service_PermiessionDenied_Results
	var mayGzipWriter io.Writer = w
	if gzipOk {
		gzipWriter := gzip.NewWriter(w)
		defer gzipWriter.Close()
		mayGzipWriter = gzipWriter
	}
	enc := json.NewEncoder(mayGzipWriter)
	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}

	s := service

	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}
	result.Err = s.PermiessionDenied()
	if result.Err != nil {
		result.Err = NewError(result.Err)
	}
	err = enc.Encode(result)
	if err != nil {
		panic(err)
	}
	return
}

type Service_GetReservedKeywordsForObjC_Params struct {
	Params struct {
	}
}

type Service_GetReservedKeywordsForObjC_Results struct {
	R api.ReservedKeywordsForObjC
	Err error

}

func Service_GetReservedKeywordsForObjC(w http.ResponseWriter, r *http.Request) {
	startAt := time.Now()
	log.Printf("Started %s \"%s\" for %s at %s\n", r.Method, r.RequestURI, r.RemoteAddr, startAt.String())
	defer func() {
		spent := time.Now().Sub(startAt) / time.Millisecond
		log.Printf("Completed \"%s\" in %d ms\n\n", r.RequestURI, spent)
	}()

	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	gzipOk := strings.Contains(r.Header.Get("Accept-Encoding"), "gzip")
	if gzipOk {
		w.Header().Set("Content-Encoding", "gzip")
	}

	var body io.Reader

	jsonpCallback := r.URL.Query().Get("callback")
	if  jsonpCallback != "" {
		body = strings.NewReader(r.URL.Query().Get("data"))
		io.WriteString(w, jsonpCallback + "(")
		defer io.WriteString(w, ")")
	} else {
		if r.Body == nil {
			panic("no body")
		}
		body = r.Body
	}

	defer r.Body.Close()

	var p Service_GetReservedKeywordsForObjC_Params

	dec := json.NewDecoder(body)
	err := dec.Decode(&p)
	var result Service_GetReservedKeywordsForObjC_Results
	var mayGzipWriter io.Writer = w
	if gzipOk {
		gzipWriter := gzip.NewWriter(w)
		defer gzipWriter.Close()
		mayGzipWriter = gzipWriter
	}
	enc := json.NewEncoder(mayGzipWriter)
	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}

	s := service

	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}
	result.R, result.Err = s.GetReservedKeywordsForObjC()
	if result.Err != nil {
		result.Err = NewError(result.Err)
	}
	err = enc.Encode(result)
	if err != nil {
		panic(err)
	}
	return
}




