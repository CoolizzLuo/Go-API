package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/role", Get)

	router.GET("/role/:id", GetOne)

	router.POST("/role", Post)

	router.PUT("/role/:id", Put)

	router.DELETE("/role/:id", Delete)

	router.Run(":8080")
}

// 取得全部資料
func Get(c *gin.Context) {
	c.JSON(http.StatusOK, Data)
}

// 取得單一筆資料
func GetOne(c *gin.Context) {
	// strconv.Atoi 轉換成 Int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	for i := 0; i < len(Data); i++ {
		if Data[i].ID == id {
			c.JSON(http.StatusOK, Data[i])
			break
		}
	}
}

// 新增資料
func Post(c *gin.Context) {
	var r Role
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	Data = append(Data, r)
	c.JSON(http.StatusOK, gin.H{"message": r.Name + " is insert"})
}

type RoleVM struct {
	ID      int    `json:"id"`      // Key
	Name    string `json:"name"`    // 角色名稱
	Summary string `json:"summary"` // 介紹
}

// 更新資料, 更新角色名稱與介紹
func Put(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var r RoleVM
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.BindJSON(&r)
	for i := 0; i < len(Data); i++ {
		if Data[i].ID == id {
			Data[i].Name = r.Name
			Data[i].Summary = r.Summary
			c.JSON(http.StatusOK, r)
			break
		}
	}
}

// 刪除資料
func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	dl := len(Data)
	for i := 0; i < dl; i++ {
		if Data[i].ID == id {
			Data[i] = Data[dl-1]
			Data = Data[:dl-1]
			break
		}
	}
	c.Status(http.StatusNoContent)

}
