module storj.io/drpc/examples/drpc

go 1.17

require (
<<<<<<< HEAD
	google.golang.org/protobuf v1.27.1
=======
	github.com/sirupsen/logrus v1.8.1 // indirect
	google.golang.org/protobuf v1.26.0
>>>>>>> e70d817 (examples/drpc: add logging and some concurrency to the client)
	storj.io/drpc v0.0.17
)

require github.com/zeebo/errs v1.2.2 // indirect

replace storj.io/drpc => ../..
