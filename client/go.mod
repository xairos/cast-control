module github.com/xairos/cast-control/client

go 1.12

require (
	github.com/AlecAivazis/survey/v2 v2.0.0-20190604222524-692473aeb872
	github.com/OpenPeeDeeP/xdg v0.2.0
	github.com/alecthomas/template v0.0.0-20160405071501-a0175ee3bccc // indirect
	github.com/alecthomas/units v0.0.0-20151022065526-2efee857e7cf // indirect
	github.com/stretchr/testify v1.3.0 // indirect
	github.com/xairos/cast-control/client/protocol v0.0.0
	google.golang.org/grpc v1.21.1
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)

replace github.com/xairos/cast-control/client/protocol => ./protocol
