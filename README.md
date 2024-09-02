Opsgenie REST client
===========

![GitHub branch check runs](https://img.shields.io/github/check-runs/ylem-co/opsgenie-client/main?color=green)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ylem-co/opsgenie-client?color=black)
<a href="https://github.com/ylem-co/opsgenie-client?tab=Apache-2.0-1-ov-file">![Static Badge](https://img.shields.io/badge/license-Apache%202.0-black)</a>
<a href="https://ylem.co" target="_blank">![Static Badge](https://img.shields.io/badge/website-ylem.co-black)</a>
<a href="https://docs.datamin.io" target="_blank">![Static Badge](https://img.shields.io/badge/documentation-docs.datamin.io-black)</a>
<a href="https://join.slack.com/t/ylem-co/shared_invite/zt-2nawzl6h0-qqJ0j7Vx_AEHfnB45xJg2Q" target="_blank">![Static Badge](https://img.shields.io/badge/community-join%20Slack-black)</a>

How to use?
------

First of initiate the client with [GeniKey](https://support.atlassian.com/opsgenie/docs/create-a-default-api-integration/). E.g., in init

```go
func init() {
	opsgenieclient.Initiate(opsgenieclient.Config{
		ApiKey: "<your api key>",
	})
}
```

Now as the client is created, we can create an alert

```go
opg, _ := opsgenieclient.CreateInstance(ctx)
_ = opg.CreateAlert(request)
```
