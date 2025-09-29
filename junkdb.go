package junkdb

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type JunkDB struct {
	uri   string
	store io.ReadWriteCloser
}

var _ io.Closer = (*JunkDB)(nil)

func Connect(uri string) (*JunkDB, error) {
	file, err := os.OpenFile(uri, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	return &JunkDB{uri: uri, store: file}, nil
}

func (db *JunkDB) Insert(key string, value string) error {
	_, err := fmt.Fprintln(db.store, key+","+value)
	return err
}

func (db *JunkDB) Get(key string) (string, error) {
	sc := bufio.NewScanner(db.store)
	value := ""
	for sc.Scan() {
		line := sc.Text()
		if strings.HasPrefix(line, key) {
			value = strings.TrimSpace(line[len(key)+1:])
		}
	}
	if err := sc.Err(); err != nil {
		return "", err
	}
	return value, nil
}

func (db *JunkDB) Delete(key string) error {
	panic("implement me")
}

func (db *JunkDB) Close() error {
	return db.store.Close()
}
