const m = require('mithril')
import Post from '../../model/Post'

const Pagination = {
  view: (vnode) => {
    let pages = vnode.attrs.pages
    if (pages.length < 2) {
      return null;
    }

    if (pages[0] && pages[0].page == 1) {
      pages = [{"url":"#","page":1,"label":"1","current": false, disable: true}].concat(pages)
    }

    return m("ul.pagination", vnode.attrs.pages.map(page => {
      return m('li.page-item' + (page.current ? '.active' : '') + (page.disable ? '.disable' : ''),
        m('a', {href: page.url, onclick: (e) => { e.preventDefault(); Post.params.page = page.page; Post.loadList(); }}, page.label))
    }))
  }
}

export { Pagination as default }
