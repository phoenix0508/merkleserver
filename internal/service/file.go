package service

import (
	"github.com/DedicatedDev/merkleserver/internal/repository"
)

type FileService interface {
	StoreFile([]byte) (string, error)
	GetFile(string) ([]byte, error)
}

type fileService struct {
	dbHandler repository.DBHandler
}

// GrantBusiness implements AdminService
func (a *fileService) StoreFile(file []byte) (string, error) {
	return a.dbHandler.NewFileQuery().StoreFile(file)
}

func (a *fileService) GetFile(file string) ([]byte, error) {
	return a.dbHandler.NewFileQuery().GetFile(file)
}

func NewFileService(dbHandler repository.DBHandler) FileService {
	return &fileService{dbHandler}
}
