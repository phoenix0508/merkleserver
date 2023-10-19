package app

import (
	"context"

	desc "github.com/DedicatedDev/merkleserver/pkg"
)

func (m *MicroserviceServer) VerifyMerkleProof(ctx context.Context, req *desc.VerifyMerkleProofRequest) (*desc.VerifyMerkleProofResponse, error) {

	return &desc.VerifyMerkleProofResponse{}, nil
}
