package storage

import "opet/API/types"

type Storage interface {
	Get(int) *types.User
}
