package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"

	uuid "github.com/nu7hatch/gouuid"
)

type Post struct {
	Id      string  `json:"id,omitempty"`
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Author  *Author `json:"author"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	r := gin.Default()

	posts := []Post{
		{Id: "1", Title: "Post One", Content: "This is content", Author: &Author{Name: "babang", Email: "babang@babang.dev"}},
		{Id: "2", Title: "Post Two", Content: "This is content", Author: &Author{Name: "babang", Email: "babang@babang.dev"}},
		{Id: "3", Title: "Post Three", Content: "This is content", Author: &Author{Name: "babang", Email: "babang@babang.dev"}},
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	r.GET("/posts", func(c *gin.Context) {
		c.JSON(200, posts)
	})

	r.GET("/posts/:id", func(c *gin.Context) {
		id := c.Param("id")

		for _, item := range posts {
			if item.Id == id {
				c.JSON(200, item)
				return
			}
		}
		c.JSON(404, &Post{})
	})

	r.POST("/posts", func(c *gin.Context) {
		var post Post
		_ = json.NewDecoder(c.Request.Body).Decode(&post)

		newId, _ := uuid.NewV4()
		post.Id = newId.String()
		posts = append(posts, post)

		c.JSON(201, post)
	})

	r.PUT("/posts/:id", func(c *gin.Context) {
		for index, item := range posts {
			if item.Id == c.Param("id") {
				posts = append(posts[:index], posts[index+1:]...)

				var post Post
				_ = json.NewDecoder(c.Request.Body).Decode(&post)
				post.Id = c.Param("id")
				posts = append(posts, post)

				c.JSON(201, post)
				return
			}
		}
		c.JSON(403, gin.H{"message": "Forbidden action."})
	})

	r.DELETE("/posts/:id", func(c *gin.Context) {
		for index, item := range posts {
			if item.Id == c.Param("id") {
				posts = append(posts[:index], posts[index+1:]...)

				c.JSON(201, gin.H{"message": "Deleted", "post": item})
				return
			}
		}
		c.JSON(403, gin.H{"message": "Forbidden action."})
	})

	r.Run()
}
