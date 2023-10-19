package app

import (
	"github.com/DedicatedDev/merkleserver/internal/service"
	desc "github.com/DedicatedDev/merkleserver/pkg"
)

type MicroserviceServer struct {
	desc.UnimplementedMicroserviceServer
	fileService       service.FileService
	merkleTreeService service.MerkleTreeService
}

func NewMicroservice(
	fileService service.FileService,
	verifyService service.MerkleTreeService,

) *MicroserviceServer {
	return &MicroserviceServer{
		fileService:       fileService,
		merkleTreeService: verifyService,
	}
}
