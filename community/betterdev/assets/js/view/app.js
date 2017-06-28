const m = require('mithril')

import session from '../session'
import lock from '../util/lock'
import Post from '../model/Post'
import ErrorView from './error'
import FlashView from './shared/flash'
import PostFormView from './shared/postform'
import Pagination from './shared/pagination'
import NavView from './shared/nav'

const PostlistView = {
  oninit: () => {
    Post.loadList()
  },
  view: () => {
    return Post.list.map((p) => {
      return m("div.column.col-12", [
        m("div.tile", [
          m("div.tile-icon", m("figure.avatar.avatar-lg", m("img", {src: p.picture, style: "width: 48px; height: 48px;", width: 48, height: 48}))),
          m('div.tile-content', [
            m('h5.tile-title', m('a', {target: 'bank', href: p.uri}, p.title)),
            m('span', p.tags ? p.tags.map((tag) => m('label.chip', tag.tag)) : null),
            m('p.tile-subtitle', p.description)
          ]),
          m('div.tile-action', [
            m('a.btn.btn-primary', {target: 'blank', href: p.uri}, 'Visit→'),
            m('button.btn.btn-primary', '♥')
          ])
        ])
      ])
    })
  }
}

class AppView {
  view() {
    return [
      m(NavView, {Post: Post}),
      m('main.container.grid-960', m("div.columns", [m(FlashView)].concat(
        m(ErrorView),
        m("section.container.post-form", m(PostFormView)),
        m("section.container", m(PostlistView)),
        m("section.container", m(Pagination, {pages: Post.pagination}))
      )))
    ]
  }
}
export { AppView as default }
