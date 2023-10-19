package repository

import (
	"crypto/sha256"
	"encoding/hex"
)

type FileQuery interface {
	StoreFile(
		file []byte,
	) (string, error)
	GetFile(
		fileID string,
	) ([]byte, error)
}

type fileQuery struct{}

// GetFile implements FileQuery.
func (u *fileQuery) GetFile(fileID string) ([]byte, error) {
	return DB.Files[fileID], nil
}

func (u *fileQuery) StoreFile(
	file []byte) (string, error) {
	hash := ComputeHashAsString(file)
	DB.Files[hash] = file
	return hash, nil
}

func ComputeHashAsString(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
