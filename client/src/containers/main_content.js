import React, { Component } from 'react';
import {Card, CardActions, CardHeader, CardText, CardTitle} from 'material-ui/Card';
import {Tabs, Tab} from 'material-ui/Tabs';
import { Divider } from 'material-ui';
import Empleados from './sections/empleados';
import Salarios from './sections/salarios';
import Departamentos from './sections/departamentos';

class MainContent extends Component {
    constructor(props){
        super(props);
    }
    handleActive(){
        
    }
    render(){
        return(
            <div>
                <Card
                    style={{marginTop: '15px'}}
                >
                    <CardHeader
                     />
                     <CardActions>
                     </CardActions>
                     <CardText>
                         <CardTitle title='Quanta Dynamics' subtitle='Towards the foundation'  className='text-center'/>
                         <Divider/>
                         <Tabs>
                         <Tab label="Empleados" >
                            <div>
                                <Empleados />
                            </div>
                            </Tab>
                            <Tab label="Salarios" >
                            <div>
                                <Salarios />
                            </div>
                            </Tab>
                            <Tab
                            label="Departamentos"
                            data-route="/home"
                            onActive={this.handleActive}
                            >
                            <div>
                                <Departamentos />
                            </div>
                            </Tab>
                         </Tabs>
                     </CardText>
                </Card>
            </div>
        )
    }
}

export default MainContent;