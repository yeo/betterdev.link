const m = require('mithril')

import session from '../session'
import lock from '../util/lock'
import Post from '../model/Post'
import Collection from '../model/Collection'
import ErrorView from './error'
import FlashView from './shared/flash'
import PostFormView from './shared/postform'
import Pagination from './shared/pagination'
import NavView from './shared/nav'

class AddToCollectionView {
  static view () {
    return m(".dropdown", [
      m("a.btn.btn-link.dropdown-toggle", {tabindex: 0, href: '#'}, ["Add to collection", m('i.icon.icon-caret')]),
      m('ul.menu', Collection.list.map((c) => {
        return m('li.menu-item', m('a', c.name))
      }))
    ])
  }
}

const PostlistView = {
  oninit: () => {
    Post.loadList()
    //TODO: Collection load should go to app init states
    Collection.loadList()
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
            m('button.btn.btn-primary', {onclick: (e) => {
              Post.postStatus = "AddToCollectionCompose"
              Post.draftItem = p
            }}, 'Add to Collection'),
            Post.postStatus == "AddToCollectionCompose" ? m(AddToCollectionView) : null
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
