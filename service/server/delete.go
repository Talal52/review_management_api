package server

import (
	"fmt"
	"strconv"
	"template/errors"

	"github.com/gin-gonic/gin"
)

func (s *Server) DeleteReviewAPI(c *gin.Context) {
	reviewIDStr := c.Param("id")
	if reviewIDStr == "" {
		errors.BadRequest(c, fmt.Errorf("user does not exist"))
		return
	}
	reviewID, err := strconv.Atoi(reviewIDStr)
	if err != nil {
		errors.BadRequest(c, err)
		return
	}

	review, err := s.api.DeleteReviewAPI(c, reviewID)
	if err != nil {
		errors.BadRequest(c, err)
		return
	}

	errors.SuccessResponse(c, review)
}
