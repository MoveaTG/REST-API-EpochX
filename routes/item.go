package routes

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sixfwa/fiber-api/database"
	"github.com/sixfwa/fiber-api/models"
)

type ItemSerializer struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Year      YearSerializer `gorm:"foreignKey:YaerRefer"`
	Date      string         `json:"date"`
	Name      string         `json:"name"`
	Text      string         `json:"text" gorm:"text"`
	ImageReal string         `json:"imageReal"` // changed from [3]string
	ImageAi   string         `json:"imageAi"`   // changed from [3]string
}

func CreateResponseItem(item models.Item, year YearSerializer) ItemSerializer {
	return ItemSerializer{
		ID:        item.ID,
		CreatedAt: item.CreatedAt,
		Year:      year,
		Date:      item.Date,
		Name:      item.Name,
		Text:      item.Text,
		ImageReal: item.ImageReal,
		ImageAi:   item.ImageAi,
	}
}

func CreateItem(c *fiber.Ctx) error {
	var item models.Item

	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var year models.Year
	if err := findYear(item.YaerRefer, &year); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&item)

	responseYear := CreateResponseYear(year)
	responseItem := CreateResponseItem(item, responseYear)

	return c.Status(200).JSON(responseItem)
}

func GetItems(c *fiber.Ctx) error {
	items := []models.Item{}

	database.Database.Db.Find(&items)
	responseItems := []ItemSerializer{}
	for _, item := range items {
		var year models.Year
		database.Database.Db.Find(&year, "id = ?", item.YaerRefer)

		responseItem := CreateResponseItem(item, CreateResponseYear(year))
		responseItems = append(responseItems, responseItem)
	}
	return c.Status(200).JSON(responseItems)
}

func FindItem(id int, item *models.Item) error {
	database.Database.Db.Find(&item, "id = ?", id)
	if item.ID == 0 {
		return errors.New("order does not exist")
	}

	return nil
}

func GetItem(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var item models.Item

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindItem(id, &item); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var year models.Year
	database.Database.Db.First(&year, item.YaerRefer)
	responseYear := CreateResponseYear(year)

	responseItem := CreateResponseItem(item, responseYear)
	return c.Status(200).JSON(responseItem)
}
