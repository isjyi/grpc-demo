package global

const (
	DB_URI        = "mongodb+srv://jerry:<p4JKI2mSGhWL8XJj>@cluster0-j7nek.azure.mongodb.net/<dbname>?retryWrites=true&w=majority"
	DB_NAME       = "grpc"
	DB_COLLECTION = "users"
)

var (
	JwtSecret = []byte("grpc-demo")
)
