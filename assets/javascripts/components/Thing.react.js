var React = require("react"),
  ReactPropTypes = React.PropTypes,
  cx = require('react/lib/cx');

var Thing = React.createClass({
  propTypes: {
    thing: ReactPropTypes.object.isRequired
  },

  // This function loads the initial values.
  getInitialState: function () {
    return {
      isEditing: false
    };
  },

  render: function () {
    var thing = this.props.thing;

    var input
    if (this.state.isEditing) {
      input =
        <ThingInput
          className="edit"
          onSave={this._onSave}
          value={thing.name}
        />;
    }

    return (
      <li
        class
    )
  }
})