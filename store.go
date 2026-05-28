package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)


func CASPathTransformFunc(key string) string {
	hash := sha1.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])
	
	blockSize := 5
	sliceLen := len(hashStr) / blockSize
	paths := make([]string, sliceLen)

	for i:=0; i < sliceLen; i++ {
		from, to := i * blockSize, (i*blockSize)+blockSize
		paths[i] = hashStr[from:to]
	}
	return strings.Join(paths, "/")
}


type PathTransformFunc func(string) string

var DefaultPathTransformFunc = func(key string) string {
	return key
}

type StoreOpts struct {
	PathTransformFunc PathTransformFunc
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}





func (s *Store) writeStream(r io.Reader, key string) error {
	pathName := s.PathTransformFunc(key)
	if err := os.MkdirAll(pathName, os.ModePerm); err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	io.Copy(buf, r)

	hash := sha1.Sum(buf.Bytes())
	fileName := hex.EncodeToString(hash[:])

	fullPath := pathName + "/" + fileName
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	
	n, err := io.Copy(file, buf)
	if err != nil {
		return err
	}	
	fmt.Printf("write %d bytes to this location : %s", n, fullPath)

	return nil 	
}