package server

import (
	"net/http"
	"strconv"
	"template/errors"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetReviewByid(c *gin.Context) {
	reviewIDStr := c.Param("id")
	if reviewIDStr == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "Review ID not provided"})
		return
	}
	reviewID, err := strconv.Atoi(reviewIDStr)
	if err != nil {
		errors.BadRequest(c, err)
		return
	}

	// Call the GetReviewByid method from the API
	review, err := s.api.GetReviewByid(c, reviewID)
	if err != nil {
		errors.InternalServerError(c, "Failed to get review")
		return
	}

	c.JSON(http.StatusOK, gin.H{"review": review})
}

// Define the GetReviewAPI method for the Server struct
func (s *Server) GetReviewAPI(c *gin.Context) {
	reviews, err := s.api.GetReviewAPI(c)
	if err != nil {
		errors.BadRequest(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}
