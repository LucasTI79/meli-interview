package jsonstore

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lucasti79/meli-interview/pkg/apperrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestEntity struct {
	ID    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Group string `json:"group,omitempty"`
}

func getTestEntityID(e TestEntity) string { return e.ID }

func writeJSONL(t *testing.T, path string, entities []TestEntity) {
	t.Helper()
	var lines []string
	for _, e := range entities {
		b, err := json.Marshal(e)
		require.NoError(t, err)
		lines = append(lines, string(b))
	}
	content := strings.Join(lines, "\n")
	if len(lines) > 0 {
		content += "\n"
	}
	require.NoError(t, os.WriteFile(path, []byte(content), 0o600))
}

func countFileLines(t *testing.T, path string) int {
	t.Helper()
	f, err := os.Open(path)
	require.NoError(t, err)
	defer f.Close()

	sc := bufio.NewScanner(f)
	n := 0
	for sc.Scan() {
		n++
	}
	require.NoError(t, sc.Err())
	return n
}

func TestNewJSONRepository_BuildsIndexAndFindByID(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "entities.jsonl")

	entities := []TestEntity{
		{ID: "1", Name: "Alice"},
		{ID: "2", Name: "Bob"},
		{ID: "3", Name: "Carol"},
	}
	writeJSONL(t, fp, entities)

	repo, err := NewJSONRepository[TestEntity](fp, getTestEntityID)
	require.NoError(t, err)

	got, err := repo.FindByID("2")
	require.NoError(t, err)
	require.Equal(t, "2", got.ID)
	require.Equal(t, "Bob", got.Name)
}

func TestSave_AppendsAndUpdatesIndex(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "entities.jsonl")

	repo, err := NewJSONRepository[TestEntity](fp, getTestEntityID)
	require.NoError(t, err)

	e := TestEntity{ID: "10", Name: "Delta"}
	require.NoError(t, repo.Save(e))

	got, err := repo.FindByID("10")
	require.NoError(t, err)
	require.Equal(t, e, got)

	count := 0
	err = repo.FindAll(func(entity TestEntity) error {
		count++
		require.Equal(t, e, entity)
		return nil
	})
	require.NoError(t, err)
	require.Equal(t, 1, count)
}

func TestFindAllWherePaginated_ReturnsExpectedPageAndTotal(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "entities.jsonl")

	ents := []TestEntity{
		{ID: "1", Group: "A"},
		{ID: "2", Group: "B"},
		{ID: "3", Group: "A"},
		{ID: "4", Group: "A"},
		{ID: "5", Group: "B"},
		{ID: "6", Group: "A"},
		{ID: "7", Group: "A"},
	}
	writeJSONL(t, fp, ents)

	repo, err := NewJSONRepository[TestEntity](fp, getTestEntityID)
	require.NoError(t, err)

	var collected []string
	total, err := repo.FindAllWherePaginated(
		func(e TestEntity) bool { return e.Group == "A" },
		2, 2,
		func(e TestEntity) error {
			collected = append(collected, e.ID)
			return nil
		},
	)
	require.NoError(t, err)
	require.Equal(t, 5, total)

	require.Equal(t, []string{"4", "6"}, collected)
}

func TestFindByID_WhenNotFound_ReturnsErrResourceNotExists(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "entities.jsonl")

	repo, err := NewJSONRepository[TestEntity](fp, getTestEntityID)
	require.NoError(t, err)

	_, err = repo.FindByID("missing")
	require.ErrorIs(t, err, apperrors.ErrResourceNotExists)
}

func TestSave_DuplicateID_RejectsWithoutWriting(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "entities.jsonl")

	orig := TestEntity{ID: "dup", Name: "Original"}
	writeJSONL(t, fp, []TestEntity{orig})

	repo, err := NewJSONRepository[TestEntity](fp, getTestEntityID)
	require.NoError(t, err)

	beforeLines := countFileLines(t, fp)
	beforeIndexSize := len(repo.index)

	dup := TestEntity{ID: "dup", Name: "NewName"}
	err = repo.Save(dup)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "already exists")

	afterLines := countFileLines(t, fp)
	afterIndexSize := len(repo.index)

	require.Equal(t, beforeLines, afterLines, "file should not be appended on duplicate")
	require.Equal(t, beforeIndexSize, afterIndexSize, "index should not change on duplicate")

	got, err := repo.FindByID("dup")
	require.NoError(t, err)
	require.Equal(t, orig, got, "original entity must remain intact")
}

func TestFindAll_WhenMalformedJSON_ReturnsInvalidDataFormat(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "entities.jsonl")

	content := `{"id":"ok","name":"Valid"}
not-a-json
`
	require.NoError(t, os.WriteFile(fp, []byte(content), 0o600))

	repo, err := NewJSONRepository[TestEntity](fp, getTestEntityID)
	require.NoError(t, err)

	visited := 0
	err = repo.FindAll(func(e TestEntity) error {
		visited++
		return nil
	})

	require.Error(t, err)
	require.ErrorIs(t, err, apperrors.ErrInvalidDataFormat)
	assert.Contains(t, err.Error(), "line 2")
	require.Equal(t, 1, visited, "handler should be called only for the first valid line")
}

func TestFindAllWhere_CallsHandlerForMatchingEntitiesInOrder(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "entities.jsonl")

	ents := []TestEntity{
		{ID: "1", Group: "A"},
		{ID: "2", Group: "B"},
		{ID: "3", Group: "A"},
		{ID: "4", Group: "A"},
		{ID: "5", Group: "B"},
	}
	writeJSONL(t, fp, ents)

	repo, err := NewJSONRepository[TestEntity](fp, getTestEntityID)
	require.NoError(t, err)

	var visited []string
	err = repo.FindAllWhere(
		func(e TestEntity) bool { return e.Group == "A" },
		func(e TestEntity) error {
			visited = append(visited, e.ID)
			return nil
		},
	)
	require.NoError(t, err)
	require.Equal(t, []string{"1", "3", "4"}, visited, "should preserve file order for matching entities")
}

func TestFindAllWhere_ReturnsNilOnEmptyFileAndDoesNotCallHandler(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "empty.jsonl")
	require.NoError(t, os.WriteFile(fp, []byte(""), 0o600))

	repo, err := NewJSONRepository[TestEntity](fp, getTestEntityID)
	require.NoError(t, err)

	count := 0
	err = repo.FindAllWhere(
		func(e TestEntity) bool { return true },
		func(e TestEntity) error {
			count++
			return nil
		},
	)
	require.NoError(t, err)
	require.Equal(t, 0, count)
}

func TestFindAllWhere_FileNotExistReturnsNilWithoutHandlerCalls(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "does_not_exist.jsonl")

	repo, err := NewJSONRepository[TestEntity](fp, getTestEntityID)
	require.NoError(t, err)

	count := 0
	err = repo.FindAllWhere(
		func(e TestEntity) bool { return true },
		func(e TestEntity) error {
			count++
			return nil
		},
	)
	require.NoError(t, err)
	require.Equal(t, 0, count)
}

func TestFindAllWhere_InvalidJSONReturnsWrappedFormatErrorWithLineNo(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "malformed.jsonl")

	content := `{"id":"1","group":"A"}
{"id":"2","group":"B"}
not-a-json
{"id":"3","group":"A"}
`
	require.NoError(t, os.WriteFile(fp, []byte(content), 0o600))

	repo, err := NewJSONRepository[TestEntity](fp, getTestEntityID)
	require.NoError(t, err)

	visited := 0
	err = repo.FindAllWhere(
		func(e TestEntity) bool { return e.Group == "A" || e.Group == "B" },
		func(e TestEntity) error {
			visited++
			return nil
		},
	)
	require.Error(t, err)
	require.ErrorIs(t, err, apperrors.ErrInvalidDataFormat)
	assert.Contains(t, err.Error(), fp)
	assert.Contains(t, err.Error(), "line 3")
	require.Equal(t, 2, visited, "should process only valid lines before malformed one")
}

func TestFindAllWhere_HandlerErrorShortCircuitsAndPropagates(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "handler_error.jsonl")

	ents := []TestEntity{
		{ID: "1", Group: "X"},
		{ID: "2", Group: "X"},
		{ID: "3", Group: "X"},
	}
	writeJSONL(t, fp, ents)

	repo, err := NewJSONRepository[TestEntity](fp, getTestEntityID)
	require.NoError(t, err)

	sentinel := assert.AnError
	count := 0
	err = repo.FindAllWhere(
		func(e TestEntity) bool { return true },
		func(e TestEntity) error {
			count++
			if e.ID == "2" {
				return sentinel
			}
			return nil
		},
	)
	require.Error(t, err)
	require.ErrorIs(t, err, sentinel)
	require.Equal(t, 2, count, "should stop processing immediately after handler error")
}

func TestFindAllWhere_ReturnsScannerErrOnScanFailure(t *testing.T) {
	dir := t.TempDir()
	fp := filepath.Join(dir, "too_long.jsonl")

	tooLong := strings.Repeat("a", bufio.MaxScanTokenSize+10)
	line := `{"id":"` + tooLong + `"}`
	require.NoError(t, os.WriteFile(fp, []byte(line), 0o600))

	repo, err := NewJSONRepository[TestEntity](fp, getTestEntityID)
	require.NoError(t, err)

	called := false
	err = repo.FindAllWhere(
		func(e TestEntity) bool { return true },
		func(e TestEntity) error {
			called = true
			return nil
		},
	)
	require.Error(t, err)
	require.ErrorIs(t, err, bufio.ErrTooLong)
	require.False(t, called, "handler must not be called when scanner fails")
}
