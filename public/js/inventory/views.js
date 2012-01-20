(function() {
  var __hasProp = Object.prototype.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor; child.__super__ = parent.prototype; return child; };

  jQuery(function() {
    var AppView, MenuView, NewVinylView, VinylView, VinylsView, _ref;
    _.templateSettings = {
      interpolate: /\{\{([\s\S]+?)\}\}/g
    };
    AppView = (function(_super) {

      __extends(AppView, _super);

      function AppView() {
        AppView.__super__.constructor.apply(this, arguments);
      }

      AppView.prototype.el = '#wrap';

      AppView.prototype.initialize = function(options) {
        this.collection.bind('reset', this.render, this);
        return this.subviews = [
          new MenuView({
            collection: this.collection
          }), new VinylsView({
            collection: this.collection
          }), new NewVinylView({
            collection: this.collection
          })
        ];
      };

      AppView.prototype.render = function() {
        var subview, _i, _len, _ref;
        $(this.el).empty();
        _ref = this.subviews;
        for (_i = 0, _len = _ref.length; _i < _len; _i++) {
          subview = _ref[_i];
          $(this.el).append(subview.render().el);
        }
        return this;
      };

      return AppView;

    })(Backbone.View);
    MenuView = (function(_super) {

      __extends(MenuView, _super);

      function MenuView() {
        MenuView.__super__.constructor.apply(this, arguments);
      }

      MenuView.prototype.tagName = 'header';

      MenuView.prototype.template = _.template($('#menu-template').html());

      MenuView.prototype.render = function() {
        $(this.el).html(this.template());
        return this;
      };

      return MenuView;

    })(Backbone.View);
    VinylsView = (function(_super) {

      __extends(VinylsView, _super);

      function VinylsView() {
        VinylsView.__super__.constructor.apply(this, arguments);
      }

      VinylsView.prototype.tagName = 'table';

      VinylsView.prototype.render = function() {
        var vinyl, vinylView, _i, _len, _ref;
        $(this.el).empty();
        _ref = this.collection.models;
        for (_i = 0, _len = _ref.length; _i < _len; _i++) {
          vinyl = _ref[_i];
          console.log("vinyl", vinyl);
          vinylView = new VinylView({
            model: vinyl
          });
          $(this.el).append(vinylView.render().el);
        }
        return this;
      };

      return VinylsView;

    })(Backbone.View);
    VinylView = (function(_super) {

      __extends(VinylView, _super);

      function VinylView() {
        VinylView.__super__.constructor.apply(this, arguments);
      }

      VinylView.prototype.tagName = 'tr';

      VinylView.prototype.template = _.template($('#vinyl-template').html());

      VinylView.prototype.render = function() {
        $(this.el).html(this.template(this.model.toJSON()));
        return this;
      };

      return VinylView;

    })(Backbone.View);
    NewVinylView = (function(_super) {

      __extends(NewVinylView, _super);

      function NewVinylView() {
        NewVinylView.__super__.constructor.apply(this, arguments);
      }

      NewVinylView.prototype.id = 'new-vinyl';

      NewVinylView.prototype.tagName = 'form';

      NewVinylView.prototype.template = _.template($('#new-vinyl-template').html());

      NewVinylView.prototype.events = {
        'keypress .vinyl-form': 'saveOnEnter'
      };

      NewVinylView.prototype.render = function() {
        $(this.el).html(this.template());
        return this;
      };

      NewVinylView.prototype.saveOnEnter = function(event) {
        var newAttributes;
        if (event.keyCode === 13) {
          console.log("hi");
          console.log($('#new-vinyl').find('[name="title"]').val());
          newAttributes = {
            title: "I See A Darkness"
          };
          console.log(newAttributes);
          console.log(this.collection);
          this["new"] = new app.Vinyl({
            title: "I See A Darkness"
          });
          return this["new"].save();
        }
      };

      NewVinylView.prototype.focus = function() {
        return $('#new-vinyl').find('[name="title"]').focus();
      };

      return NewVinylView;

    })(Backbone.View);
    this.app = (_ref = window.app) != null ? _ref : {};
    return this.app.AppView = AppView;
  });

}).call(this);
