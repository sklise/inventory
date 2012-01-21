class Vinyl extends Backbone.Model
  initialize: (attributes, options) ->
    if !attributes.label
      @attributes.label = {name: "&nbsp;"}

@app = window.app ? {}
@app.Vinyl = Vinyl