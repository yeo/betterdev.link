// Brunch automatically concatenates all files in your
// watched paths. Those paths can be configured at
// config.paths.watched in "brunch-config.js".
//
// However, those files will only be executed if
// explicitly imported. The only exception are files
// in vendor, which are never wrapped in imports and
// therefore are always executed.

// Import dependencies
//
// If you no longer want to use a dependency, remember
// to also remove its path from "config.paths.watched".
import "phoenix_html"

// Import local files
//
// Local files can be imported directly using relative
// paths "./socket" or full ones "web/static/js/socket".

// import socket from "./socket"
const m = require('mithril')
import profileView from './view/profile'
import appView from './view/app'
import postFormView from './view/shared/postform'
//import mylinksView from './view/mylinks'

//m.mount(document.getElementById("app-wrapper"), appView)
//m.mount(document.getElementById("app-postform"), postFormView)
m.mount(document.getElementById("app-nav"), profileView)

m.route(document.getElementById("app-wrapper"), "/", {
  "/": appView, // defines `http://localhost/#!/home`
//  //"/mylinks": mylinksView, // defines `http://localhost/#!/home`
})
