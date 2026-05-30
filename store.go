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


func CASPathTransformFunc(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])
	
	blockSize := 5
	sliceLen := len(hashStr) / blockSize
	paths := make([]string, sliceLen)

	for i:=0; i < sliceLen; i++ {
		from, to := i * blockSize, (i*blockSize)+blockSize
		paths[i] = hashStr[from:to]
	}

	return PathKey{
		Pathname: strings.Join(paths, "/"),
		Original: hashStr,
	}
}


type PathTransformFunc func(string) PathKey

type PathKey struct {
	Pathname string
	Original string
}
func (p PathKey) FileName() string {
	return fmt.Sprintf("%s/%s", p.Pathname, p.Original)
}

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
func (s *Store) Read(key string) (io.Reader, error) {
	f, err := s.readStream(key)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	buffer := new(bytes.Buffer)
	_, err = io.Copy(buffer, f)
	return buffer, err

}

func (s *Store) readStream(key string) (io.ReadCloser, error) {
	pathKey := s.PathTransformFunc(key)
	return os.Open(pathKey.FileName())
}


func (s *Store) writeStream(r io.Reader, key string) error {
	pathKey := s.PathTransformFunc(key)
	if err := os.MkdirAll(pathKey.Pathname, os.ModePerm); err != nil {
		return err
	}

	fullPath := pathKey.FileName()
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	
	n, err := io.Copy(file, r)
	if err != nil {
		return err
	}	
	fmt.Printf("write %d bytes to this location : %s", n, fullPath)

	return nil 	
}