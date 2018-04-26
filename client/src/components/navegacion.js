import React, { Component } from 'react';
import AppBar from 'material-ui/AppBar';
import Drawer from 'material-ui/Drawer';
import MenuItem from 'material-ui/MenuItem';
import Paper from 'material-ui/Paper';
import Menu from 'material-ui/Menu';
import Divider from 'material-ui/Divider';
import Avatar from 'material-ui/Avatar';

const style = {margin: 5, width: '100px'};

class Navegacion extends Component {
    constructor(props){
        super(props);
        this.state = { open: false};
    }
    clickMenu(){
        this.setState({open: !this.state.open});
    }
    render(){
        return(
            <div>
                 <AppBar
                    title="TestApp"
                    iconClassNameRight="muidocs-icon-navigation-expand-more"
                    onLeftIconButtonClick={this.clickMenu.bind(this)}
                />
                 <Drawer
                    docked={false} 
                    open={this.state.open}
                    onRequestChange={(open) => this.setState({open})}>
                        <Avatar
                            src="https://www.releasemama.com/wp-content/uploads/2018/01/dr-manhattans-child-the-secret-to-dcs-watchmen-sequel.jpg"
                            size={100}
                            style={style}
                            />
                            <h3>Dr. Manhattan</h3>
                        <Divider />
                        <MenuItem>Empleados</MenuItem>
                        <MenuItem>Plantillas de salario</MenuItem>
                        <MenuItem>Departamentos</MenuItem>
                </Drawer>
            </div>
        );
    }
}

export default Navegacion;