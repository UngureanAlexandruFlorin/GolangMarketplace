import React from 'react';

class User extends React.Component {

    constructor(props) {
        super(props);
    }

    getAll() {
        fetch('http://ec2-54-219-132-254.us-west-1.compute.amazonaws.com:8082/get', {
                method: 'get',
                headers: {
                    'Access-Control-Request-Headers': 'Authorization',
                    'Authorization': localStorage.getItem('token')
                }
            })
            .then(response => {
                response.text()
                    .then(data => console.log(data))
                    .catch(error => console.log(error));
            })
            .catch(error => console.log(error));
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