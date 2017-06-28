const m = require('mithril')
import Post from '../model/Post'

const ErrorView = {
  view: (vnode) => {
    return Post.errors.map((e) => {
      return m('.toast.toast-primary.toast-warning', [
        m('button.btn.btn-clear.float-right', {onclick: () => { Post.errors = [] }}), e
      ])
    })
  }
}

export { ErrorView as default }
