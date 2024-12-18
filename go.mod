module github.com/cloudlink-omega/backend

go 1.23.3

replace (
	github.com/cloudlink-omega/accounts => ../accounts
	github.com/cloudlink-omega/signaling => ../signaling
)

require (
	github.com/cloudlink-omega/accounts v0.0.0-00010101000000-000000000000
	github.com/cloudlink-omega/signaling v0.0.0-00010101000000-000000000000
	github.com/cloudlink-omega/backend v0.0.0-20241015194227-80fc91596c4c
	github.com/elithrar/simple-scrypt v1.3.0
	github.com/go-chi/chi/v5 v5.2.0
	github.com/go-chi/cors v1.2.1
	github.com/go-playground/validator/v10 v10.22.1
	github.com/goccy/go-json v0.10.3
	github.com/gofiber/fiber/v2 v2.52.5
	github.com/gofiber/template/html/v2 v2.1.2
	github.com/gorilla/websocket v1.5.3
	github.com/huandu/go-sqlbuilder v1.32.0
	github.com/joho/godotenv v1.5.1
	github.com/ncruces/go-sqlite3 v0.21.2
	github.com/oklog/ulid/v2 v2.1.0
	gopkg.in/mail.v2 v2.3.1
)

require (
	cloud.google.com/go/compute/metadata v0.3.0 // indirect
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/fasthttp/websocket v1.5.8 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/gofiber/contrib/websocket v1.3.2 // indirect
	github.com/gofiber/template v1.8.3 // indirect
	github.com/gofiber/utils v1.1.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/huandu/xstrings v1.4.0 // indirect
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mrz1836/go-sanitize v1.3.3 // indirect
	github.com/ncruces/julianday v1.0.0 // indirect
	github.com/pion/datachannel v1.5.9 // indirect
	github.com/pion/dtls/v3 v3.0.3 // indirect
	github.com/pion/ice/v4 v4.0.2 // indirect
	github.com/pion/interceptor v0.1.37 // indirect
	github.com/pion/logging v0.2.2 // indirect
	github.com/pion/mdns/v2 v2.0.7 // indirect
	github.com/pion/randutil v0.1.0 // indirect
	github.com/pion/rtcp v1.2.14 // indirect
	github.com/pion/rtp v1.8.9 // indirect
	github.com/pion/sctp v1.8.33 // indirect
	github.com/pion/sdp/v3 v3.0.9 // indirect
	github.com/pion/srtp/v3 v3.0.4 // indirect
	github.com/pion/stun/v3 v3.0.0 // indirect
	github.com/pion/transport/v3 v3.0.7 // indirect
	github.com/pion/turn/v4 v4.0.0 // indirect
	github.com/pion/webrtc/v4 v4.0.1 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/savsgio/gotils v0.0.0-20240303185622-093b76447511 // indirect
	github.com/tetratelabs/wazero v1.8.2 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.52.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	github.com/wlynxg/anet v0.0.3 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/oauth2 v0.24.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
)
