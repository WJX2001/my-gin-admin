package system

type ServiceGroup struct {
	JwtService
	UserService
	SecurityConfigService
	LoginLogService
}
