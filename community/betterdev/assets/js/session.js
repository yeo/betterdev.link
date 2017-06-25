import lock from './util/lock'
const m = require('mithril')

const jwtDecode = require('jwt-decode')

class Session {
  static flash(message) {
    if (message) {
      if (!Session._flash) {
        Session._flash = []
      }
      Session._flash.concat(message)
      return Session._flash
    }

    setTimeout(() => {
      Session._flash = []
    }, 10000)
    if (Session._flash) {
      return Session._flash
    }
    return []
  }


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
    Session.clear()
    lock.logout({returnTo: document.URL})
  }

  static login(e) {
    e.preventDefault()
    lock.show({autoclose: true})
  }
}

export { Session as default }
