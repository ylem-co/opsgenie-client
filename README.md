Opsgenie REST client
===========

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
