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
  static view (vnode) {
    return m(".dropdown.dropdown-right", [

      m("a.btn.dropdown-toggle.btn-primary", {tabindex: 0, href: '#', onclick: (e) => {
        e.preventDefault()
        Post.postStatus = "AddToCollectionCompose"
        Post.draftItem = vnode.attrs.link
      }}, ["Save", m('i.icon.icon-caret')]),

      Post.postStatus == "AddToCollectionCompose" ?
        m('ul.menu', [
          Collection.list.map((c) => {
            return m('li.menu-item', [m('button.btn.btn-sm',{onclick: (e) => {
              Collection.append(vnode.attrs.link, c)
            }}, m('i.icon' + vnode.attrs.link.collections.indexOf(c.id) >= 0 ? '.icon-plus' : '.icon-minus')), ' ', c.name])
          }),
          m('li.divider'),
          m('li.menu-item', m('.input-group', [
            m('input.form-input[type=text][placeholder=New Collection]', { oninput: m.withAttr("value", function(value) { Collection.draftCollection = value })}),
            m('button.btn.btn-primary.input-group-btn', {onclick: Collection.create}, 'Create'),
          ]))
        ]) : null

    ])
  }

}

const PostlistView = {
  oninit: () => {
    console.log(m.route.param('user_id'))
    if (m.route.param('user_id')) {
      Post.params['user_id'] = m.route.param('user_id')
    } else {
      console.log("delete")
      delete Post.params.user_id
    }

    if (m.route.param('collection_id')) {
      Post.params['collection'] = m.route.param('collection_id')
    } else {
      delete Post.params.collection
    }

    Post.loadList()
    //TODO: Collection load should go to app init states
    Collection.loadList()
  },
  view: () => {
    return Post.list.map((p) => {
      return m("div.column.col-12", [
        m(".article",m('div.tile', [
          m("div.tile-icon", m("figure.avatar.avatar-lg", m("img", {src: p.picture.startsWith("https") ? p.picture : "/img?url=" + p.picture, style: "width: 48px; height: 48px;", width: 48, height: 48}))),
          m('div.tile-content', [
            m('h6.tile-title', m('a', {target: 'bank', href: p.uri}, p.title)),
            m('p.tile-subtitle.article-title', p.description),
            m('p.tile-subtitle', [m('a', {href: '#!/user_links/' + p.user.id, onclick: (e) => { Post.switchUser(p.user.id) }}, p.user.name), ' post at ', p.inserted_at, p.tags.length > 0 ? ' in ' : '', p.tags ? p.tags.map((tag) => m('span.label', tag.tag)) : null]),
          ]),
          m('.tile-action', [
            m(AddToCollectionView, {link: p}),
          ])
        ]))
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
