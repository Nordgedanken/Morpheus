package scalar

import "github.com/matrix-org/gomatrix"

type OpenIDTokenResponse struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	MatrixServerName string `json:"matrix_server_name"`
}

func GetOpenIDToken(cli *gomatrix.Client) (resp *OpenIDTokenResponse, err error) {
	urlPath := cli.BuildURL("user", cli.UserID, "openid", "request_token")
	_, err = MakeRequest(cli, "POST", urlPath, "{}", &resp)
	return
}

type ScalarRegisterResp struct {
	ScalarToken string `json:"scalar_token"`
}

//GetScalarToken returns the scalar token for the current user. Used inside the Integration manager.
func GetScalarToken(cli *gomatrix.Client, openIDToken *OpenIDTokenResponse) (resp *ScalarRegisterResp, err error) {
	urlPath := "https://scalar.vector.im/api/register"
	_, err = cli.MakeRequest("POST", urlPath, openIDToken, &resp)
	return
}
