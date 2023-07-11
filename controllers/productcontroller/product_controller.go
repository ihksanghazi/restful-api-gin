package productcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/restful-api-gin/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context){
	var products []models.Products

	models.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func Show(c *gin.Context){
	var product models.Products
	id:= c.Param("id")

	if err := models.DB.First(&product,id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound,gin.H{"msg":"Data Tidak Ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"msg":err.Error()})
			return
		}		
	}

	c.JSON(http.StatusOK,gin.H{"product":product})
}
func Create(c *gin.Context){
	var product models.Products

	if err:= c.BindJSON(&product); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"msg":err.Error()})
		return
	}

	models.DB.Create(&product)

	c.JSON(http.StatusOK,gin.H{"product":product})
}
func Update(c *gin.Context){
	var product models.Products
	id:= c.Param("id")

	if err:= c.BindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"msg":err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ?",id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"msg":"Tidak Dapat Mengupdate Product"})
		return
	}

	c.JSON(http.StatusOK,gin.H{"msg":"Data Berhasil Diperbarui"})
}
func Delete(c *gin.Context){
	var product models.Products
	id:=c.Param("id")

	if models.DB.Delete(&product,id).RowsAffected == 0{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"msg":"Gagal Hapus Data"})
		return
	}

	c.JSON(http.StatusOK,gin.H{"msg":"Data Berhasil Dihapus"})
}