app = window.app ? {}

describe "Vinyl", ->
  
  describe "new vinyl", ->
    beforeEach ->
      @vinyl = new app.Vinyl
        title: 'Illmatic'
        size: 12
        records: 1
        year: 1994

    it "populates title", ->
      expect(@vinyl.get('title')).toEqual 'Illmatic'
    it "populates size", ->
      expect(@vinyl.get('size')).toEqual 12