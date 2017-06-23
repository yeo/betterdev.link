import lock from './util/lock'
const m = require('mithril')

const jwtDecode = require('jwt-decode')

class Session {
  static load(token) {
    this.currentUser = {}
    if (!token) {
      return
    }

    const user = jwtDecode(token)
		console.log(user, '2')
    if (user.email) {
      this.currentUser = {
        email: user.email,
        provider: user.sub.split("|")[0],
        userId: user.sub.split("|")[1],
        nonce: user.nonce,
      }
      if (this.currentUser.provider == 'github') {
        this.currentUser.avatar = `https://avatars1.githubusercontent.com/u/${this.currentUser.userId}?v=3&s=64`
      } else {
      }
			this.validate()
		} else {
			localStorage.removeItem('accessToken')
    }
  }

	static validate() {
		m.request({
			method: "GET",
      url: "/api/me",
      withCredentials: true,
      headers: {
        Authorization: `Bearer ${localStorage.accessToken}`
      }
		}).then(data => {
			console.log(data)
		}).catch(e => console.log(e))
	}

  static isSignedIn() {
    return this.currentUser && this.currentUser.email
  }
}

export { Session as default }
