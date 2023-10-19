package repository

import (
	"github.com/DedicatedDev/merkle"
	_ "github.com/go-sql-driver/mysql"
)

type DBHandler interface {
	NewFileQuery() FileQuery
	NewMerkleTreeQuery() MerkleTreeQuery
}

type dbHandler struct {
	db *MemoryDB
}

type MemoryDB struct {
	Tree  *merkle.MerkleTree
	Files map[string][]byte
}

var DB *MemoryDB

func NewDBHandler(db *MemoryDB) DBHandler {
	return &dbHandler{db}
}

func NewDB(dbName string) (*MemoryDB, error) {
	DB = &MemoryDB{
		Tree:  merkle.NewMerkleTree([][]byte{}),
		Files: make(map[string][]byte),
	}
	return DB, nil
}

func (d *dbHandler) NewFileQuery() FileQuery {
	return &fileQuery{}
}

func (d *dbHandler) NewMerkleTreeQuery() MerkleTreeQuery {
	return &merkleTreeQuery{}
}
