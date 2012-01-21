class Vinyls extends Backbone.Collection
  model: app.Vinyl
  url: '/vinyls'
  byAuthor: ->
    _.sortBy @models, (vinyl) ->
      vinyl.get('author').name
  byTitle: ->
    _.sortBy @models, (task) ->
      task.get('title')

@app = window.app ? {}
# Make a new instance of the collection, don't call on the abstract definition
@app.Vinyls = new Vinyls