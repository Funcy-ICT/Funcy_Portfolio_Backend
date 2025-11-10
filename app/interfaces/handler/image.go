package handler

import (
	"backend/app/interfaces/response"
	"backend/app/packages/storage"
	"context"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type ImageHandler struct {
	gcsClient *storage.GCSClient
}

func NewImageHandler(gcsClient *storage.GCSClient) *ImageHandler {
	return &ImageHandler{
		gcsClient: gcsClient,
	}
}

// UploadImage handles image upload to GCS
func (h *ImageHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Parse multipart form
	err := r.ParseMultipartForm(32 << 20) // 32 MB max
	if err != nil {
		log.Printf("Failed to parse multipart form: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "Failed to parse form")
		return
	}

	files := r.MultipartForm.File["file"]
	if len(files) == 0 {
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "No files provided")
		return
	}

	var urls []string

	for _, fileHeader := range files {
		// Open uploaded file
		file, err := fileHeader.Open()
		if err != nil {
			log.Printf("Failed to open uploaded file: %v", err)
			_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, "Failed to open file")
			return
		}
		defer file.Close()

		// Generate unique filename
		ext := filepath.Ext(fileHeader.Filename)
		fileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)

		// Upload to GCS
		url, err := h.gcsClient.Upload(ctx, fileName, file)
		if err != nil {
			log.Printf("Failed to upload to GCS: %v", err)
			_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, "Failed to upload file")
			return
		}

		urls = append(urls, url)
	}

	// Return response in same format as file-server
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"urls":%q}`, urls)
}

// DeleteImage handles image deletion from GCS
func (h *ImageHandler) DeleteImage(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	fileName := chi.URLParam(r, "filename")
	if fileName == "" {
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "filename is required")
		return
	}

	err := h.gcsClient.Delete(ctx, fileName)
	if err != nil {
		log.Printf("Failed to delete from GCS: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, "Failed to delete file")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message":"File deleted successfully"}`)
}
