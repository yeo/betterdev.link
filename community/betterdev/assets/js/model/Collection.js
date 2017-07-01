const m = require('mithril')

import session from '../session'
import Url from '../util/url'

const Collection = {
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

  append: (link, collection) => {
    return m.request({
      method: "PATCH",
      url: `api/collections/${collection.id}`,
      data: {link_id: link.id},
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
