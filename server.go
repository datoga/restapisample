package restapisample

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Server handles the creation of a gin server.
type Server struct {
	repository Repository
}

//NewServer creates a new Server, using a Repository in it.
func NewServer(repository Repository) *Server {
	return &Server{repository: repository}
}

//Start configures and starts the repository.
func (s Server) Start() error {
	router := gin.Default()

	router.GET("/jobs/:id", func(c *gin.Context) {
		id := c.Param("id")

		if id == "" {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		job, found, err := s.repository.Get(c, id)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		if !found {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, job)
	})

	router.POST("/jobs", func(c *gin.Context) {
		var job Job

		if err := c.ShouldBindJSON(&job); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		id, err := s.repository.Create(c, job)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": id})
	})
	router.PUT("/jobs/:id", func(c *gin.Context) {
		var job Job

		if err := c.ShouldBindJSON(&job); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		id := c.Param("id")

		if id == "" {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		err := s.repository.Update(c, id, job)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusOK)
	})
	router.DELETE("/jobs/:id", func(c *gin.Context) {
		id := c.Param("id")

		if id == "" {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		err := s.repository.Delete(c, id)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusNoContent)
	})
	router.PATCH("/jobs/:id/:key/:value", func(c *gin.Context) {
		id := c.Param("id")
		key := c.Param("key")
		value := c.Param("value")

		if id == "" || key == "" || value == "" {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		job, found, err := s.repository.Get(c, id)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		if !found {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		switch key {
		case "company":
			job.Company = value
		case "title":
			job.Title = value
		case "description":
			job.Description = value
		default:
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("not supported param %v", key))
			return
		}

		err = s.repository.Update(c, id, job)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusOK)
	})

	return router.Run()
}
