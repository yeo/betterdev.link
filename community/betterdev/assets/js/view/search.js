const m = require('mithril')
import Post from '../model/Post'

var query = ""

const Search = {
  view: () => {
    return m(".input-group.input-inline.has-icon-left", [
      m("input.form-input", {
        type: "text", placeholder: "search",
        oninput: m.withAttr("value", function(value) { query = value })
      }),
      m("i.form-icon.icon.icon-search"),
      m("button.btn.btn-primary.input-group-btn", {
        onclick: (e) => {
          if (query != "") {
            Post.search(query)
          }
        }
      }, "Search")
    ])
  }
}

export { Search as default }
