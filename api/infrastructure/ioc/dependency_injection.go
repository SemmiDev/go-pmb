package ioc

import (
	"github.com/SemmiDev/fiber-go-clean-arch/api/application/services"
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/interfaces"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/adapters"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/bus"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/database"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/environments"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/payments/midtrans"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/repositories"
	"github.com/SemmiDev/fiber-go-clean-arch/api/presentation/controllers"
)

var (
	MongoConnection *database.Mongo
	RedisConnection *database.RedisConnection

	HashAdapter adapters.IHashAdapter
	JwtAdapter  adapters.IJwtAdapter

	RabbitMQ interfaces.IRabbitMQ
	Midtrans midtrans.IMidtrans

	RegistrationRepository interfaces.IRegistrationRepository
	RedisAuthRepository    interfaces.IRedisAuthRepository
	RegistrationService    interfaces.IRegistrationService

	RegistrationController *controllers.RegistrationController
	AuthController         *controllers.AuthController
)

func SetupDependencyInjection(mc *database.Mongo, rc *database.RedisConnection) {
	// Database
	MongoConnection = mc
	RedisConnection = rc

	// Adapters
	HashAdapter = adapters.NewHashAdapter()
	JwtAdapter = adapters.NewJwtAdapter()

	// Bus
	RabbitMQ = bus.New(environments.RabbitMQConnection)

	// Payment
	Midtrans = midtrans.New()

	// Repositories
	RegistrationRepository = repositories.NewRegistrationRepository(MongoConnection)
	RedisAuthRepository = repositories.NewRedisAuthRepository(RedisConnection)

	// Services
	RegistrationService = services.NewRegistrationService(
		HashAdapter,
		Midtrans,
		RabbitMQ,
		RegistrationRepository,
	)
	AuthService := services.NewAuthService(
		RegistrationRepository,
		RedisAuthRepository,
		HashAdapter,
		JwtAdapter,
	)

	// Controllers
	RegistrationController = controllers.NewRegistrationController(RegistrationService)
	AuthController = controllers.NewAuthController(JwtAdapter, RedisAuthRepository, AuthService)
}
