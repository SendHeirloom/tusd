package handler_test

import (
	"github.com/sendheirloom/tusd/pkg/filestore"
	"github.com/sendheirloom/tusd/pkg/handler"
	"github.com/sendheirloom/tusd/pkg/memorylocker"
)

func ExampleNewStoreComposer() {
	composer := handler.NewStoreComposer()

	fs := filestore.New("./data")
	fs.UseIn(composer)

	ml := memorylocker.New()
	ml.UseIn(composer)

	config := handler.Config{
		StoreComposer: composer,
	}

	_, _ = handler.NewHandler(config)
}
