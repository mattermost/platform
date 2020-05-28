// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package api4

func (api *API) InitUserLocal() {
	api.BaseRoutes.User.Handle("/tokens", api.ApiLocal(createUserAccessToken)).Methods("POST")

	api.BaseRoutes.Users.Handle("/tokens/revoke", api.ApiLocal(revokeUserAccessToken)).Methods("POST")
}
