package app

import (
	"context"
	"fmt"

	desc "github.com/DedicatedDev/merkleserver/pkg"
)

func (m *MicroserviceServer) UploadFile(ctx context.Context, req *desc.UploadFileRequest) (*desc.UploadFileResponse, error) {
	files := req.GetFiles()
	fileIDs := make([]string, len(files))
	for index, file := range files {
		fileID, err := m.fileService.StoreFile(file)
		if err != nil {
			return nil, err
		}
		m.merkleTreeService.AddLeafToTree(file)
		fileIDs[index] = fileID
	}
	return &desc.UploadFileResponse{
		FileIds: fileIDs,
	}, nil
}

func (m *MicroserviceServer) DownloadFile(ctx context.Context, req *desc.DownloadFileRequest) (*desc.DownloadFileResponse, error) {
	// Retrieve the file content by its ID
	fileID := req.GetFileId()
	fileContent, err := m.fileService.GetFile(fileID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving file: %v", err)
	}
	// Compute the Merkle proof for the requested file
	proof, verified := m.merkleTreeService.GenerateMerkleProof(fileContent)
	if !verified {
		return nil, fmt.Errorf("error generating Merkle proof: %v", err)
	}

	// Return both the file content and the Merkle proof to the client
	return &desc.DownloadFileResponse{
		Content: fileContent,
		MerkleProof: &desc.MerkleProof{
			Hashes: proof.Hashes,
			IsLeft: proof.IsLeft,
			Target: proof.Target,
		}, // this should be serialized in a format your client can understand
	}, nil
}
