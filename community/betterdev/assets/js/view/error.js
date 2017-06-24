const ErrorView = {
  view: (vnode) => {
    return vnode.attrs.errors.map((e) => {
      m('.toast.toast-primary', [
        m('button.btn.btn-clear'),
        e
      ])
    })
  }
}

export { ErrorView as default }
