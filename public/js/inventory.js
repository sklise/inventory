(function() {

  jQuery(function() {
    new app.AppView({
      collection: app.Vinyls
    });
    return app.Vinyls.fetch();
  });

}).call(this);
