const m = require('mithril')
import session from '../session'

const avatarView = {
  view: () => {
    return m("figure.avatar", [m("img", {src: session.currentUser.avatar})])
  }
}

const ProfileView = {
  oninit: () => {
    session.load(localStorage.accessToken)
  },
	view: () => {
    if (session.isSignedIn()) {
		  return m("section.navbar-section", {id: "profile-menu"}, m("div.form-group", {}, [
          m("a.btn.btn-primary", {onclick: (e) => { }, href: '#'}, "+"),
          m("figure.avatar", [m("img", {src: session.currentUser.avatar})])
      ]))
    } else {
		  return m("section.navbar-section", {id: "profile-menu"}, m("a.btn.btn-primary", {onclick: (e) => { lock.show(); e.preventDefault()  }, href: "#"}, "Login"))
    }
	}
}

export {ProfileView as default}
