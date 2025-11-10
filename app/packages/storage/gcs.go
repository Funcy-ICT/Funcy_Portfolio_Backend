package storage

import (
	"context"
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/storage"
)

type GCSClient struct {
	client     *storage.Client
	bucketName string
}

// NewGCSClient creates a new GCS client
func NewGCSClient(ctx context.Context) (*GCSClient, error) {
	bucketName := os.Getenv("GCS_BUCKET_NAME")
	if bucketName == "" {
		return nil, fmt.Errorf("GCS_BUCKET_NAME environment variable is not set")
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCS client: %w", err)
	}

	return &GCSClient{
		client:     client,
		bucketName: bucketName,
	}, nil
}

// Upload uploads a file to GCS and returns the public URL
func (g *GCSClient) Upload(ctx context.Context, fileName string, file io.Reader) (string, error) {
	bucket := g.client.Bucket(g.bucketName)
	obj := bucket.Object(fileName)

	writer := obj.NewWriter(ctx)
	defer writer.Close()

	if _, err := io.Copy(writer, file); err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %w", err)
	}

	// Return public URL
	// For emulator: http://localhost:4443/storage/v1/b/{bucket}/o/{object}?alt=media
	// For production: https://storage.googleapis.com/{bucket}/{object}
	emulatorHost := os.Getenv("GCS_EMULATOR_HOST")
	if emulatorHost != "" {
		// Use localhost for external access with GCS API path
		return fmt.Sprintf("http://localhost:4443/storage/v1/b/%s/o/%s?alt=media", g.bucketName, fileName), nil
	}

	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", g.bucketName, fileName), nil
}

// Delete deletes a file from GCS
func (g *GCSClient) Delete(ctx context.Context, fileName string) error {
	bucket := g.client.Bucket(g.bucketName)
	obj := bucket.Object(fileName)

	if err := obj.Delete(ctx); err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

// Close closes the GCS client
func (g *GCSClient) Close() error {
	return g.client.Close()
}

// CreateBucketIfNotExists creates a bucket if it doesn't exist (for emulator)
func (g *GCSClient) CreateBucketIfNotExists(ctx context.Context) error {
	bucket := g.client.Bucket(g.bucketName)

	// Check if bucket exists
	_, err := bucket.Attrs(ctx)
	if err == storage.ErrBucketNotExist {
		// Create bucket for emulator
		if err := bucket.Create(ctx, os.Getenv("GCS_PROJECT_ID"), &storage.BucketAttrs{
			Location: "US",
		}); err != nil {
			return fmt.Errorf("failed to create bucket: %w", err)
		}
		return nil
	}

	return err
}
