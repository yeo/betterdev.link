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
  view() {
    return [
      m(NavView, {Post: Post}),
      m('main.container.grid-960', m("div.columns", [m(FlashView)].concat(
        m(ErrorView),
        m("section.container.post-form", m(PostFormView)),
        m("section", m("p", 'link'))
      )))
    ]
  }
}
export { CollectionView as default }
