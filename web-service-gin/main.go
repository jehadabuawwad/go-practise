package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)


type DataItem struct {
    ID     string  `json:"id"`
    Name  string  `json:"name"`
    DESCRIPTION string  `json:"description"`
    IMAGEURL  string `json:"imageUrl"`
}


func InitDb() *gorm.DB {
    db, err := gorm.Open("sqlite3", "./data.db")
    db.LogMode(true)

    if err != nil {
        panic(err)
    }

    if !db.HasTable(&DataItem{}) {
        db.CreateTable(&DataItem{})
        db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&DataItem{})
    }

    return db
}

func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func getData(c *gin.Context){
	db := InitDb()
    defer db.Close()

    var dataItems []DataItem
    db.Find(&dataItems)
    c.JSON(200, dataItems)
}


func createData(c *gin.Context) {

	db := InitDb()
    defer db.Close()

    var data DataItem
    c.Bind(&data)

    if data.ID != ""  {
    
        db.Create(&data)
    
        c.JSON(201, gin.H{"success": data})
    } else {
    
        c.JSON(422, gin.H{"error": "Fields are empty"})
    }
}


func updateData(c *gin.Context){
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var dataItem DataItem
	db.First(&dataItem, id)

	if dataItem.ID != ""{

		if dataItem.ID != "0" {
			var newDataItem DataItem
			c.Bind(&newDataItem)

			result := DataItem {
				ID:        dataItem.ID,
				Name: newDataItem.Name,
				DESCRIPTION:  newDataItem.DESCRIPTION,
				IMAGEURL:  newDataItem.IMAGEURL,
			}


			db.Save(&result)

			c.JSON(200, gin.H{"success": result})
		} else {
			c.JSON(404, gin.H{"error": "User not found"})
		}

	} else {
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

func deleteData(c *gin.Context){
    db := InitDb()
    defer db.Close()

    id := c.Params.ByName("id")
    var dataItem DataItem

    db.First(&dataItem, id)

    if dataItem.ID != "0" {
        db.Delete(&dataItem)
        c.JSON(200, gin.H{"success": "User #" + id + " deleted"})
    } else {
        c.JSON(404, gin.H{"error": "User not found"})
    }
}


func main(){
	r := gin.Default()
	r.Use(Cors())
    v1 := r.Group("api/v1")
    {
        v1.POST("/products", createData)
        v1.GET("/products", getData)
        v1.PUT("/products/:id", updateData)
        v1.DELETE("/products/:id", deleteData)
    }

    r.Run(":8000")

}