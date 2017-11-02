package controllers

import (
	"testing"
	"fmt"
	"github.com/satori/go.uuid"
)

func TestTemp(t *testing.T) {

	fmt.Println(uuid.NewV4().String(),len(uuid.NewV4().String()))
}

