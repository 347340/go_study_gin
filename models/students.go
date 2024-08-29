package models

import (
    "gorm.io/gorm"
)

type Students struct {
  Name string
  Age int
  gorm.Model
}

func main() {
  db.Create(&Students{Name:"王五", Age:"15"})
}
