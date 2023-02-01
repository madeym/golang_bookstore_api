package request

type BookRatingRequest struct {
	BookID int    `json:"book_id" binding:"required"`
	Rating int    `json:"rating" binding:"required"`
	Review string `json:"review" binding:"required"`
}
