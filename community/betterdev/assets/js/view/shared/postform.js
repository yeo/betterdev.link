const m = require('mithril')
import session from '../../session'
import Post from '../../model/Post'

const Postform = {
  oninit: () => {
    Post.loadDraft()
  },

  createPost: ()  => {
    if (Post.draft != "") {
      Post.create(Post.draft)
    }
  },

  view: () => {
    return [
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
      m('a.btn.btn-primary', {onclick: session.login}, 'Login'),
      m('footer.text-center', [
        m('h4', 'Or share with our bot'),
        m('p', [
          m('a.btn.btn-lg', {href: `https://telegram.me/${window.app.config.telegram.bot}`},'Telegram'),
          ' ',
          m('a.btn.btn-lg', {href: `https://slack.com/oauth/authorize?&client_id=${window.app.config.slack.client_id}&scope=bot,chat:write:bot,links:read,incoming-webhook&redirect_uri=http://127.0.0.1:4000/bot/slack`}, 'Slack')
        ])
      ]),

    ])
  }
}

const PostformView = {
  view: () => {
    return session.isSignedIn() ? m(Postform) : m(LoginRemind)
  }
}

export { PostformView as default }
