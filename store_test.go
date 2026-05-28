package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	pathName := CASPathTransformFunc(key)
	fmt.Println(pathName)
	assert.Equal(t, pathName, "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff")
}
func TestStore(t *testing.T) {
	opts := StoreOpts{
		CASPathTransformFunc,
	}	
	store := NewStore(opts)
	data := bytes.NewReader([]byte("Some distributed JPEGs bytes"))

	assert.Nil(t, store.writeStream(data, "dsys-key"))



}