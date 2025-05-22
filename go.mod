module github.com/cloudlink-omega/backend

go 1.24.1

replace (
	github.com/cloudlink-omega/accounts => ..\accounts
	github.com/cloudlink-omega/signaling => ..\signaling
	github.com/cloudlink-omega/storage => ..\storage
)

require (
	github.com/cloudlink-omega/accounts v0.0.0-00010101000000-000000000000
	github.com/cloudlink-omega/signaling v0.0.0-00010101000000-000000000000
	github.com/cloudlink-omega/storage v0.0.0-00010101000000-000000000000
	github.com/goccy/go-json v0.10.5
	github.com/gofiber/fiber/v2 v2.52.6
	github.com/gofiber/template/html/v2 v2.1.3
	github.com/joho/godotenv v1.5.1
	github.com/mileusna/useragent v1.3.5
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.26.1
)

require (
	cloud.google.com/go/compute/metadata v0.6.0 // indirect
	github.com/andybalholm/brotli v1.1.1 // indirect
	github.com/boombuler/barcode v1.0.2 // indirect
	github.com/chuckpreslar/emission v0.0.0-20170206194824-a7ddd980baf9 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/elithrar/simple-scrypt v1.3.0 // indirect
	github.com/fasthttp/websocket v1.5.12 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/gofiber/contrib/websocket v1.3.4 // indirect
	github.com/gofiber/template v1.8.3 // indirect
	github.com/gofiber/utils v1.1.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/mrz1836/go-sanitize v1.3.5 // indirect
	github.com/muka/peerjs-go v0.0.0-20240401061429-5b28944b9e4f // indirect
	github.com/oklog/ulid/v2 v2.1.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/philhofer/fwd v1.1.3-0.20240916144458-20a13a1f6b7c // indirect
	github.com/pion/datachannel v1.5.10 // indirect
	github.com/pion/dtls/v2 v2.2.12 // indirect
	github.com/pion/ice/v2 v2.3.37 // indirect
	github.com/pion/interceptor v0.1.37 // indirect
	github.com/pion/logging v0.2.3 // indirect
	github.com/pion/mdns v0.0.12 // indirect
	github.com/pion/randutil v0.1.0 // indirect
	github.com/pion/rtcp v1.2.15 // indirect
	github.com/pion/rtp v1.8.13 // indirect
	github.com/pion/sctp v1.8.38 // indirect
	github.com/pion/sdp/v3 v3.0.11 // indirect
	github.com/pion/srtp/v2 v2.0.20 // indirect
	github.com/pion/stun v0.6.1 // indirect
	github.com/pion/transport/v2 v2.2.10 // indirect
	github.com/pion/transport/v3 v3.0.7 // indirect
	github.com/pion/turn/v2 v2.1.6 // indirect
	github.com/pion/webrtc/v3 v3.3.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/pquerna/otp v1.4.0 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/savsgio/gotils v0.0.0-20250408102913-196191ec6287 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	github.com/tinylib/msgp v1.2.5 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.60.0 // indirect
	github.com/wlynxg/anet v0.0.5 // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/oauth2 v0.29.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/mail.v2 v2.3.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
