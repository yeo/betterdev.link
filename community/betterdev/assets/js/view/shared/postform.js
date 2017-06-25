const m = require('mithril')
import session from '../../session'
import ErrorView from '../error'
import Post from '../../model/Post'

const Postform = {
  error: false,
  oninit: () => {
    //TOOD: store in localstorage and reload?
    Post.loadDraft()
  },

  createPost: ()  => {
    if (Post.draft != "") {
      Post.create(Post.draft)
    }
  },

  view: () => {
    return [
      m(ErrorView, {errors: Post.errors}),
      m('header.text-center', m('h4', "Share a link! Help other dev learn!")),
      m('div.input-group', [
        m('input.form-input.input-lg', {
          type: 'url', placeholder: 'http://awesome.link/', value: Post.draft,
          oninput: m.withAttr("value", function(value) {Post.draft = value})
        }),
        m('button.btn.btn-primary.btn-action.btn-lg' + (Post.postStatus == "posting" ? ".loading" : ''), {onclick: (e) => Post.create(Post.draft)}, m('i.icon.icon-plus'))
      ]),
    ]
  }
}

const LoginRemind = {
  view: () => {
    return m('header.text-center', [
      m('h4', "Login to share links! Help people learn!"),
      m('a.btn.btn-primary', {onclick: session.login}, 'Login with Github')
    ])
  }
}

const PostformView = {
  view: () => {
    return session.isSignedIn() ? m(Postform) : m(LoginRemind)
  }
}

export { PostformView as default }
