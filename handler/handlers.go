package handler

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GradeType float32

const (
	GradeDunno GradeType = -0.5
	GradeHard  GradeType = -0.25
	GradeGood  GradeType = 0.25
	GradeEasy  GradeType = 0.5
)

type Tag struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"not null,unique"`
}

type Repository struct {
	Database *gorm.DB
}

type Review struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	Grade     GradeType
}

type Item struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	Stability float64
	// Reviews   []Review
	Tags []Tag `gorm:"many2many:item_tags;"`
}

type ItemInput struct {
	ID        string    `json:"id" binding:"required"`
	Tags      []Tag     `json:"tags" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

func (r *Repository) AddItem(c *gin.Context) {
	var input ItemInput
	c.ShouldBindJSON(&input)

	fmt.Println("input before inside db:", input)

	item := Item{
		ID:        input.ID,
		CreatedAt: input.CreatedAt,
		Stability: 0.0,
		Tags:      input.Tags,
	}

	r.Database.Create(&item)

	fmt.Println("newItem is ", item)

}

func (d *Repository) GetItemById(c *gin.Context) {
	id := c.Params.ByName("id")
	var item Item
	result := d.Database.First(&item, "id = ?", id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": result.Error})
		return
	}
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("item is: %v", item),
	})
	// Get the item from the database
	// if found

	//if not found
	// c.JSON(404, gin.H{"message": "item not found"})
}

// func ItemsByTagHandler(w http.ResponseWriter, r *http.Request) {
// }

// func AddItemHandler(w http.ResponseWriter, r *http.Request) {
// }

// func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
// }
