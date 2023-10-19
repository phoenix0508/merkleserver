package service

import (
	"github.com/DedicatedDev/merkleserver/internal/dto"
	"github.com/DedicatedDev/merkleserver/internal/repository"
)

type MerkleTreeService interface {
	VerifyMerkleProof(tree []byte) error
	AddLeafToTree(file []byte)
	GenerateMerkleProof(data []byte) (dto.MerkleProofDTO, bool)
}

type merkleTreeService struct {
	dbHandler repository.DBHandler
}

// GenerateMerkleProof implements MerkleTreeService.
func (m *merkleTreeService) GenerateMerkleProof(data []byte) (dto.MerkleProofDTO, bool) {
	return m.dbHandler.NewMerkleTreeQuery().GenerateMerkleProof(data)
}

// GrantBusiness implements AdminService
func (m *merkleTreeService) VerifyMerkleProof(tree []byte) error {
	return m.dbHandler.NewMerkleTreeQuery().VerifyMerkleProof(tree)
}

// GrantBusiness implements AdminService
func (m *merkleTreeService) AddLeafToTree(file []byte) {
	m.dbHandler.NewMerkleTreeQuery().AddLeaf(file)
}

func NewVerifyService(dbHandler repository.DBHandler) MerkleTreeService {
	return &merkleTreeService{dbHandler}
}
