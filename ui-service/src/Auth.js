import React from 'react';
import './Auth.css';

class Auth extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            email: '',
            password: ''
        };

        this.handleLogin = this.handleLogin.bind(this);
        this.handleRegister = this.handleRegister.bind(this);
        this.handleEmailChange = this.handleEmailChange.bind(this);
        this.handlePasswordChange = this.handlePasswordChange.bind(this);
    }

    handleLogin(event) {
        if (this.state.email.length < 2 || this.state.password < 2) {
            alert('Email or password too short!');
        } else {
            fetch('http://ec2-54-219-132-254.us-west-1.compute.amazonaws.com:8081/login', {
                method: 'post',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    email: this.state.email,
                    password: this.state.password
                })
            }).then(data => {
                console.log(data);
                data.text()
                    .then(token => { console.log(token); })
                    .catch(err => { console.log(err); });
            }).catch(error => {
                alert('Error!');
                console.log(error);
            });
        }
    }

    handleRegister(event) {
        alert('Register');
    }

    handleEmailChange(event) {
        this.setState({ email: event.target.value });
    }

    handlePasswordChange(event) {
        this.setState({ password: event.target.value });
    }

    render() {
        return <div id = "authPanel">
        <h1>Golang Marketplace</h1>
          <form>
            <input className = "authInput" type = "text" placeholder = "Email" onChange={ this.handleEmailChange } /> <br/>
            <input className = "authInput" type = "password"  placeholder = "Password" onChange={ this.handlePasswordChange }/>
          </form>

          <div className = "authButtons">
            <button className = "authButton" onClick = {this.handleLogin}>Login</button> <br/>
            <button className = "authButton" onClick = {this.handleRegister}>Register</button>
          </div>
        </div>;
    }

}

export default Auth;
