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
import AppView from './app'

class CollectionView extends AppView {
  oninit() {
    Collection.loadList()
  }

  view() {
    return [
      m(NavView, {Post: Post}),
      m('main.container.grid-960', m("div.columns", [m(FlashView)].concat(
        m(ErrorView),
        m("section.container.post-form", m(PostFormView)),
        m("section", m('ul.menu', Collection.list.map((c) => {
          return m("li.menu-item", m('a', {href: '#!/collections/' + c.id}, c.name))
        })))
      )))
    ]
  }
}
export { CollectionView as default }
