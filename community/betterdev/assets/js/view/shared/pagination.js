const m = require('mithril')
import Post from '../../model/Post'

const Pagination = {
  view: (vnode) => {
    return m("ul.pagination", vnode.attrs.pages.map(page => {
      return m('li.page-item' + (page.current ? '.active' : ''),
        m('a', {href: page.url, onclick: (e) => { e.preventDefault(); Post.params.page = page.page; Post.loadList(); }}, page.label))
    }))
  }
}

export { Pagination as default }
