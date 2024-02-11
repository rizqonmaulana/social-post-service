package models

type Post struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  Content  string `json:"title"`
  Author string `json:"author"`
}