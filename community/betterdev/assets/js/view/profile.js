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
		  return m("section.navbar-section", {id: "profile-menu"}, m("div.form-group.dropdown", [
          m("a.dropdown-toggle", {tabindex: 0}, [
            m("figure.avatar", m("img", {src: session.currentUser.avatar})),
            m('i.icon.icon-caret')
          ]),
          m('ul.menu', [
            m('li.menu-item', m('a', {href: '#mylink'}, 'My links')),
            m('li.menu-item', m('a', {href: '#', onclick: session.logout }, 'Logout'))
          ])
      ]))
    } else {
		  return m("section.navbar-section", {id: "profile-menu"}, m("a.btn.btn-primary", {onclick: session.login, href: "#"}, "Login"))
    }
	}
}

export {ProfileView as default}
