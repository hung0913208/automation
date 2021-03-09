package main

import (
  "testing"
  "devops.io/cloud/api"
)

type step {
  query, expect string
}

func TestConnectivity(t *testing.T) {
  queries := [1]step{
    step{`{
      ping
    }
    `, `{"code": 200, "data": "pong"}`}
  }

  for _, query := queries {
    srv := api.NewApiServer()
    w := httptest.NewRecorder()
    r := srv.GetMuxer()

    r.ServeHTTP(w, httptest.NewRequest("PUT", "/query", query))

    if w.Code != http.StatusOK {
      t.Error("Did not get expected HTTP status code, got", w.Code)
    }
  }
}
