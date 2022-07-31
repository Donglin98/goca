package httpserver

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var sever *http.Server
var running bool

func init() {

}

func Run() {
	if running {
		return
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/",rootHandler)
	mux.HandleFunc("/cer-template",getCsrTemplateHandler)
	mux.HandleFunc("/csr",sinCsrHandler)
	server := http.Server{
		Addr:              ":8001",
		Handler:           mux,
	}
	running = true
	fmt.Printf("server listening at %v,http",server.Addr)
	if server.ListenAndServe() !=nil{
		running = false
		log.Printf("can't start http server at %v",server.Addr)
	}
	running = false
}

func sinCsrHandler(writer http.ResponseWriter, request *http.Request) {
	
}

func getCsrTemplateHandler(w http.ResponseWriter, r *http.Request) {
	//if r.Method != "GET" {
	//	w.WriteHeader(http.StatusMethodNotAllowed)
	//	return
	//}
	//
	//csr := ca.CertificateSigningRequest{
	//	SubjectCountry:            []string{"China"},
	//	SubjectOrganization:       []string{"Qinghua"},
	//	SubjectOrganizationalUnit: []string{"ComputerScience"},
	//	SubjectProvince:           []string{"Beijing"},
	//	SubjectLocality:           []string{"北京"},
	//
	//	SubjectCommonName: "tsinghua.edu.cn",
	//	EmailAddresses:    []string{"ex@example.com"},
	//	DNSNames:          []string{"localhost"},
	//}
	//
	//csrBytes, err := json.Marshal(csr)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//
	//w.WriteHeader(http.StatusOK)
	//w.Header().Set("content-type", "application/json")
	//w.Write(csrBytes)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	body ,err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("error when reading body"))
		return
	}
	w.Write(body)
}
