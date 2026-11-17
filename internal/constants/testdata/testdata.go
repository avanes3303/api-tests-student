package testdata

const (
	AdvertisementPrice     = "1250.50"
	AdvertisementQuantity  = "2"
	AdvertisementText      = "Учебное объявление для проверки API"
	TitleLimit             = 50
	DescriptionLimit       = 500
	PhotoLimit             = 3
	PhotoWidth             = 32
	PhotoHeight            = 32
	FavoritesLimit         = 50
	InvalidToken           = "incorrect_token"
	InvalidID              = "not-a-uuid"
	MissingAdvertisementID = "00000000-0000-4000-8000-000000000001"

	InvalidRequest = "invalid_request"
	Unauthorized   = "unauthorized"
	Forbidden      = "forbidden"
	NotFound       = "not_found"
	AlreadyExists  = "already_exists"
	LimitReached   = "limit_reached"

	TitleRequired         = "title is required"
	DescriptionRequired   = "description is required"
	PhotoRequired         = "at least 1 photo is required"
	PriceRange            = "price must be between 0 and 1000000 inclusive"
	QuantityRange         = "quantity must be between 1 and 100 inclusive"
	AuthorizationRequired = "Authorization header required"
	InvalidAuthorization  = "Invalid or expired token"
	OwnFavorite           = "Cannot add own advertisement to favorites"
	DuplicateFavorite     = "Advertisement is already in favorites"
	AdvertisementNotFound = "Advertisement not found"
	FavoriteNotFound      = "Favorite entry not found"
	FavoriteIDRequired    = "Key: 'AddFavoriteRequest.AdvertisementID' Error:Field validation for 'AdvertisementID' failed on the 'required' tag"
	InvalidUUID           = "invalid UUID length: 10"
	FavoritesLimitReached = "Favorites limit of 50 reached"
)
