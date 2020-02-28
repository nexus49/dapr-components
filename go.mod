module github.com/nexus49/dapr-components

go 1.13

replace k8s.io/client => github.com/kubernetes-client/go v0.0.0-20190928040339-c757968c4c36

require (
	github.com/coreos/go-oidc v2.1.0+incompatible
	github.com/dapr/components-contrib v0.0.0-20200226172056-467062ce38b5
	github.com/didip/tollbooth v4.0.2+incompatible
	github.com/fasthttp-contrib/sessions v0.0.0-20160905201309-74f6ac73d5d5
	github.com/google/uuid v1.1.1
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0
	github.com/valyala/fasthttp v1.6.0
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
)
