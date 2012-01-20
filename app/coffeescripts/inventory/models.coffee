class Vinyl extends Backbone.Model
  url: ->
    if @id
      "/vinyls/#{@id}"
    else
      "/vinyls"

@app = window.app ? {}
@app.Vinyl = Vinyl