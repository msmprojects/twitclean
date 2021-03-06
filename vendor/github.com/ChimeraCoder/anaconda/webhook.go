package anaconda

import (
	"net/url"
)

//GetActivityWebhooks represents the twitter account_activity webhook
//Returns all URLs and their statuses for the given app. Currently,
//only one webhook URL can be registered to an application.
//https://developer.twitter.com/en/docs/accounts-and-users/subscribe-account-activity/api-reference/get-webhook-config
func (a TwitterApi) GetActivityWebhooks(v url.Values) (u []WebHookResp, err error) {
	responseCh := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/account_activity/webhooks.json", v, &u, _GET, responseCh}
	return u, (<-responseCh).err
}

//WebHookResp represents the Get webhook responses
type WebHookResp struct {
	ID        string
	URL       string
	Valid     bool
	CreatedAt string
}

//SetActivityWebhooks represents to set twitter account_activity webhook
//Registers a new webhook URL for the given application context.
//The URL will be validated via CRC request before saving. In case the validation fails,
//a comprehensive error is returned. message to the requester.
//Only one webhook URL can be registered to an application.
//https://developer.twitter.com/en/docs/accounts-and-users/subscribe-account-activity/api-reference/new-webhook-config
func (a TwitterApi) SetActivityWebhooks(v url.Values) (u []WebHookResp, err error) {
	responseCh := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/account_activity/webhooks.json", v, &u, _POST, responseCh}
	return u, (<-responseCh).err
}

//DeleteActivityWebhooks Removes the webhook from the provided application’s configuration.
//https://developer.twitter.com/en/docs/accounts-and-users/subscribe-account-activity/api-reference/delete-webhook-config
func (a TwitterApi) DeleteActivityWebhooks(v url.Values, webhookID string) (u interface{}, err error) {
	responseCh := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/account_activity/webhooks/" + webhookID + ".json", v, &u, _DELETE, responseCh}
	return u, (<-responseCh).err
}

//PutActivityWebhooks update webhook which reenables the webhook by setting its status to valid.
//https://developer.twitter.com/en/docs/accounts-and-users/subscribe-account-activity/api-reference/validate-webhook-config
func (a TwitterApi) PutActivityWebhooks(v url.Values, webhookID string) (u interface{}, err error) {
	responseCh := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/account_activity/webhooks/" + webhookID + ".json", v, &u, _PUT, responseCh}
	return u, (<-responseCh).err
}

//SetWHSubscription Subscribes the provided app to events for the provided user context.
//When subscribed, all DM events for the provided user will be sent to the app’s webhook via POST request.
//https://developer.twitter.com/en/docs/accounts-and-users/subscribe-account-activity/api-reference/new-subscription
func (a TwitterApi) SetWHSubscription(v url.Values, webhookID string) (u interface{}, err error) {
	responseCh := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/account_activity/webhooks/" + webhookID + "/subscriptions.json", v, &u, _POST, responseCh}
	return u, (<-responseCh).err
}

//GetWHSubscription Provides a way to determine if a webhook configuration is
//subscribed to the provided user’s Direct Messages.
//https://developer.twitter.com/en/docs/accounts-and-users/subscribe-account-activity/api-reference/get-subscription
func (a TwitterApi) GetWHSubscription(v url.Values, webhookID string) (u interface{}, err error) {
	responseCh := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/account_activity/webhooks/" + webhookID + "/subscriptions.json", v, &u, _GET, responseCh}
	return u, (<-responseCh).err
}

//DeleteWHSubscription Deactivates subscription for the provided user context and app. After deactivation,
//all DM events for the requesting user will no longer be sent to the webhook URL..
//https://developer.twitter.com/en/docs/accounts-and-users/subscribe-account-activity/api-reference/delete-subscription
func (a TwitterApi) DeleteWHSubscription(v url.Values, webhookID string) (u interface{}, err error) {
	responseCh := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/account_activity/webhooks/" + webhookID + "/subscriptions.json", v, &u, _DELETE, responseCh}
	return u, (<-responseCh).err
}
