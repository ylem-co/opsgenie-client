Opsgenie REST client
===========

![GitHub branch check runs](https://img.shields.io/github/check-runs/datamin-io/opsgenie-client/main?color=green)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/datamin-io/opsgenie-client?color=blue)
<a href="https://github.com/datamin-io/opsgenie-client?tab=Apache-2.0-1-ov-file">![Static Badge](https://img.shields.io/badge/license-Apache%202.0-blue)</a>
<a href="https://datamin.io" target="_blank">![Static Badge](https://img.shields.io/badge/website-datamin.io-blue)</a>
<a href="https://docs.datamin.io" target="_blank">![Static Badge](https://img.shields.io/badge/documentation-docs.datamin.io-blue)</a>
<a href="https://join.slack.com/t/datamincommunity/shared_invite/zt-2nawzl6h0-qqJ0j7Vx_AEHfnB45xJg2Q" target="_blank">![Static Badge](https://img.shields.io/badge/community-join%20Slack-blue)</a>

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
