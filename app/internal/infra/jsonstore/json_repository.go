package jsonstore

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"sync"

	"github.com/lucasti79/meli-interview/pkg/helpers"
)

type IDGetter[T any] func(entity T) string

type JSONRepository[T any] struct {
	filePath string
	mutex    sync.Mutex
	index    map[string]int64
	getID    IDGetter[T]
}

func NewJSONRepository[T any](fileName string, getID IDGetter[T]) (*JSONRepository[T], error) {
	path := filepath.Join(helpers.ProjectRoot(), fileName)
	repo := &JSONRepository[T]{
		filePath: path,
		index:    make(map[string]int64),
		getID:    getID,
	}
	if err := repo.buildIndex(); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *JSONRepository[T]) buildIndex() error {
	r.index = make(map[string]int64)

	f, err := os.Open(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer f.Close()

	var offset int64 = 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var entity T
		line := scanner.Bytes()
		if err := json.Unmarshal(line, &entity); err != nil {
			return err
		}
		id := r.getID(entity)
		r.index[id] = offset
		offset += int64(len(line) + 1) // +1 para o '\n'
	}
	return scanner.Err()
}

func (r *JSONRepository[T]) FindByID(id string) (T, error) {
	var zero T
	r.mutex.Lock()
	defer r.mutex.Unlock()

	offset, ok := r.index[id]
	if !ok {
		return zero, fmt.Errorf("%s not found", reflect.TypeOf(zero).Name())
	}

	f, err := os.Open(r.filePath)
	if err != nil {
		return zero, err
	}
	defer f.Close()

	if _, err := f.Seek(offset, io.SeekStart); err != nil {
		return zero, err
	}

	reader := bufio.NewReader(f)
	line, err := reader.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return zero, err
	}

	var entity T
	if err := json.Unmarshal(line, &entity); err != nil {
		return zero, err
	}

	return entity, nil
}

func (r *JSONRepository[T]) Save(entity T) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	id := r.getID(entity)
	if _, exists := r.index[id]; exists {
		return fmt.Errorf("%s with ID %s already exists", reflect.TypeOf(entity).Name(), id)
	}

	dir := filepath.Dir(r.filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	f, err := os.OpenFile(r.filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	offset, _ := f.Seek(0, io.SeekEnd)
	enc := json.NewEncoder(f)
	if err := enc.Encode(entity); err != nil {
		return err
	}

	r.index[id] = offset
	return nil
}

func (r *JSONRepository[T]) FindAll(handler func(entity T) error) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	f, err := os.Open(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var entity T
		if err := json.Unmarshal(scanner.Bytes(), &entity); err != nil {
			return err
		}
		if err := handler(entity); err != nil {
			return err
		}
	}

	return scanner.Err()
}
