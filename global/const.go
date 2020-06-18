package global

const (
	dburi       = "mongodb+srv://jerry:<p4JKI2mSGhWL8XJj>@cluster0-j7nek.azure.mongodb.net/<dbname>?retryWrites=true&w=majority"
	dbname      = "grpc"
	performance = 100
)

var (
	jwtSecret = []byte("grpc-demo")
)
