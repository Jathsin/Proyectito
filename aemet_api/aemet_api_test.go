package aemet_api

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	uv, err := get_uv_prediction(0, "11012")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("Tal rayos UV: %d", uv)
}

func Test2(t *testing.T) {

	uv_array, err := get_uv_local_prediction(2106004)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(uv_array)
}
