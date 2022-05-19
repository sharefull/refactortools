package a

type ARequest struct {
	ID int `json:"id"` // want "field ID is not required"
}

type BRequest struct {
	ID int `json:"id" binding:"required"` // OK
}

type CRequest struct {
	ID int `json:"id" binding:"required,abc"` // OK
}

type DRequest struct {
	ID int `json:"id" binding:"abc"` // want "field ID is not required"
}
