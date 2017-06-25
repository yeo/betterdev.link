const m = require('mithril')
import session from '../../session'

const Flash = {
  view: () => {
    return session.flash().map(m => {
      console.log(m)
      return m('toast.toast-primary', m.message)
    })
  }
}

export { Flash as default }
