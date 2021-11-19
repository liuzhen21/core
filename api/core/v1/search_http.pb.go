// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http 0.1.0

package v1

import (
	context "context"
	json "encoding/json"
	go_restful "github.com/emicklei/go-restful"
	http "net/http"
)

import transportHTTP "github.com/tkeel-io/kit/transport/http"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the tkeel package it is being compiled against.
// import package.context.http.go_restful.json.

const _ = transportHTTP.ImportAndUsed

type SearchHTTPServer interface {
	Index(context.Context, *IndexObject) (*IndexResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
}

type SearchHTTPHandler struct {
	srv SearchHTTPServer
}

func newSearchHTTPHandler(s SearchHTTPServer) *SearchHTTPHandler {
	return &SearchHTTPHandler{srv: s}
}

func (h *SearchHTTPHandler) Index(req *go_restful.Request, resp *go_restful.Response) {
	in := IndexObject{}
	if err := transportHTTP.GetBody(req, &in.Obj); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	out, err := h.srv.Index(req.Request.Context(), &in)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *SearchHTTPHandler) Search(req *go_restful.Request, resp *go_restful.Response) {
	in := SearchRequest{}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	out, err := h.srv.Search(req.Request.Context(), &in)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func RegisterSearchHTTPServer(container *go_restful.Container, srv SearchHTTPServer) {
	var ws *go_restful.WebService
	for _, v := range container.RegisteredWebServices() {
		if v.RootPath() == "/v1" {
			ws = v
			break
		}
	}
	if ws == nil {
		ws = new(go_restful.WebService)
		ws.ApiVersion("/v1")
		ws.Path("/v1").Produces(go_restful.MIME_JSON)
		container.Add(ws)
	}

	handler := newSearchHTTPHandler(srv)
	ws.Route(ws.POST("/index").
		To(handler.Index))
	ws.Route(ws.GET("/search/{data}").
		To(handler.Search))
}
