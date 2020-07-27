import React from 'react';
import ReactDOM from 'react-dom';
import AddProductToMarket from './modals/AddProductToMarket'

class User extends React.Component {

    constructor(props) {
        super(props);

        this.state = { data: '' };

        this.getAll = this.getAll.bind(this);
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

    render() {
        return (
            <div>
                <h1>User page!</h1>
                <button onClick={this.getAll}>Get all</button>

                <p>Result: {this.state.data}</p>

                <div id="addProductToMarkedId"></div>
            </div>
        );
    }
}

ReactDOM.render(
    <AddProductToMarket />,
    document.getElementById('addProductToMarkedId')
);

export default User;