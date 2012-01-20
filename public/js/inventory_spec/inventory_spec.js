(function() {
  var app, _ref;

  app = (_ref = window.app) != null ? _ref : {};

  describe("Vinyl", function() {
    return describe("new vinyl", function() {
      beforeEach(function() {
        return this.vinyl = new app.Vinyl({
          title: 'Illmatic',
          size: 12,
          records: 1,
          year: 1994
        });
      });
      it("populates title", function() {
        return expect(this.vinyl.get('title')).toEqual('Illmatic');
      });
      return it("populates size", function() {
        return expect(this.vinyl.get('size')).toEqual(12);
      });
    });
  });

}).call(this);
