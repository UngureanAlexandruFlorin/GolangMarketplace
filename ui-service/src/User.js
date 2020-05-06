import React from 'react';

class User extends React.Component {

    constructor(props) {
        super(props);
    }

    getAll() {
        fetch('ec2-54-219-132-254.us-west-1.compute.amazonaws.com:8082/get', {
                method: 'post',
                headers: {
                    'Content-Type': 'application/json',
                    'Accept': '*/*',
                    'Authorization': localStorage.getItem('token')
                },
                body: JSON.stringify({
                    email: this.state.email,
                    password: this.state.password
                })
            })
            .then()
            .catch();
    }

    render() {
        return (
            <div>
        		<h1>User page!</h1>
        		<button onClick = {this.getAll}>Get all</button>
        	</div>
        );
    }
}

export default User;