package model

type NetworkSocket interface {
	Read() (Message, error)
}
