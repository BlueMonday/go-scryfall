package scryfall

import "context"

// OAuthScope is the level of access.
type OAuthScope string

const (
	// OAuthScopeRead grants the ability to inspect data on a user's
	// account. No methods that change data will be allowed.
	OAuthScopeRead OAuthScope = "read"

	// OAuthScopeReadWrite grants full API access to a user's account. The
	// application will be able to use methods that update, delete, and add
	// account data on behalf of the user.
	OAuthScopeReadWrite OAuthScope = "read_write"

	// OAuthScopeEphemeral will grant access to the user's public account
	// information, and then revoke access immediately afterward.
	//
	// Useful for creating software such as polls or petitions that only
	// need to make sure that a unique and valid account is signing or voting
	// as a one-time action.
	OAuthScopeEphemeral OAuthScope = "ephemeral"
)

// Account represents a Scryfall account.
type Account struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	DisplayName  string `json:"display_name"`
	Twitter      string `json:"twitter"`
	FullFeatured bool   `json:"full_featured"`
	Verified     bool   `json:"verified"`
}

// Application represents a Scryfall application.
type Application struct {
	ClientID     string `json:"client_id"`
	Name         string `json:"name"`
	HomepageURI  string `json:"homepage_uri"`
	ContactURI   string `json:"contact_uri"`
	ContactEmail string `json:"contact_email"`
}

// OAuthGrant is an OAuth grant.
type OAuthGrant struct {
	GrantID     string     `json:"grant_id"`
	CreatedAt   Timestamp  `json:"created_at"`
	Scope       OAuthScope `json:"scope"`
	GrantSecret string     `json:"grant_secret"`
	Revoked     bool       `json:"revoked"`
	Account     Account    `json:"account"`
}

// GetAccount returns an object describing the currently authenticated Scryfall
// account.
//
// Requires an OAuth grant with OAuthScopeRead or higher.
func (c *Client) GetAccount(ctx context.Context) (Account, error) {
	account := Account{}
	err := c.get(ctx, "account", &account)
	if err != nil {
		return Account{}, err
	}

	return account, nil
}

// GetApplication returns an object describing the currently authenticated
// application.
//
// Requires application authentication.
func (c *Client) GetApplication(ctx context.Context) (Application, error) {
	application := Application{}
	err := c.get(ctx, "application", &application)
	if err != nil {
		return Application{}, err
	}

	return application, nil
}

// OAuthConvertRequest is an OAuth convert request.
type OAuthConvertRequest struct {
	Code string `json:"code"`
}

// OAuthConvert exchanges an OAuth code for a full OAuth grant object. The
// returned object will contain the GrantSecret that you should use for future
// requests inside the Authorization header for that account.
//
// Each code expires in 5 minutes, and can only be used once. Repeated requests
// sent to this method with the same code will fail.
//
// Ensure that you save both the GrantID and the GrantSecret you receive, as
// well as recording any other data in the object that your application needs to
// run, such as the information inside the account object.
//
// Requires application authentication.
func (c *Client) OAuthConvert(ctx context.Context, code string) (OAuthGrant, error) {
	oAuthConvertRequest := OAuthConvertRequest{
		Code: code,
	}
	var oAuthGrant OAuthGrant
	err := c.post(ctx, "oauth/convert", &oAuthConvertRequest, &oAuthGrant)
	if err != nil {
		return OAuthGrant{}, err
	}

	return oAuthGrant, nil
}

// OAuthDowngradeRequest is an OAuth downgrade request.
type OAuthDowngradeRequest struct {
	GrantID string `json:"grant_id"`
}

// OAuthDowngrade downgrades the scope of the OAuth grant identified by the
// submitted grantID.
//
// If the scope of the grant is OAuthScopeReadWrite, it will change to
// OAuthScopeRead. If the scope was already OAuthScopeRead, the grant object is
// returned unchanged (OAuthScopeRead is the lowest permission scope).
//
// Downgraded grants cannot be upgraded later, this change is permanent.
//
// This method is designed to allow your application to proactively relinquish
// rights to a user's account if you no longer need OAuthScopeReadWrite scope.
//
// Requires application authentication.
func (c *Client) OAuthDowngrade(ctx context.Context, grantID string) (OAuthGrant, error) {
	oAuthDowngradeRequest := OAuthDowngradeRequest{
		GrantID: grantID,
	}
	var oAuthGrant OAuthGrant
	err := c.post(ctx, "oauth/downgrade", &oAuthDowngradeRequest, &oAuthGrant)
	if err != nil {
		return OAuthGrant{}, err
	}

	return oAuthGrant, nil
}

// OAuthRevokeRequest is an OAuth revoke request.
type OAuthRevokeRequest struct {
	GrantID string `json:"grant_id"`
}

// OAuthRevokeResponse is an OAuth revoke response.
type OAuthRevokeResponse struct {
	GrantID   string    `json:"grant_id"`
	CreatedAt Timestamp `json:"created_at"`
	Revoked   bool      `json:"revoked"`
}

// OAuthRevoke revokes the OAuth grant identified by the provided grantID. The
// entire grant is immediately invalidated, no further requests may be made
// using its grant ID or grant secret. A minimal revoked version of the grant is
// returned as confirmation.
//
// The user must perform the full OAuth flow to establish a new grant with your
// application if they so desire.
//
// This method is designed to allow your application to proactively disconnect
// from this user's Scryfall account from your side.
//
// Requires application authentication.
func (c *Client) OAuthRevoke(ctx context.Context, grantID string) (OAuthRevokeResponse, error) {
	oAuthRevokeRequest := OAuthRevokeRequest{
		GrantID: grantID,
	}
	var oAuthRevokeResponse OAuthRevokeResponse
	err := c.post(ctx, "oauth/revoke", &oAuthRevokeRequest, &oAuthRevokeResponse)
	if err != nil {
		return OAuthRevokeResponse{}, err
	}

	return oAuthRevokeResponse, nil
}
