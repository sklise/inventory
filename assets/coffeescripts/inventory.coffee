@app = window.app ? {}

jQuery ->
  new app.AppView collection: app.Vinyls
  app.Vinyls.fetch()