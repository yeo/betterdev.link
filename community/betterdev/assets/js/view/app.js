const m = require('mithril')
import session from '../session'

const AppView = {
  oninit: () => {
  },
	view: () => {
    if (session.isSignedIn()) {
		  return m("section.navbar-section", {id: "profile-menu"},
                m("div.form-group", {}, [
                  m("a.btn.btn-primary", {onclick: (e) => { }, href: '#'}, "+"),
                  m("figure.avatar", {}, [ m("img", {src: session.currentUser.avatar})])
                ]))
    } else {
		  return m("section.navbar-section", {id: "profile-menu"}, m("a.btn.btn-primary", {onclick: show_lock, href: "#"}, "Login"))
    }
	}
}

export default AppView
