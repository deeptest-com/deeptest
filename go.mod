module github.com/deeptest-com/deeptest

go 1.22.0

require (
	github.com/fullstorydev/grpcurl v1.9.1
	github.com/jhump/protoreflect v1.17.0
	google.golang.org/grpc v1.67.0
)

replace (
	github.com/fullstorydev/grpcurl v1.9.1 => github.com/fullstorydev/grpcurl v1.3.2
	github.com/jhump/protoreflect v1.16.0 => github.com/jhump/protoreflect v1.5.0
	google.golang.org/grpc v1.63.2 => google.golang.org/grpc v1.21.0
)

require (
	github.com/474420502/requests v1.35.0
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible
	github.com/andybalholm/brotli v1.1.0
	github.com/antchfx/htmlquery v1.3.0
	github.com/antchfx/jsonquery v1.3.0
	github.com/antchfx/xmlquery v1.3.17
	github.com/antchfx/xpath v1.2.4
	github.com/beevik/etree v1.2.0
	github.com/bndr/gojenkins v1.1.0
	github.com/casbin/casbin/v2 v2.37.2
	github.com/casbin/gorm-adapter/v3 v3.4.2
	github.com/clbanning/mxj v1.8.4
	github.com/d5/tengo/v2 v2.13.0
	github.com/dlclark/regexp2 v1.11.0
	github.com/dop251/goja v0.0.0-20240516125602-ccbae20bcec2
	github.com/dop251/goja_nodejs v0.0.0-20240418154818-2aae10d4cbcf
	github.com/emirpasic/gods v1.12.0
	github.com/facebookgo/inject v0.0.0-20180706035515-f23751cae28b
	github.com/fatih/color v1.17.0
	github.com/fsnotify/fsnotify v1.7.0
	github.com/getkin/kin-openapi v0.107.0
	github.com/go-ldap/ldap v3.0.3+incompatible
	github.com/go-ozzo/ozzo-routing/v2 v2.4.0
	github.com/go-playground/validator/v10 v10.9.0
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-sourcemap/sourcemap v2.1.4+incompatible
	github.com/go-sql-driver/mysql v1.7.1
	github.com/goccy/go-json v0.10.2
	github.com/gofrs/uuid v4.2.0+incompatible
	github.com/gookit/color v1.4.2
	github.com/gorilla/websocket v1.5.1
	github.com/invopop/yaml v0.2.0
	github.com/iris-contrib/middleware/cors v0.0.0-20210110101738-6d0a4d799b5d
	github.com/iris-contrib/swagger v0.0.0-20210110101208-2670b336b278
	github.com/jameskeane/bcrypt v0.0.0-20120420032655-c3cd44c1e20f
	github.com/jinzhu/copier v0.3.5
	github.com/kataras/iris/v12 v12.2.11
	github.com/kataras/neffos v0.0.24-0.20240408172741-99c879ba0ede
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lib/pq v1.10.9
	github.com/lucasjones/reggen v0.0.0-20200904144131-37ba4fa293bb
	github.com/mholt/archiver/v3 v3.5.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/muonsoft/language v0.3.1
	github.com/muonsoft/openapi-mock v0.3.9
	github.com/oklog/ulid/v2 v2.1.0
	github.com/oliveagle/jsonpath v0.0.0-20180606110733-2e52cf6e6852
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/rs/zerolog v1.31.0
	github.com/satori/go.uuid v1.2.0
	github.com/savsgio/gotils v0.0.0-20210617111740-97865ed5a873
	github.com/sirupsen/logrus v1.9.3
	github.com/snowlyg/helper v0.0.0-20210919081922-08e1fd247338
	github.com/snowlyg/multi v0.0.0-20210828154328-14d6bfb7d0a6
	github.com/spf13/viper v1.19.0
	github.com/stretchr/testify v1.9.0
	github.com/swaggo/swag v1.16.3
	github.com/tidwall/gjson v1.17.0
	github.com/tidwall/sjson v1.2.5
	github.com/xuri/excelize/v2 v2.6.1
	github.com/zwgblue/yaml-encoder v0.0.0-20221226083717-a0bdbda0d998
	go.uber.org/zap v1.21.0
	golang.org/x/crypto v0.27.0
	golang.org/x/net v0.29.0
	golang.org/x/text v0.18.0
	google.golang.org/protobuf v1.34.2
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/mysql v1.4.5
	gorm.io/driver/postgres v1.4.5
	gorm.io/driver/sqlite v1.4.4
	gorm.io/driver/sqlserver v1.4.1
	gorm.io/gorm v1.24.1-0.20221019064659-5dd2bb482755
	gorm.io/plugin/dbresolver v1.4.0
	syreclabs.com/go/faker v1.2.3
)

require (
	cloud.google.com/go v0.113.0 // indirect
	cloud.google.com/go/auth v0.4.1 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.2 // indirect
	cloud.google.com/go/compute/metadata v0.5.0 // indirect
	cloud.google.com/go/firestore v1.15.0 // indirect
	cloud.google.com/go/longrunning v0.5.7 // indirect
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/CloudyKit/jet/v6 v6.2.0 // indirect
	github.com/Joker/jade v1.1.3 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/Shopify/goreferrer v0.0.0-20220729165902-8cddb4f5de06 // indirect
	github.com/TylerBrock/colorjson v0.0.0-20200706003622-8a50f05110d2 // indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/bufbuild/protocompile v0.14.1 // indirect
	github.com/bwmarrin/snowflake v0.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/chindeo/pkg v0.0.2-0.20210813090029-854fe1390d17 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0-20190314233015-f79a8a8ca69d // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dsnet/compress v0.0.2-0.20210315054119-f66993602bf5 // indirect
	github.com/eknkc/amber v0.0.0-20171010120322-cdade1c07385 // indirect
	github.com/facebookgo/ensure v0.0.0-20200202191622-63f1cf65ac4c // indirect
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/facebookgo/structtag v0.0.0-20150214074306-217e25fb9691 // indirect
	github.com/facebookgo/subset v0.0.0-20200203212716-c811ad88dec4 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/flosch/pongo2/v4 v4.0.2 // indirect
	github.com/gavv/httpexpect/v2 v2.16.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/spec v0.20.9 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.3.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/golang/gddo v0.0.0-20210115222349-20d68f94ee1f // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gomarkdown/markdown v0.0.0-20240328165702-4d01890c35c0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/pprof v0.0.0-20240521024322-9665fa269a30 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.4 // indirect
	github.com/gorilla/css v1.0.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hashicorp/consul/api v1.28.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/imkira/go-interpol v1.1.0 // indirect
	github.com/iris-contrib/go.uuid v2.0.0+incompatible // indirect
	github.com/iris-contrib/schema v0.0.6 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.14.3 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.3 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/jackc/pgx/v4 v4.17.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kataras/blocks v0.0.8 // indirect
	github.com/kataras/golog v0.1.11 // indirect
	github.com/kataras/pio v0.0.13 // indirect
	github.com/kataras/sitemap v0.0.6 // indirect
	github.com/kataras/tunnel v0.0.4 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/klauspost/pgzip v1.2.4 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mailgun/raymond/v2 v2.0.48 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	github.com/mediocregopher/radix/v3 v3.8.1 // indirect
	github.com/microcosm-cc/bluemonday v1.0.26 // indirect
	github.com/microsoft/go-mssqldb v0.17.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/nats-io/nats.go v1.34.1 // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/nwaples/rardecode v1.1.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml v1.9.3 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/pierrec/lz4/v4 v4.0.3 // indirect
	github.com/pkoukk/tiktoken-go v0.1.6 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/r3labs/sse/v2 v2.10.0 // indirect
	github.com/richardlehane/mscfb v1.0.4 // indirect
	github.com/richardlehane/msoleps v1.0.3 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sagikazarmark/crypt v0.19.0 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sanity-io/litter v1.5.5 // indirect
	github.com/schollz/closestmatch v2.1.0+incompatible // indirect
	github.com/schollz/progressbar v1.0.0 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/tdewolff/minify/v2 v2.20.19 // indirect
	github.com/tdewolff/parse/v2 v2.7.12 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tmc/langchaingo v0.1.12 // indirect
	github.com/ulikunitz/xz v0.5.8 // indirect
	github.com/urfave/cli/v2 v2.3.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.56.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	github.com/xo/terminfo v0.0.0-20210125001918-ca9a967f8778 // indirect
	github.com/xuri/efp v0.0.0-20220603152613-6918739fd470 // indirect
	github.com/xuri/nfp v0.0.0-20220409054826-5e722a1d9e22 // indirect
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0 // indirect
	github.com/yosssi/ace v0.0.5 // indirect
	github.com/yudai/gojsondiff v1.0.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	go.etcd.io/etcd/api/v3 v3.5.12 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.12 // indirect
	go.etcd.io/etcd/client/v2 v2.305.12 // indirect
	go.etcd.io/etcd/client/v3 v3.5.12 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.51.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.51.0 // indirect
	go.opentelemetry.io/otel v1.26.0 // indirect
	go.opentelemetry.io/otel/metric v1.26.0 // indirect
	go.opentelemetry.io/otel/trace v1.26.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/exp v0.0.0-20240404231335-c0f41cb1a7a0 // indirect
	golang.org/x/oauth2 v0.22.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d // indirect
	golang.org/x/xerrors v0.0.0-20231012003039-104605ab7028 // indirect
	google.golang.org/api v0.180.0 // indirect
	google.golang.org/genproto v0.0.0-20240401170217-c3f982113cda // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240924160255-9d4c2d233b61 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/cenkalti/backoff.v1 v1.1.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	moul.io/http2curl v1.0.1-0.20190925090545-5cd742060b0e // indirect
	moul.io/http2curl/v2 v2.3.0 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)
