package core

import (
	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
)

type ContextV2 struct {
	echo.Context
}

func ToContextV2(ctx *echo.Context) *ContextV2 {
	return &ContextV2{*ctx}
}

func (c *ContextV2) GetAuthToken() *auth.Token {
	token := &auth.Token{}
	t := c.Get("authToken")
	token, ok := t.(*auth.Token)
	if !ok {
		return nil
	}
	return token
}
