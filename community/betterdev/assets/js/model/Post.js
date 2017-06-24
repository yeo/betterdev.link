const m = require('mithril')

import session from '../session'

const Post = {
  errors: [],
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
  },

  create: (uri) => {
    return m.request({
      method: "POST",
      url: "/api/links",
      withCredentials: true,
      headers: {
        Authorization: `Bearer ${localStorage.accessToken}`
      },
      data: {links: {uri: uri, title: uri}},
    }).catch((e) => { Post.errors.append("Cannot save");console.log(e) })
  }
}

export default Post
