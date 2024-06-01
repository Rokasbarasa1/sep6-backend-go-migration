package types

type Movie struct {
	MovieId uint32 `json:"movieId" binding:"required,gte=1,lte=9999999"`
}
