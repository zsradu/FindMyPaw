package models

import "github.com/jinzhu/gorm"

type Type int

const (
	Undefined Type = iota
	Pisică
	Câine
)

func (t Type) toString() string {
	switch t {
	case Pisică:
		return "Pisică"
	case Câine:
		return "Câine"
	}
	return "Undefined"
}

type Location struct {
	gorm.Model
	X        float64
	Y        float64
	AnimalID uint
}

type Animal struct {
	gorm.Model
	AnimalType string `json:"type"`
	Breed      string
	Color      string
	Name       string
	Found      bool `json:"found"`
	Location   string
	FirstX     float64
	FirstY     float64
	LastX      float64
	LastY      float64
}
