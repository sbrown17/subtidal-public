package main

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var boardList = [...]string{"programming", "kcdc2024"}

type Authed struct {
	Authed bool
}

type Boards struct {
	Boards []Board
}

type Board struct {
	Name    string
	Threads []Thread
}

type Thread struct {
	Board    string
	Name     string
	Id       int
	Position int
	Posts    []Post
	Body     string
}

type Post struct {
	Name string
	Id   int
	Body string
}

var globalId = 0

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	//auth for admin stuff
	router.GET("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "DontUseThisAsAPassword",
	}), func(c *gin.Context) {
		cookie, err := c.Cookie("sesh_cook")

		if err != nil {
			c.SetCookie("sesh_cook", sessionId, 300, "/", "localhost", false, true) // false should be set to true before deploy, also need to change domain from localhost to subtidal.org
		}
		if cookie == "" {
			c.SetCookie("sesh_cook", sessionId, 300, "/", "localhost", false, true)
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"boardList": boardList,
		})
	})
	router.GET("/login", getLoginPage)
	router.GET("/", getHome)
	router.GET("/:board", getBoard)
	router.GET("/:board/:thread", getThread)
	router.POST("/:board", postThread)
	router.POST("/:board/:thread", postPostInAThread)
	router.POST("/:board/remove/thread/:id", removeThread)
	router.POST("/:board/remove/post/:id", removePost)
	router.Static("/styles", "./styles")
	router.Run("localhost:8080")
}

func getLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
func getHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"boardList": boardList,
	})
}

func postThread(c *gin.Context) {
	var newThread Thread
	globalId = globalId + 1
	newThread.Id = globalId
	newThread.Position = globalId
	newThread.Board = c.Param("board")
	newThread.Name = c.PostForm("name")
	newThread.Body = c.PostForm("body")

	for i, b := range boards.Boards {
		if b.Name == newThread.Board {
			// newThread.Id = mockBoards.Boards[i].Threads[len(mockBoards.Boards[i].Threads)-1].Id + 1

			boards.Boards[i].Threads = append(b.Threads, newThread)

			c.HTML(http.StatusOK, "board.html", gin.H{
				"boardList": boardList,
				"board":     boards.Boards[i],
			})
			break
		}
	}
}

func getBoard(c *gin.Context) {
	// set authed here by checking cookie
	var returnBoard Board
	var boardName string
	seshCookie, _ := c.Cookie("sesh_cook")
	showAdminTools := false
	if seshCookie == sessionId {
		showAdminTools = true
	}
	for _, b := range boards.Boards {
		if b.Name == c.Param("board") {
			returnBoard = b
			boardName = b.Name
			sort.Slice(b.Threads, func(i, j int) bool {
				return b.Threads[i].Position > b.Threads[j].Position // this sorts by thread id, but need to sort by post id
			})
			c.HTML(http.StatusOK, "board.html", gin.H{
				"boardList":      boardList,
				"board":          returnBoard,
				"boardName":      boardName,
				"showAdminTools": showAdminTools,
			})
		}
	}
}

func getThread(c *gin.Context) {
	// set authed here by checking cookie
	boardName := c.Param("board")
	threadId, err := strconv.Atoi(c.Param("thread"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "", gin.H{})
	}
	seshCookie, err := c.Cookie("sesh_cook")
	showAdminTools := false
	if seshCookie == sessionId {
		showAdminTools = true
	}

	var returnThread Thread
	for i, b := range boards.Boards {
		if b.Name == boardName {
			for _, t := range boards.Boards[i].Threads {
				if t.Id == threadId {
					returnThread = t
				}
			}
		}
	}

	c.HTML(http.StatusOK, "thread.html", gin.H{
		"boardList":      boardList,
		"boardName":      boardName,
		"thread":         returnThread,
		"showAdminTools": showAdminTools,
	})
}

func postPostInAThread(c *gin.Context) {
	//this doesn't check for admin on return... all things that return values should probably check for that
	var returnThread Thread
	var newPost Post
	globalId = globalId + 1
	newPost.Id = globalId

	newPost.Name = c.PostForm("name")
	newPost.Body = c.PostForm("body")
	boardName := c.Param("board")
	threadId, err := strconv.Atoi(c.Param("thread"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "", gin.H{})
	}
	for i, b := range boards.Boards {
		if b.Name == boardName {
			for ii, t := range b.Threads {
				if t.Id == threadId {
					boards.Boards[i].Threads[ii].Posts = append(t.Posts, newPost)
					returnThread = boards.Boards[i].Threads[ii]
					boards.Boards[i].Threads[ii].Position = globalId
					break
				}

			}
			break
		}
	}

	c.HTML(http.StatusOK, "thread.html", gin.H{
		"boardList": boardList,
		"boardName": boardName,
		"thread":    returnThread,
	})
}

func removeThread(c *gin.Context) {
	var returnBoard Board
	boardName := c.Param("board")
	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "", gin.H{})
	}
	seshCookie, err := c.Cookie("sesh_cook")
	showAdminTools := false
	if seshCookie == sessionId {
		showAdminTools = true
	}
	for i, b := range boards.Boards {
		for ii, t := range b.Threads {
			if t.Id == Id {
				boardThreads := append(boards.Boards[i].Threads[:ii], boards.Boards[i].Threads[ii+1:]...)
				boards.Boards[i].Threads = boardThreads
				break
			}

		}
	}
	c.HTML(http.StatusOK, "board.html", gin.H{
		"boardList":      boardList,
		"board":          returnBoard,
		"boardName":      boardName,
		"showAdminTools": showAdminTools,
	})
}

func removePost(c *gin.Context) {
	var returnThread Thread
	boardName := c.Param("board")

	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "", gin.H{})
	}
	seshCookie, err := c.Cookie("sesh_cook")
	showAdminTools := false
	if seshCookie == sessionId {
		showAdminTools = true
	}
	for i, b := range boards.Boards {
		for ii, t := range b.Threads {
			for iii, p := range t.Posts {
				if p.Id == Id {
					threadPosts := append(boards.Boards[i].Threads[ii].Posts[:iii], boards.Boards[i].Threads[ii].Posts[ii+1:]...)
					boards.Boards[i].Threads[ii].Posts = threadPosts
					returnThread = boards.Boards[i].Threads[ii]
					break
				}
			}
		}
	}
	c.HTML(http.StatusOK, "thread.html", gin.H{
		"boardList":      boardList,
		"boardName":      boardName,
		"thread":         returnThread,
		"showAdminTools": showAdminTools,
	})
}

var sessionId = uuid.NewString() // this probably needs to get saved as a global to check and updated when the password is entered in
var boards = &Boards{
	Boards: []Board{
		{
			Name:    "programming",
			Threads: []Thread{},
		},
		{
			Name:    "kcdc2024",
			Threads: []Thread{},
		},
	},
}

// var mockBoards = &Boards{
// Boards: []Board{
// {
// Name: "programming",
// Threads: []Thread{
// {
// Id:   1,
// Name: "first thread",
// Posts: []Post{
// {
// Name: "first post",
// Id:   3,
// Body: "this is the text body",
// },
// {
// Name: "second post",
// Id:   5,
// Body: "this is the text body",
// },
// },
// Body: "first thread body",
// },
// {
// Id:   2,
// Name: "second thread",
// Posts: []Post{
// {
// Name: "second thread, first post",
// Id:   4,
// Body: "here is another post",
// },
// {
// Name: "second thread, second post",
// Id:   6,
// Body: "here is another other post",
// },
// },
// Body: "second thread body",
// },
// },
// },
// },
// }
