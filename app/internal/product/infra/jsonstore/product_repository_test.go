package jsonstore_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/lucasti79/meli-interview/internal/product"
	"github.com/lucasti79/meli-interview/internal/product/infra/jsonstore"
	"github.com/lucasti79/meli-interview/internal/product/repository"
	"github.com/lucasti79/meli-interview/pkg/apperrors"
)

func writeProductsJSONL(t *testing.T, products []product.Product) string {
	t.Helper()

	dir := t.TempDir()
	fp := filepath.Join(dir, "products.jsonl")

	var lines []string
	for _, p := range products {
		data, err := json.Marshal(p)
		require.NoError(t, err)
		lines = append(lines, string(data))
	}

	content := strings.Join(lines, "\n")
	if len(lines) > 0 {
		content += "\n" // Add final newline
	}

	err := os.WriteFile(fp, []byte(content), 0o600)
	require.NoError(t, err)

	return fp
}

func newRepository(t *testing.T, filePath string) repository.Repository {
	t.Helper()
	repo, err := jsonstore.NewProductRepository(filePath)
	require.NoError(t, err)
	return repo
}

func TestGetByIDWithContext_ReturnsProductForValidID(t *testing.T) {
	fp := writeProductsJSONL(t, []product.Product{
		{Id: "123", Name: "Prod123", Category: "CatA", Price: 99.99},
		{Id: "456", Name: "Prod456", Category: "CatB", Price: 10.0},
	})
	repo := newRepository(t, fp)

	ctx := context.Background()
	got, err := repo.GetByIDWithContext(ctx, "123")
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, "123", got.Id)
	require.Equal(t, "Prod123", got.Name)
}

func TestGetByIDWithContext_UsesContextAndSucceedsBeforeDeadline(t *testing.T) {
	fp := writeProductsJSONL(t, []product.Product{
		{Id: "abc", Name: "TimeSensitive", Category: "CatX", Price: 1.23},
	})
	repo := newRepository(t, fp)

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	start := time.Now()
	got, err := repo.GetByIDWithContext(ctx, "abc")
	elapsed := time.Since(start)

	require.NoError(t, err)
	require.NotNil(t, got)
	require.True(t, elapsed < 300*time.Millisecond, "operation exceeded deadline: took %s", elapsed)
}

func TestGetByIDWithContext_MapsRowToDomainProduct(t *testing.T) {
	expected := product.Product{
		Id:       "m-001",
		Name:     "Mapped Product",
		Category: "Mapping",
		Price:    1234.56,
	}
	fp := writeProductsJSONL(t, []product.Product{expected})
	repo := newRepository(t, fp)

	ctx := context.Background()
	got, err := repo.GetByIDWithContext(ctx, expected.Id)
	require.NoError(t, err)
	require.NotNil(t, got)

	require.Equal(t, expected.Id, got.Id)
	require.Equal(t, expected.Name, got.Name)
	require.Equal(t, expected.Category, got.Category)
	require.Equal(t, expected.Price, got.Price)
}

func TestGetByIDWithContext_ProductNotFound(t *testing.T) {
	fp := writeProductsJSONL(t, []product.Product{
		{Id: "exists", Name: "Exists", Category: "Cat", Price: 5},
	})
	repo := newRepository(t, fp)

	ctx := context.Background()
	got, err := repo.GetByIDWithContext(ctx, "missing-id")
	require.Error(t, err)
	require.Nil(t, got)
	require.ErrorIs(t, err, apperrors.ErrResourceNotExists)
}

func TestGetByIDWithContext_PropagatesDatastoreError(t *testing.T) {
	// Create a valid repository and then remove the underlying file to induce a read error on lookup.
	fp := writeProductsJSONL(t, []product.Product{
		{Id: "x", Name: "X", Category: "Y", Price: 1},
	})
	repo := newRepository(t, fp)

	require.NoError(t, os.Remove(fp))

	ctx := context.Background()
	got, err := repo.GetByIDWithContext(ctx, "x")
	require.Error(t, err)
	require.Nil(t, got)
	// Ensure it's not mistaken as a not-found logical error
	require.False(t, errors.Is(err, apperrors.ErrResourceNotExists), "expected underlying datastore error, not a not-found error")
}

func TestGetAll_NoFilters_ReturnsAll(t *testing.T) {
	fp := writeProductsJSONL(t, []product.Product{
		{Id: "1", Name: "A", Category: "Cat1", Price: 10},
		{Id: "2", Name: "B", Category: "Cat2", Price: 20},
	})
	repo := newRepository(t, fp)

	products, total, err := repo.GetAll(product.ProductFilter{})
	require.NoError(t, err)
	require.Len(t, products, 2)
	require.Equal(t, 2, total)
}

func TestGetAll_FilterByName(t *testing.T) {
	fp := writeProductsJSONL(t, []product.Product{
		{Id: "1", Name: "Phone X", Category: "Electronics", Price: 100},
		{Id: "2", Name: "Shoes", Category: "Fashion", Price: 50},
	})
	repo := newRepository(t, fp)

	products, total, err := repo.GetAll(product.ProductFilter{Name: "phone"})
	require.NoError(t, err)
	require.Len(t, products, 1)
	require.Equal(t, 1, total)
	require.Equal(t, "Phone X", products[0].Name)
}

func TestGetAll_FilterByCategories(t *testing.T) {
	fp := writeProductsJSONL(t, []product.Product{
		{Id: "1", Name: "Phone", Category: "Electronics", Price: 100},
		{Id: "2", Name: "T-shirt", Category: "Fashion", Price: 20},
	})
	repo := newRepository(t, fp)

	products, total, err := repo.GetAll(product.ProductFilter{Categories: []string{"fashion"}})
	require.NoError(t, err)
	require.Len(t, products, 1)
	require.Equal(t, "T-shirt", products[0].Name)
	require.Equal(t, 1, total)
}

func TestGetAll_FilterByPriceRange(t *testing.T) {
	fp := writeProductsJSONL(t, []product.Product{
		{Id: "1", Name: "Cheap", Category: "Misc", Price: 5},
		{Id: "2", Name: "Mid", Category: "Misc", Price: 50},
		{Id: "3", Name: "Expensive", Category: "Misc", Price: 500},
	})
	repo := newRepository(t, fp)

	products, total, err := repo.GetAll(product.ProductFilter{MinPrice: 10, MaxPrice: 100})
	require.NoError(t, err)
	require.Len(t, products, 1)
	require.Equal(t, "Mid", products[0].Name)
	require.Equal(t, 1, total)
}

func TestGetAll_Pagination(t *testing.T) {
	fp := writeProductsJSONL(t, []product.Product{
		{Id: "1", Name: "P1", Category: "C", Price: 1},
		{Id: "2", Name: "P2", Category: "C", Price: 2},
		{Id: "3", Name: "P3", Category: "C", Price: 3},
	})
	repo := newRepository(t, fp)

	products, total, err := repo.GetAll(product.ProductFilter{Page: 2, PageSize: 1})
	require.NoError(t, err)
	require.Len(t, products, 1)
	require.Equal(t, "P2", products[0].Name)
	require.Equal(t, 3, total)
}

func TestGetAllWithContext_BehavesLikeGetAll(t *testing.T) {
	fp := writeProductsJSONL(t, []product.Product{
		{Id: "1", Name: "CtxProduct", Category: "Ctx", Price: 42},
	})
	repo := newRepository(t, fp)

	ctx := context.Background()
	products, total, err := repo.GetAllWithContext(ctx, product.ProductFilter{Categories: []string{"ctx"}})
	require.NoError(t, err)
	require.Len(t, products, 1)
	require.Equal(t, 1, total)
	require.Equal(t, "CtxProduct", products[0].Name)
}

func TestGetAll_NoResults(t *testing.T) {
	fp := writeProductsJSONL(t, []product.Product{
		{Id: "1", Name: "One", Category: "C1", Price: 10},
	})
	repo := newRepository(t, fp)

	products, total, err := repo.GetAll(product.ProductFilter{Name: "nonexistent"})
	require.NoError(t, err)
	require.Len(t, products, 0)
	require.Equal(t, 0, total)
}
