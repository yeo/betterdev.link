const m = require('mithril')

import session from '../session'
import Url from '../util/url'

const Collection = {
  draftCollection: "",
  loaded: false,
  errors: [],
  list: [],

  loadList: () => {
    return m.request({
      method: "GET",
      url: "/api/collections",
      withCredentials: true,
      headers: {
        Authorization: `Bearer ${localStorage.accessToken}`
      }
    })
    .then(function(result) {
      if (result && result.data) {
        Collection.list = result.data
        Collection.loaded = true
      }
    })
    .catch((e) => {
      Collection.errors = ["Cannot fetch collections from server. Try again"]
    })
  },

  create: (e) => {
    if (Collection.draftCollection == "") {
      return
    }

    return m.request({
      method: "POST",
      url: `api/collections/`,
      data: {collection: {name: Collection.draftCollection }},
      headers: {
        Authorization: `Bearer ${localStorage.accessToken}`
      },
    }).then(function(result) {
      Collection.loadList()
    }).catch((e) => {
      Post.errors = [{message: "Cannot create collection"}]
      m.redraw()
    })

  },

  append: (link, collection) => {
    return m.request({
      method: "PATCH",
      url: `api/collections/${collection.id}`,
      data: {link: link.id},
      headers: {
        Authorization: `Bearer ${localStorage.accessToken}`
      },
    }).then(function(result) {
      console.log(result)
    }).catch((e) => {
      console.log(e)
    })
  }
}

export default Collection
