package a

type ARequest struct { // want "struct ARequest is not required"
	ID int `json:"id"`
}

type BRequest struct { // OK
	ID int `json:"id" binding:"required"`
}

type CRequest struct { // OK
	ID int `json:"id" binding:"required,abc"`
}

type DRequest struct { // want "struct DRequest is not required"
	ID int `json:"id" binding:"abc"`
}

type ERequest struct { // OK
	ID  int `json:"id" binding:"required"`
	Foo int `json:"id2"`
}
