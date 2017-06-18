// Brunch automatically concatenates all files in your
// watched paths. Those paths can be configured at
// config.paths.watched in "brunch-config.js".
//
// However, those files will only be executed if
// explicitly imported. The only exception are files
// in vendor, which are never wrapped in imports and
// therefore are always executed.

// Import dependencies
//
// If you no longer want to use a dependency, remember
// to also remove its path from "config.paths.watched".
import "phoenix_html"

// Import local files
//
// Local files can be imported directly using relative
// paths "./socket" or full ones "web/static/js/socket".

// import socket from "./socket"<
import Auth0Lock from 'auth0-lock'
const lock = window.lock = new Auth0Lock('a-wb-XWRmIfAm9v0U9eUbfawoJKsGG99', 'yeo.auth0.com', {
	auth: {
		redirect: false,
		//redirectUrl: 'http://127.0.0.1:4000',
		//responseType: 'token',
		responseType: 'id_token',
		params: {
			scope: 'openid email'
		}
	}
})

lock.on("authenticated", function(authResult) {
	localStorage.setItem('accessToken', authResult.idToken);
})
