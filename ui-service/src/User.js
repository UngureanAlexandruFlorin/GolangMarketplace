import React from 'react';
import AddProductToMarket from './modals/AddProductToMarket'
import './User.css';

class User extends React.Component {
    constructor(props) {
        super(props);

        this.state = { data: '', output: '' };

        this.getAll = this.getAll.bind(this);
        this.getAllByEmail = this.getAllByEmail.bind(this);
        this.changeModalState = this.changeModalState.bind(this);
    }

    async getAll() {
        let response;
        try {
            response = await fetch('http://192.168.1.13:8082/get', {
                method: 'get',
                headers: {
                    'Access-Control-Request-Headers': 'Authorization',
                    'Authorization': localStorage.getItem('token')
                }
            });
        } catch (error) {
            alert(error);
            return;
        }

        this.setState({
            data: await response.text()
        });
    }

    async getAllByEmail() {
        let response;
        try {
            response = await fetch('http://192.168.1.13:8082/getBtEmail', {
                method: 'get',
                headers: {
                    'Access-Control-Request-Headers': 'Authorization',
                    'Authorization': localStorage.getItem('token')
                }
            });
        } catch (error) {
            alert(error);
            return;
        }

        this.setState({
            output: await response.text()
        });
    }

    changeModalState(state) {
        console.log(state);
        this.setState({ showModal: state });
    }

    render() {
        return (
            <div>
                <h1>User page!</h1>
                <button onClick={this.getAll}>Get all</button>

                <p>Result: {this.state.data}</p>

                <button onClick={() => { this.changeModalState(true) }}> Add product to market </button> <br />
                <button onClick={() => { this.changeModalState(true) }}> Show my published products </button>

                <div id="outputPanel">
                    <p>Output</p>
                    {this.state.output}
                </div>

                {this.state.showModal ? <AddProductToMarket changeModalState={this.changeModalState} /> : null}
            </div >
        );
    }
}

export default User;