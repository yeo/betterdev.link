const m = require('mithril')

import session from '../session'
import Url from '../util/url'

const Post = {
  errors: [],
  list: [],
  pagination: [],

  params : {},

  loadList: () => {
    Post.postStatus = 'loading'

    return m.request({
      method: "GET",
      url: "/api/links" + Url.serialize(Post.params),
      withCredentials: true,
      headers: {
        Authorization: `Bearer ${localStorage.accessToken}`
      }
    })
    .then(function(result) {
      if (result && result.data) {
        Post.list = result.data
        Post.pagination = result.pagination
      }
    })
    .catch((e) => { console.log("Error", e) })
  },

  postStatus: "init",
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
      data: {link: {uri: uri}},
    }).catch((e) => { Post.errors.append("Cannot save");console.log(e) })
  }
}

export default Post
