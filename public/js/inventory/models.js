(function() {
  var Vinyl, _ref,
    __hasProp = Object.prototype.hasOwnProperty,
    __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor; child.__super__ = parent.prototype; return child; };

  Vinyl = (function(_super) {

    __extends(Vinyl, _super);

    function Vinyl() {
      Vinyl.__super__.constructor.apply(this, arguments);
    }

    Vinyl.prototype.initialize = function(attributes, options) {
      if (!attributes.label) {
        return this.attributes.label = {
          name: "&nbsp;"
        };
      }
    };

    return Vinyl;

  })(Backbone.Model);

  this.app = (_ref = window.app) != null ? _ref : {};

  this.app.Vinyl = Vinyl;

}).call(this);
