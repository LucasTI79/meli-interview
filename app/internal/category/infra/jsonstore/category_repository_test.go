package jsonstore_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lucasti79/meli-interview/internal/category"
	"github.com/lucasti79/meli-interview/internal/category/infra/jsonstore"
	"github.com/lucasti79/meli-interview/internal/category/repository"
	"github.com/lucasti79/meli-interview/pkg/apperrors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func newRepository(t *testing.T, filePath string) repository.Repository {
	t.Helper()
	repo, err := jsonstore.NewCategoryRepository(filePath)
	require.NoError(t, err)
	return repo
}

func newCategoryRepositoryForTest(t *testing.T, filePath string) repository.Repository {
	t.Helper()
	r, err := jsonstore.NewCategoryRepository(filePath)
	require.NoError(t, err)

	cr, ok := r.(repository.Repository)
	require.True(t, ok, "expected repository to be *repository.CategoryRepository")
	return cr
}

type DummyMock struct{ mock.Mock }

func writeProductsJSONL(t *testing.T, categories []string) string {
	t.Helper()

	dir := t.TempDir()
	fp := filepath.Join(dir, "products.jsonl")

	var lines []string
	for _, c := range categories {
		data, err := json.Marshal(map[string]any{
			"category": c,
		})
		require.NoError(t, err)
		lines = append(lines, string(data))
	}

	content := strings.Join(lines, "\n")
	if len(lines) > 0 {
		content += "\n"
	}

	err := os.WriteFile(fp, []byte(content), 0o600)
	require.NoError(t, err)

	return fp
}

func writeInvalidProductsJSONL(t *testing.T) string {
	t.Helper()

	dir := t.TempDir()
	fp := filepath.Join(dir, "products_invalid.jsonl")

	content := `{"category":"valid"}
not-a-json
`

	err := os.WriteFile(fp, []byte(content), 0o600)
	require.NoError(t, err)
	return fp
}

func TestCategoryRepository_GetAll_ReturnsUniqueCategories(t *testing.T) {
	fp := writeProductsJSONL(t, []string{"electronics", "fashion", "electronics", "home"})

	repo := newRepository(t, fp)

	got, err := repo.GetAll()
	require.NoError(t, err)

	expected := []category.Category{
		{Name: "electronics"},
		{Name: "fashion"},
		{Name: "home"},
	}
	require.ElementsMatch(t, expected, got)
}

func TestCategoryRepository_GetByName_ReturnsExistingCategory(t *testing.T) {
	fp := writeProductsJSONL(t, []string{"A", "B"})

	repo := newRepository(t, fp)

	got, err := repo.GetByName("B")
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "B", got.Name)
}

func TestCategoryRepository_GetAllWithContext_ReturnsSameAsGetAll(t *testing.T) {
	fp := writeProductsJSONL(t, []string{"x", "y", "x"})

	repo := newRepository(t, fp)

	got1, err := repo.GetAll()
	require.NoError(t, err)

	got2, err := repo.GetAllWithContext(context.Background())
	require.NoError(t, err)

	require.ElementsMatch(t, got1, got2)
}

func TestNewCategoryRepository_ReturnsErrorOnJSONRepositoryFailure(t *testing.T) {
	dir := t.TempDir()

	r, err := jsonstore.NewCategoryRepository(dir)
	require.Error(t, err)
	require.Nil(t, r)
}

func TestCategoryRepository_GetAll_PropagatesFindAllError(t *testing.T) {
	fp := writeInvalidProductsJSONL(t)

	repo := newRepository(t, fp)

	_, err := repo.GetAll()
	require.Error(t, err)
}

func TestCategoryRepository_GetByName_ReturnsErrResourceNotExistsWhenMissing(t *testing.T) {
	fp := writeProductsJSONL(t, []string{"A", "B"})

	repo := newRepository(t, fp)

	got, err := repo.GetByName("C")
	require.ErrorIs(t, err, apperrors.ErrResourceNotExists)
	require.Nil(t, got)
}

func TestGetByNameWithContext_ReturnsCategoryOnExactMatch(t *testing.T) {
	fp := writeProductsJSONL(t, []string{"home", "garden", "toys"})
	repo := newCategoryRepositoryForTest(t, fp)

	got, err := repo.GetByNameWithContext(context.Background(), "garden")
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "garden", got.Name)
}

func TestGetByNameWithContext_SucceedsWhenMultipleMatchesExist(t *testing.T) {
	fp := writeProductsJSONL(t, []string{"A", "B", "B", "C"})
	repo := newCategoryRepositoryForTest(t, fp)

	got, err := repo.GetByNameWithContext(context.Background(), "B")
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "B", got.Name)
}

func TestGetByNameWithContext_HandlesUnicodeCategoryNames(t *testing.T) {
	name := "Café-電子📦"
	fp := writeProductsJSONL(t, []string{"alpha", name, "omega"})
	repo := newCategoryRepositoryForTest(t, fp)

	got, err := repo.GetByNameWithContext(context.Background(), name)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, name, got.Name)
}

func TestGetByNameWithContext_ReturnsErrResourceNotExistsWhenNoMatch(t *testing.T) {
	fp := writeProductsJSONL(t, []string{"A", "B"})
	repo := newCategoryRepositoryForTest(t, fp)

	got, err := repo.GetByNameWithContext(context.Background(), "C")
	require.ErrorIs(t, err, apperrors.ErrResourceNotExists)
	require.Nil(t, got)
}

func TestGetByNameWithContext_PropagatesFindAllWhereError(t *testing.T) {
	fp := writeInvalidProductsJSONL(t)
	repo := newCategoryRepositoryForTest(t, fp)

	got, err := repo.GetByNameWithContext(context.Background(), "valid")
	require.Error(t, err)
	require.Nil(t, got)
}

func TestGetByNameWithContext_NoMatchWhenCaseDiffers(t *testing.T) {
	fp := writeProductsJSONL(t, []string{"Books"})
	repo := newCategoryRepositoryForTest(t, fp)

	got, err := repo.GetByNameWithContext(context.Background(), "books")
	require.ErrorIs(t, err, apperrors.ErrResourceNotExists)
	require.Nil(t, got)
}
