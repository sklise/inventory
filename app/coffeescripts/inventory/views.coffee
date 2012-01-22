jQuery ->
  _.templateSettings = {
      interpolate : /\{\{([\s\S]+?)\}\}/g
  }

  class AppView extends Backbone.View
    el: '#wrap'
    initialize: (options) ->
      @collection.bind 'reset', @render, @
      @subviews = [
        new MenuView      collection: @collection
        new NewVinylView  collection: @collection
        new VinylsView    collection: @collection
        ]
    render: ->
      $(@el).empty()
      $(@el).append subview.render().el for subview in @subviews
      @

  class MenuView extends Backbone.View
    tagName: 'header'
    template: _.template($('#menu-template').html())
    render: ->
      $(@el).html @template()
      @

  class VinylsView extends Backbone.View
    tagName: 'ul'
    initialize: (options) ->
      @collection.bind 'add', @render, @
      @collection.bind 'remove', @render, @
      @collection.bind 'change', @render, @
    render: ->
      $(@el).empty()
      $(@el).append @collection.models.length
      for vinyl in @collection.byAuthor()
        vinylView = new VinylView model: vinyl
        $(@el).append vinylView.render().el
      @

  class VinylView extends Backbone.View
    tagName: 'li'
    className: 'vinyl'
    template: _.template($('#vinyl-template').html())
    editTemplate: _.template($('#vinyl-edit-template').html())
    events:
      'click .delete': 'destroy'
      'click .edit'  : 'edit'
      'keypress .vinyl-edit-form' : 'saveOnEnter'
    render: ->
      $(@el).html @template(@model.toJSON())
      @
    edit: ->
      $(@el).html @editTemplate(@model.toJSON())
      @$('.vinyl-edit-form').eq(0).focus()
      @
    destroy: ->
      @model.destroy()
    saveOnEnter: ->
      if(event.keyCode is 13)
        @model.save({
          author:{
            name:@$('.author input').val().trim().trim()
            id:@model.get('author').id
          }
          label:{
            name:@$('.label input').val().trim().trim()
          }
          title: @$('.title input').val().trim().trim()
          records: @$('.record input').val().trim().trim()
          year: @$('.year input').val().trim().trim()
        })
        @render()

  class NewVinylView extends Backbone.View
    id: 'new-vinyl'
    tagName: 'form'
    template: _.template($('#new-vinyl-template').html())
    events:
      'keypress .vinyl-form': 'saveOnEnter'
    render: ->
      $(@el).html @template()
      @
    saveOnEnter: (event) ->
      if (event.keyCode is 13) # ENTER
        event.preventDefault()
        newAttributes = {
          vinyl: {
            title: $('#new-vinyl').find('[name="title"]').val().trim()
            year: $('#new-vinyl').find('[name="year"]').val().trim()
            size: parseInt $('#new-vinyl').find('[name="size"]').val().trim()
            records: parseInt $('#new-vinyl').find('[name="records"]').val().trim()
          }
          author: {
            name: $('[name="author"]').val().trim()
          }
          label: {
            name: $('[name="label"]').val().trim()
          }
        }
        if @collection.create(newAttributes)
          @emptyFields()
          @focus()
    emptyFields: ->
      $('.vinyl-form').val('')
      $('#new-vinyl').find('[name="size"]').val("12")
      $('#new-vinyl').find('[name="records"]').val("1")
      true
    focus: ->
      $('#new-vinyl').find('[name="title"]').focus()
      

  @app = window.app ? {}
  @app.AppView = AppView