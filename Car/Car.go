package Car

import (
	"PemWeb_BE/Auth"
	"PemWeb_BE/User"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Routes(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/mobil")
	r.POST("/managemen", Auth.Authorization(), func(c *gin.Context) {
		var input Add
		id, _ := c.Get("id")
		user := User.User{}
		if err := db.Where("id=?", id).Take(&user); err.Error != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Query user gagal",
				"error":   err.Error.Error(),
			})
		}
		if user.Telepon != "081913910239" {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Halaman ini hanya dapat diakses Administrator",
				"error":   "Forbidden Access",
			})
		}
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Value sent was incorrect",
				"error":   err.Error(),
			})
			return
		}
		create := Car{
			Name:   input.Name,
			Engine: input.Engine,
			Price:  input.Price,
			Stock:  input.Stock,
			Photo:  input.Photo,
		}
		if err := db.Create(&create); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong with the database row creation",
				"error":   err.Error.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Mobil berhasil ditambahkan",
			"error":   nil,
			"data":    create,
		})
	})
	r.PATCH("/managemen", Auth.Authorization(), func(c *gin.Context) {
		id, _ := c.Get("id")
		user := User.User{}
		if err := db.Where("id=?", id).Take(&user); err.Error != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Query user gagal",
				"error":   err.Error.Error(),
			})
		}
		if user.Telepon != "081913910239" {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Halaman ini hanya dapat diakses Administrator",
				"error":   "Forbidden Access",
			})
		}
		var input Car
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Value sent was incorrect",
				"error":   err.Error(),
			})
			return
		}
		update := Car{}
		if err := db.Where("id=?", input.ID).Take(&update); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Car was not found or ID was not given",
				"error":   err.Error.Error(),
			})
			return
		}
		if len(input.Photo) <= 0 {
			input.Photo = " "
		}
		update = Car{
			ID:     update.ID,
			Name:   input.Name,
			Engine: input.Engine,
			Photo:  input.Photo,
			Price:  input.Price,
			Stock:  input.Stock,
		}
		err := db.Model(&update).Updates(update)
		if err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong on car update",
				"error":   err.Error.Error(),
			})
			return
		}
		if err.RowsAffected < 1 {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "No data has been changed",
				"error":   err.Error.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Data updated successfully",
			"error":   nil,
			"data": gin.H{
				"id":     update.ID,
				"nama":   update.Name,
				"mesin":  update.Engine,
				"foto":   update.Photo,
				"harga":  update.Price,
				"jumlah": update.Stock,
			},
		})
	})
	r.DELETE("/managemen", Auth.Authorization(), func(c *gin.Context) {
		id, _ := c.Get("id")
		user := User.User{}
		if err := db.Where("id=?", id).Take(&user); err.Error != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Query user gagal",
				"error":   err.Error.Error(),
			})
		}
		if user.Telepon != "081913910239" {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Halaman ini hanya dapat diakses Administrator",
				"error":   "Forbidden Access",
			})
		}
		var input Car
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Value sent was incorrect",
				"error":   err.Error(),
			})
			return
		}
		delete := Car{}
		if err := db.Where("id=?", input.ID).Take(&delete); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Car was not found",
				"error":   err.Error.Error(),
			})
			return
		}
		if err := db.Delete(&delete); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to delete car",
				"error":   err.Error.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Account deleted successfully",
			"error":   nil,
			"data": gin.H{
				"id":     delete.ID,
				"nama":   delete.Name,
				"mesin":  delete.Engine,
				"foto":   delete.Photo,
				"harga":  delete.Price,
				"jumlah": delete.Stock,
			},
		})
	})
	r.GET("/", func(c *gin.Context) {
		var query []Car
		if err := db.Find(&query); err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Something went wrong in query",
				"error":   err.Error.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"error":   nil,
			"message": "Query successful",
			"result":  query,
		})
	})
}
