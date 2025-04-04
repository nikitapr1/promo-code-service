package dto

type Mode string

const (
	COMMON Mode = "COMMON"
	UNIQUE Mode = "UNIQUE"
)

type Target struct {
	ageFrom    uint32
	ageUntil   uint32
	country    string
	categories []string
}

type PromoDTO struct {
	promoId     string
	companyId   string
	companyName string
	likeCount   int64
	usedCount   int64
	active      bool
	mode        Mode
	promoCommon string
	promoUnique []string
	description string
	imageUrl    string
	target      Target
	maxCount    int64
	activeFrom  string
	activeUntil string
}
