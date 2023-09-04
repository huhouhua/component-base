package core

import (
	"github.com/huhouhua/errors"
	"testing"
)

func Test_WriteResponseDetail(t *testing.T) {
	//WriteResponseDetail(nil, errors.New(""), nil)
	WriteResponseDetail(nil, errors.WithMessage(errors.New(""), ""), nil)
	//WriteResponseDetail(nil, errors.WithCode(400, "400"), nil)
}
