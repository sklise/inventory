class Vinyls extends Backbone.Collection
  model: app.Vinyl
  url: '/vinyls'

@app = window.app ? {}
# Make a new instance of the collection, don't call on the abstract definition
@app.Vinyls = new Vinyls