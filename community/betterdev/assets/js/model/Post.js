const m = require('mithril')

import session from '../session'

const Post = {
  list: [],
  loadList: () => {
    console.log(localStorage.accessToken)
    return m.request({
      method: "GET",
      url: "/api/links",
      withCredentials: true,
      headers: {
        Authorization: `Bearer ${localStorage.accessToken}`
      }
    })
    .then(function(result) {
      Post.list = result
    })
  }
}

Post.loadList()
export default Post
