var React = require("react"),
  ReactPropTypes = React.PropTypes,
  Thing = require("./Thing.react");

var Things = React.createClass({
  propTypes: {
    allThings: ReactPropTypes.object.isRequired
  },

  render: function () {

    var allThings = this.props.allThings;
    var things = []
    for (var i in allThings) {
      things.push(<Thing key={i} thing={allThings[i]} />);
    }

    return (
      <section id="things">
        <ul id="things-list">
        {things}
        </ul>
      </section>
    )
  }
});

module.exports = Things;