import Auth0Lock from 'auth0-lock'
import session from '../session'

const lock = window.lock = new Auth0Lock('dN3hp6MDmXoLD412wEFkdQHjbXbSp6ZT', 'yeo.auth0.com', {
	auth: {
		redirect: false,
		//redirectUrl: 'http://127.0.0.1:4000',
		//responseType: 'token',
		responseType: 'id_token',
		params: {
			scope: 'openid email profile'
		}
	},
  theme: {
    logo: '/images/icon-s.png',
    primaryColor: '#f36c3d'
  },
  languageDictionary: {
    title: "Share, Save, Help!"
  },
})

lock.on("authenticated", (authResult) => {
  localStorage.setItem('accessToken', authResult.idToken)
  session.load(authResult.idToken)
  session.flash({message: "Login succesfully", type: 'success'})
})

export { lock as default }
