const m = require('mithril')
import Post from '../model/Post'
import session from '../session'
import AppView from './app'

class MyLinksView extends AppView {
  oninit() {
    Post.params["user_id"] = session.currentUser.id
  }
}
export { MyLinksView as default }
