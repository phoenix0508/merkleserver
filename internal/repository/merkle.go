package repository

import (
	"github.com/DedicatedDev/merkle"
	"github.com/DedicatedDev/merkleserver/internal/dto"
)

type MerkleTreeQuery interface {
	VerifyMerkleProof(
		file []byte,
	) error

	AddLeaf(hash []byte)

	GenerateMerkleProof(data []byte) (dto.MerkleProofDTO, bool)
}

type merkleTreeQuery struct{}

// AddLeaf adds a new leaf to the tree and recalculates the tree.
func (u *merkleTreeQuery) AddLeaf(data []byte) {
	if DB.Tree == nil {
		DB.Tree = merkle.NewMerkleTree([][]byte{data})
	} else {
		DB.Tree.AddLeaf(data)
	}
}

func (u *merkleTreeQuery) VerifyMerkleProof(
	file []byte) error {

	return nil
}

func (u *merkleTreeQuery) GenerateMerkleProof(data []byte) (dto.MerkleProofDTO, bool) {
	// Convert the data to its hash
	dataHash := merkle.ComputeHashAsString(data)
	if DB.Tree != nil {
		proof, ok := DB.Tree.GenerateProof(dataHash)
		return dto.MerkleProofDTO(proof), ok
	} else {
		return dto.MerkleProofDTO{}, false
	}
}

func findParentAndSibling(root, node *merkle.MerkleNode) (parent, sibling *merkle.MerkleNode) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return nil, nil
	}

	if root.Left == node {
		return root, root.Right
	}
	if root.Right == node {
		return root, root.Left
	}

	// Search left subtree
	parent, sibling = findParentAndSibling(root.Left, node)
	if parent != nil {
		return parent, sibling
	}

	// Search right subtree
	return findParentAndSibling(root.Right, node)
}
