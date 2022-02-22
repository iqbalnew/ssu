package main

import (
	"context"
	"encoding/json"
	"net/http"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func CustomHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(runtime.HTTPStatusFromCode(grpc.Code(err)))

	body := &pb.ErrorBodyResponse{
		Error:   true,
		Code:    uint32(runtime.HTTPStatusFromCode(grpc.Code(err))),
		Message: grpc.ErrorDesc(err),
	}

	jErr := json.NewEncoder(w).Encode(body)

	if jErr != nil {
		w.Write([]byte(fallback))
	}
}
