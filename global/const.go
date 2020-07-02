package global

const (
	// DB_URI        = "mongodb+srv://jerry:p4JKI2mSGhWL8XJj@cluster0-j7nek.azure.mongodb.net/test?retryWrites=true&w=majority"
	DB_URI        = "mongodb://localhost:27017/"
	DB_NAME       = "grpc"
	DB_COLLECTION = "users"
)

var (
	JwtSecret = "grpc-demo"
)
