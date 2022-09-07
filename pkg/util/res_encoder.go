package util

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	nhttp "net/http"
	"strings"
)

type CustomHttpResponse struct {
	Code    int32
	Message string
	Data    interface{}
}

func CustomResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if v == nil {
		return nil
	}
	if rd, ok := v.(http.Redirector); ok {
		url, code := rd.Redirect()
		nhttp.Redirect(w, r, url, code)
		return nil
	}
	codec, _ := http.CodecForRequest(r, "Accept")
	_, err := codec.Marshal(v)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", strings.Join([]string{"application", codec.Name()}, "/"))

	res := &CustomHttpResponse{
		Code:    200,
		Message: "Success",
		Data:    v,
	}
	data, err := codec.Marshal(res)
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}
