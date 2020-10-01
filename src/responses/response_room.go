package responses

// RoomDetailResponse room response
type RoomDetailResponse struct {
	ID        int64  `json:"id" swaggertype:"integer" example:"1"`
	Name      string `json:"name" swaggertype:"string" example:"Lich Truong"`
	CreatedAt string `json:"created_at" swaggertype:"string" example:"1991-02-13 10:10:10"`
	UpdatedAt string `json:"updated_at" swaggertype:"string" example:"2020-07-15 10:10:10"`
}
