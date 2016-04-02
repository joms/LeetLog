import React from 'react';
import FlatButton from '../../../node_modules/material-ui/lib/flat-button';
import RaisedButton from '../../../node_modules/material-ui/lib/raised-button';
import Dialog from '../../../node_modules/material-ui/lib/dialog';

const styles = {
  container: {
    textAlign: 'center',
    paddingTop: 200,
  },
};

class Modal extends React.Component {
  constructor(props, context) {
    super(props, context);
    this.handleRequestClose = this.handleRequestClose.bind(this);
    this.handleTouchTap = this.handleTouchTap.bind(this);

    this.state = {
      open: false,
    };
  }

  handleRequestClose() {
    this.setState({
      open: false,
    });
  }

  handleTouchTap() {
    this.setState({
      open: true,
    });
  }

  render() {
    const standardActions = (
      <FlatButton
        label="Okey"
        secondary={true}
        onTouchTap={this.handleRequestClose}
      />
    );

    return (
        <div style={styles.container}>
          <Dialog
            open={this.state.open}
            title="Super Secret Password"
            actions={standardActions}
            onRequestClose={this.handleRequestClose}
          >
            1-2-3-4-5
          </Dialog>
          <h1>material-ui</h1>
          <h2>example project</h2>
          <RaisedButton
            label="Super Secret Password!"
            primary={true}
            onTouchTap={this.handleTouchTap}
          />
        </div>
    );
  }
}

export default Modal;
