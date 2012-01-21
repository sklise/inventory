(function() {
  var Vinyls, _ref,
    __hasProp = Object.prototype.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor; child.__super__ = parent.prototype; return child; };

  Vinyls = (function(_super) {

    __extends(Vinyls, _super);

    function Vinyls() {
      Vinyls.__super__.constructor.apply(this, arguments);
    }

    Vinyls.prototype.model = app.Vinyl;

    Vinyls.prototype.url = '/vinyls';

    Vinyls.prototype.byAuthor = function() {
      return _.sortBy(this.models, function(vinyl) {
        return vinyl.get('author').name;
      });
    };

    Vinyls.prototype.byTitle = function() {
      return _.sortBy(this.models, function(task) {
        return task.get('title');
      });
    };

    return Vinyls;

  })(Backbone.Collection);

  this.app = (_ref = window.app) != null ? _ref : {};

  this.app.Vinyls = new Vinyls;

}).call(this);
