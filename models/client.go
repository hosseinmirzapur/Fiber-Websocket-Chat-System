package models

import "sync"

type Client struct {
	IsClosing bool
	Mu        sync.Mutex
}
