package util

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	nhttp "net/http"
	"strconv"
	"strings"
)

func CustomErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	se := errors.FromError(err)
	codec, _ := http.CodecForRequest(r, "Accept")
	_, err = codec.Marshal(se)
	if err != nil {
		w.WriteHeader(nhttp.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", strings.Join([]string{"application", codec.Name()}, "/"))

	b := se.Code > 0 && se.Code <= 600
	if b {
		w.WriteHeader(int(se.Code))
	}

	code, _ := strconv.ParseInt(se.Reason, 10, 32)
	res := &CustomHttpResponse{
		Code:    int32(code),
		Message: se.Message,
		Data:    se.Metadata,
	}
	data, err := codec.Marshal(res)
	w.Write(data)

}
