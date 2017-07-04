const m = require('mithril')
import session from '../../session'
import SearchView from './search'

const avatarView = {
  view: () => {
    return m("figure.avatar", [m("img", {src: session.currentUser.avatar})])
  }
}

const ProfileView = {
  oninit: () => {
    session.load(localStorage.accessToken)
  },
	view: (vnode) => {
    if (session.isSignedIn()) {
		  return m("section.navbar-section", {id: "profile-menu"}, m("div.form-group.dropdown.dropdown-right", [
          m("a.dropdown-toggle", {tabindex: 0}, [
            m("figure.avatar", m("img", {src: session.currentUser.avatar})),
            m('i.icon.icon-caret')
          ]),
          m('ul.menu', [
            m('li.menu-item', m('a', {href: '#!/mylinks'}, 'My links')),
            m('li.menu-item', m('a', {href: '#!/collections'}, 'My collections')),
            m('li.divider'),
            m('li.menu-item', m('a', {href: '#', onclick: session.logout }, 'Logout'))
          ])
      ]))
    } else {
		  return m("section.navbar-section", {id: "profile-menu"}, m("a.btn.btn-primary", {onclick: session.login, href: "#"}, "Login"))
    }
	}
}

const NavView = {
  view: (vnode) => {
    return m('section.section.section-header.bg-gray', m('section.grid-header.container.grid-960', m('nav.navbar', [
        m("section.navbar-section", [
            m("a.btn.btn-lg.btn-link.btn-action.show-sm", {href: "#sidebar"}, m("i.icon.icon-menu")),
            m("a.navbar-brand.mr-10", {href: '#'}, 'BetterDev')]),
        m("section.navbar-section", m(SearchView)),
        m(ProfileView)
      ])))
  }
}

export {NavView as default}
