package server

import (
	"template/errors"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetReviewAPI(c *gin.Context) {
	reviews, err := s.api.GetReviewAPI(c)
	if err != nil {
		errors.BadRequest(c, err)
		return
	}

	errors.SuccessResponse(c, reviews)
}

// func (s *Server) GetReviewByid(c *gin.Context){
// 	reviewIDStr:=c.Param("id")
// 	if reviewIDStr==""{
// 		errors.BadRequest(c,fmt.Errorf("user does not exist"))
// 		return
// 	}
// 	reviewID,err:=strconv.Atoi(reviewIDStr)
// 	if err!=nil{
// 		errors.BadRequest(c,err)
// 		return
// 	}
// 	review,err:=s.api.GetReviewByid(c,reviewID)
// 	if err!=nil{
// 		errors.BadRequest(c,err)
// 		return
// 	}
// 	errors.SuccessResponse(c,review)
// }
