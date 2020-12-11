package user

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestAddressGenerator(t *testing.T) {
	AddressGenerator(new(gin.Context))
}