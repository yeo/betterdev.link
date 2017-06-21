const m = require('mithril')

import session from '../session'

const Post = {
  list: [],
  loadList: () => {
    return m.request({
      method: "GET",
      url: "/api/links",
      withCredentials: true,
      headers: {
        Authorization: `Bearer ${localStorage.accessToken}`
      }
    })
    .then(function(result) {
      if (result && result.data) {
        Post.list = result.data
      }
    })
    .catch((e) => { console.log("Error", e) })
  },

  draft: "",
  loadDraft: () => {
    Post.draft = ""
  }
}

export default Post
