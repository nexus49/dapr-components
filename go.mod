module github.com/nexus49/dapr-components

go 1.13

replace k8s.io/client => github.com/kubernetes-client/go v0.0.0-20190928040339-c757968c4c36

require (
	github.com/dapr/components-contrib v0.0.0-20200226172056-467062ce38b5
	github.com/stretchr/testify v1.4.0
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
)
