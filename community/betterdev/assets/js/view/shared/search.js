const m = require('mithril')
import Post from '../../model/Post'

var query = ""
const Search = {
  view: (vnode) => {
    return m(".input-group.input-inline", [
      m("input.form-input.input-group-addon", {
        type: "text", placeholder: "search",
        oninput: m.withAttr("value", (value) => { query = value })
      }),
      m("button.btn.btn-primary.input-group-btn" + (Post.postStatus == "searching" ? ".loading" : ''), {
        onclick: (e) => {
          if (query != "") {
            Post.params["q"] = query
            return Post.loadList()
          } else {
            delete Post.params.q
            return Post.loadList()
          }
        }
      }, "Search"),
      query == "" ? null : m("button.btn.btn-primary.input-group-btn", {onclick: () => { delete Post.params.q; Post.loadList(); query = ""; }}, "Reset")
    ])
  }
}

export { Search as default }
