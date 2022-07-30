package doc

import (
	"fmt"
	"time"

	"github.com/admpub/go-zinc/doc/schemas"
	resty "github.com/admpub/resty/v2"
)

type ZincDocSDK interface {
	InsertDocument(index string, doc interface{}) error
	DeleteDocument(index string, id string) error
	UpdateDocument(index string, id string, doc interface{}) error
	SearchDocuments(index string, req *schemas.SearchRequest) (*schemas.SearchResponse, error)
}

type zincDocImpl struct {
	client *resty.Client
	host   string
}

func NewSDK(host, user, pwd string, timeout ...time.Duration) (ZincDocSDK, error) {
	client := resty.New()
	client.SetBasicAuth(user, pwd)
	client.SetBaseURL(host)
	client.SetDisableWarn(true)
	if len(timeout) > 0 {
		client.SetTimeout(timeout[0])
	}
	return &zincDocImpl{
		client: client,
		host:   host,
	}, nil
}

func (sdk *zincDocImpl) InsertDocument(index string, doc interface{}) error {
	resp, err := sdk.client.R().SetBody(doc).Put(fmt.Sprintf("/api/%s/document", index))
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("code=%d, msg=%s", resp.StatusCode(), string(resp.Body()))
	}
	return nil
}

func (sdk *zincDocImpl) DeleteDocument(index string, id string) error {
	resp, err := sdk.client.R().Delete(fmt.Sprintf("/api/%s/_doc/%s", index, id))
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("code=%d, msg=%s", resp.StatusCode(), string(resp.Body()))
	}
	return nil
}

func (sdk *zincDocImpl) UpdateDocument(index string, id string, doc interface{}) error {
	resp, err := sdk.client.R().SetBody(doc).Put(fmt.Sprintf("/api/%s/_doc/%s", index, id))
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("code=%d, msg=%s", resp.StatusCode(), string(resp.Body()))
	}
	return nil
}

func (sdk *zincDocImpl) SearchDocuments(index string, req *schemas.SearchRequest) (*schemas.SearchResponse, error) {
	out := &schemas.SearchResponse{}
	resp, err := sdk.client.R().SetBody(req).SetResult(out).Post(fmt.Sprintf("/api/%s/_search", index))
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("code=%d, msg=%s", resp.StatusCode(), string(resp.Body()))
	}
	return out, nil
}
