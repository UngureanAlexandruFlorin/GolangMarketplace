import React from 'react';

class User extends React.Component {

    constructor(props) {
        super(props);

        this.state = { data: '' };

        this.getAll = this.getAll.bind(this);
    }

    async getAll() {
        let response;
        try {
            response = await fetch('http://ec2-54-219-132-254.us-west-1.compute.amazonaws.com:8082/get', {
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
        		<button onClick = {this.getAll}>Get all</button>

        		<p>Result: {this.state.data}</p>
        	</div>
        );
    }
}

export default User;