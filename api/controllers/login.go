package controllers

import (
	"api/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Retrieving all the user details
func GetUser(c *gin.Context) {
	var user []models.User
	_, err := dbmap.Select(&user, "select * from user")
	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}

}

// Retrieving all the posts details
func GetPosts(c *gin.Context) {
	var post []models.Posts
	_, err := dbmap.Select(&post, "select * from posts")
	if err == nil {

		c.JSON(200, post)
	} else {

		c.JSON(404, gin.H{"error": " posts not found"})
	}

}

//Retrieve a user by id
func GetUserDetail(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=? LIMIT 1", id)
	if err == nil {
		user_id, _ := strconv.ParseInt(id, 0, 64)
		content := &models.User{
			Id:        user_id,
			Username:  user.Username,
			Password:  user.Password,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}

//Retrieve a post by id
func GetPostDetail(c *gin.Context) {
	id := c.Params.ByName("postid")

	var post models.Posts
	err := dbmap.SelectOne(&post, "SELECT * FROM posts WHERE postid=? LIMIT 1 ", id)
	if err == nil {
		post_id, _ := strconv.ParseInt(id, 0, 64)
		content := &models.Posts{
			Postid: post_id,
			Userid: post.Userid,
			Title:  post.Title,
			Body:   post.Body,
		}
		c.JSON(200, content)

	} else {
		c.JSON(404, gin.H{" error": "Post not found"})
	}

}

//Retrive all posts of specific user by userid
func GetPostsDetailByUserID(c *gin.Context) {
	id := c.Params.ByName("userid")

	var post []models.Posts

	_, err := dbmap.Select(&post, "SELECT * FROM posts WHERE userid=?  ", id)

	if err == nil {
		c.JSON(200, post)
	} else {
		c.JSON(404, gin.H{" error": "Post not found"})
	}

}

// Login function with database validation
func Login(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	err := dbmap.SelectOne(&user, "select * from user where Username=? LIMIT 1", user.Username)
	if err == nil {
		user_id := user.Id
		content := &models.User{
			Id:        user_id,
			Username:  user.Username,
			Password:  user.Password,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
		}
		c.JSON(200, content)

	} else {
		c.JSON(404, gin.H{"error": "user not found"})

	}
}

//Inserting a user detail
func PostUser(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	log.Println(user)
	if user.Username != "" && user.Password != "" && user.Firstname != "" && user.Lastname != "" {
		if insert, _ := dbmap.Exec(`INSERT INTO user (Username, Password, Firstname, Lastname) VALUES (?, ?, ?, ?)`, user.Username, user.Password, user.Firstname, user.Lastname); insert != nil {
			user_id, err := insert.LastInsertId()
			if err == nil {
				content := &models.User{
					Id:        user_id,
					Username:  user.Username,
					Password:  user.Password,
					Firstname: user.Firstname,
					Lastname:  user.Lastname,
				}
				c.JSON(201, content)
			} else {
				checkErr(err, "Insert failed")
			}
		}
	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}
}

// inserting a post detail
/*func AddPost( c *gin.Context){
	var post models.Posts
	c.Bind(&post)
	log.Println(&user)

	if post.Userid !=  &&  post.Title !=  " "   && post.Body != "" {

 		if insert , _ := dbmap.Exec(`INSERT INTO posts (Userid, Title, Body) VALUES (?, ?, ?)`,post.Userid , post.Title , post.Body); insert != nil{
			post_id , err := inset.LastInsertId()

			if err == nil {

				content := &models.Posts{
						Postid : post_id,
						Userid : post.Userid,
						Title  : post.Title,
						Body   : post.Body,


				}
				c.JSON(201 , content)
				}else {
					checkErr(err, "Insert failed")
				}




			}

		}else {
			c.JSON(400, gin.H{"error": "Fields are empty"})
		}

	}
*/

//deleting a user
func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")

	_, err := dbmap.Exec("DELETE FROM user WHERE ID = ?", id)

	if err == nil {

		c.JSON(200, gin.H{" Deleted": "Success"})
	} else {

		c.JSON(404, gin.H{"Error": "Not Deleted"})
	}

}

//deleting a post
func DeletePost(c *gin.Context) {

	id := c.Params.ByName("postid")

	_, err := dbmap.Exec("DELETE FROM posts WHERE Postid = ?", id)

	if err == nil {

		c.JSON(200, gin.H{" Deleted": "Success"})
	} else {

		c.JSON(404, gin.H{"Error": "Not Deleted"})
	}

}

// Updating a user details
func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=?", id)
	if err == nil {
		var uuser models.User
		c.Bind(&uuser)
		user_id, _ := strconv.ParseInt(id, 0, 64)
		user := models.User{
			Id:        user_id,
			Username:  uuser.Username,
			Password:  uuser.Password,
			Firstname: uuser.Firstname,
			Lastname:  uuser.Lastname,
		}
		if user.Firstname != "" && user.Lastname != "" {
			_, err = dbmap.Update(&user)
			if err == nil {
				c.JSON(200, user)
			} else {
				checkErr(err, "Updated failed")
			}
		} else {
			c.JSON(400, gin.H{"error": "fields are empty"})
		}
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}

// Updating a post details
func UpdatePost(c *gin.Context) {
	id := c.Params.ByName("postid")

	var post models.Posts

	err := dbmap.SelectOne(&post, " SELECT * FROM posts WHERE postid= ?", id)

	if err == nil {
		var ppost models.Posts
		c.Bind(&ppost)

		post_id, _ := strconv.ParseInt(id, 0, 64)

		post := models.Posts{
			Postid: post_id,
			Userid: ppost.Userid,
			Title:  ppost.Title,
			Body:   ppost.Body,
		}
		if post.Title == "" && post.Body == "" {
			_, err := dbmap.Update(&post)
			if err == nil {
				c.JSON(200, post)
			} else {
				checkErr(err, "Updated failed")
			}

		} else {
			c.JSON(400, gin.H{"error": "fields are empty"})
		}

	} else {
		c.JSON(404, gin.H{"error": "post not found"})
	}

}
