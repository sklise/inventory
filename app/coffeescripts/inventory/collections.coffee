class Vinyls extends Backbone.Collection
  model: app.Vinyl
  url: '/vinyls'

@app = window.app ? {}
@app.Vinyls = new Vinyls