package tokenservice

const (
	AccessTokenExpireTime  = 1000
	RefreshTokenExpireTime = 1000

	UserID     = "user_id"
	Name       = "name"
	NationalID = "national_id"
	Phone      = "phone"
	Status     = "status"
)

type Service struct {
}

func New() Service {
	return Service{}
}
