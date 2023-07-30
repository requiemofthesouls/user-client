package userclient

const (
	// NGINX заголовки
	HeaderRequestID    = "X-Request-ID"
	HeaderUserIP       = "X-Forwarded-For"
	HeaderUserLocation = "X-geoip-city-country-code3"
	HeaderUserHost     = "X-Forwarded-Host"

	// Клиентские заголовки
	HeaderUserAgent           = "X-User-Agent"
	HeaderUserLanguage        = "X-User-Language"
	HeaderApplicationPlatform = "X-Application-Platform"

	// Заголовки для авторизации
	HeaderAuthorization     = "Authorization"
	HeaderAuthorizationType = "Bearer"

	// Заголовки для внутреннего тестирования
	HeaderDebugUserID  = "X-Debug-User-ID"
	HeaderDebugCountry = "X-Debug-Country"
)
