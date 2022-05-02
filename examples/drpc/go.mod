module storj.io/drpc/examples/drpc

go 1.17

require (
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/protobuf v1.28.0
	storj.io/drpc v0.0.30
)

require (
	github.com/zeebo/errs v1.2.2 // indirect
	golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect
)

replace storj.io/drpc => ../..
