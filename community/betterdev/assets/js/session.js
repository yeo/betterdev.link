import lock from './util/lock'

const jwtDecode = require('jwt-decode')

class Session {
  static load(token) {
    this.currentUser = {}
    if (!token) {
      return
    }

    const user = jwtDecode(token)
    console.log(user)
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
    }
  }

  static isSignedIn() {
    return this.currentUser && this.currentUser.email
  }
}

export { Session as default }