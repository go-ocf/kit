module github.com/plgd-dev/kit

go 1.13

require (
	cloud.google.com/go v0.54.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.4.9
	github.com/fxamacker/cbor/v2 v2.2.0
	github.com/go-acme/lego v2.7.2+incompatible
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.1
	github.com/gorilla/mux v1.7.4
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/jessevdk/go-flags v1.4.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/lestrrat-go/jwx v1.0.2
	github.com/miekg/dns v1.1.31 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pion/dtls/v2 v2.0.1-0.20200503085337-8e86b3a7d585
	github.com/plgd-dev/go-coap/v2 v2.1.4-0.20201201213140-b8c428d8fccf
	github.com/stretchr/testify v1.6.1
	github.com/ugorji/go/codec v1.1.7
	github.com/valyala/fasthttp v1.12.0
	go.uber.org/zap v1.15.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	google.golang.org/genproto v0.0.0-20200511104702-f5ebc3bea380
	google.golang.org/grpc v1.29.1
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.3.0
)

replace gopkg.in/yaml.v2 v2.2.8 => github.com/cizmazia/yaml v0.0.0-20200220134304-2008791f5454
