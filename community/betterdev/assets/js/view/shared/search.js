const m = require('mithril')
import Post from '../../model/Post'

var query = ""
const Search = {
  view: (vnode) => {
    return m(".input-group.input-inline.has-icon-left", [
      m("input.form-input", {
        type: "text", placeholder: "search",
        oninput: m.withAttr("value", function(value) { query = value })
      }),
      m("i.form-icon.icon.icon-search"),
      m("button.btn.btn-primary.input-group-btn" + (Post.postStatus == "searching" ? ".loading" : ''), {
        onclick: (e) => {
          if (query != "") {
            return Post.search(query)
          }
        }
      }, "Search"),
      query == "" ? null : m("button.btn.btn-primary.input-group-btn", {onclick: () => { query = ""; Post.loadList()}}, "Reset")
    ])
  }
}

export { Search as default }
