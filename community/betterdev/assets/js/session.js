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
        this.currentUser.avatar = `https://avatars1.githubusercontent.com/u/${this.currentUser.userId}?v=3&s=64`
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
      this.currentUser.id = data.id
    }).catch(e => {
      this.currentUser = {}
      localStorage.removeItem('accessToken')
    })
  }

  static isSignedIn() {
    return this.currentUser && this.currentUser.email
  }

  static clear() {
    this.currentUser = {}
    localStorage.removeItem('accessToken')
  }

  static logout(e) {
    e.preventDefault()
    session.clear()
    lock.logout({returnTo: "/"})
  }

  static login(e) {
    console.log('Login')
    e.preventDefault()
    lock.show()
  }
}

export { Session as default }
