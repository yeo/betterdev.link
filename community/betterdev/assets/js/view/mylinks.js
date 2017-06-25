const m = require('mithril')
import session from '../session'
import lock from '../util/lock'
import Post from '../model/Post'
import ErrorView from './error'

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
            m('p.tile-title', m('a', {target: 'bank', href: p.uri}, p.title)),
            m('label.chip', 'Elixir'),
            m('label.chip', 'Pi'),
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

const Postform = {
  error: false,
  oninit: () => {
    //TOOD: store in localstorage and reload?
    Post.loadDraft()
  },

  createPost: ()  => {
    Post.create(Post.draft)
  },

  view: () => {
    return [
      m(ErrorView, {errors: Post.errors}),
      m('header.text-center', m('h4', "Share a link! Help other dev learn!")),
      m('div.input-group', [
        m('input.form-input.input-lg', {
          type: 'text', placeholder: 'http://awesome.link/', value: Post.draft,
          oninput: m.withAttr("value", function(value) {Post.draft = value})
        }),
        m('button.btn.btn-primary.btn-action.btn-lg', {onclick: (e) => Post.create(Post.draft)}, m('i.icon.icon-plus'))
      ]),
    ]
  }
}

const LoginRemind = {
  view: () => {
    return m('header.text-center', [
      m('h4', "Login to share links! Help people learn!"),
      m('a.btn.btn-primary', {}, 'Login')
    ])
  }
}

const AppView = {
  view: () => {
    return m("div.columns", [
      m("section.container.post-form", session.isSignedIn() ? m(Postform) : m(LoginRemind)),
      m("section.cotainer", m(PostlistView))
    ])
  }
}
export { MyLinksView as default }
