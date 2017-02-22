package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

	"github.com/artefactual-labs/archivematica-workflow/endpoints"
	httptransport "github.com/artefactual-labs/archivematica-workflow/http"
	"github.com/artefactual-labs/archivematica-workflow/service"
)

func TestWiring(t *testing.T) {
	svc := service.New(log.NewNopLogger())
	eps := endpoints.New(svc, log.NewNopLogger())
	mux := httptransport.NewHandler(context.Background(), eps, log.NewNopLogger())
	srv := httptest.NewServer(mux)
	defer srv.Close()

	for _, testcase := range []struct {
		method string
		url    string
		want   int
	}{
		{"GET", srv.URL + "/workflows/default", 200},
		{"GET", srv.URL + "/workflows/foobar", 404},
	} {
		req, _ := http.NewRequest(testcase.method, testcase.url, strings.NewReader(``))
		resp, _ := http.DefaultClient.Do(req)

		o, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(o))

		if want, have := testcase.want, resp.StatusCode; want != have {
			t.Errorf("%s %s: want=%d have=%d", testcase.method, testcase.url, want, have)
		}
	}
}
