package v2action

import (
	"fmt"

	"code.cloudfoundry.org/cli/api/uaa/constant"
)

// Authenticate authenticates the user in UAA and sets the returned tokens in
// the config.
//
// It unsets the currently targeted org and space whether authentication
// succeeds or not.
func (actor Actor) Authenticate(ID string, secret string, grantType constant.GrantType) error {
	actor.Config.UnsetOrganizationInformation()
	actor.Config.UnsetSpaceInformation()

	accessToken, refreshToken, err := actor.UAAClient.Authenticate(ID, secret, grantType)
	if err != nil {
		actor.Config.SetTokenInformation("", "", "")
		return err
	}

	accessToken = fmt.Sprintf("bearer %s", accessToken)
	actor.Config.SetTokenInformation(accessToken, refreshToken, "")
	return nil
}
