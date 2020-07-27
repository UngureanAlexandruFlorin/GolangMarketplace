import React from 'react';

class AddProductToMarket extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            name: '',
            description: '',
            price: ''
        };

        this.addProductToMarket = this.addProductToMarket.bind(this);
        this.handleNameChanged = this.handleNameChanged.bind(this);
        this.handleDescriptionChanged = this.handleDescriptionChanged.bind(this);
        this.handlePriceChanged = this.handlePriceChanged.bind(this);

    }

    handleNameChanged(event) {
        this.setState({ name: event.target.value })
    }

    handleDescriptionChanged(event) {
        this.setState({ description: event.target.value })
    }

    handlePriceChanged(event) {
        this.setState({ price: event.target.value })
    }

    async addProductToMarket() {
        let response;
        try {
            response = await fetch('http://192.168.1.13:8082/create', {
                method: 'post',
                headers: {
                    'Access-Control-Request-Headers': 'Authorization',
                    'Authorization': localStorage.getItem('token')
                }
            });

            if (response.status !== 200) {
                alert('Failed to create the product!');
            }
        } catch (error) {
            alert(error);
            return;
        }
    }

    render() {
        return (
            <div>
                <div className="panel">
                    <p></p>
                    <form>
                        <input type="text" placeholder="Name" onChange={this.handleNameChanged} /><br />
                        <input type="text" placeholder="Description" onChange={this.handleDescriptionChanged} /><br />
                        <input type="number" placeholder="Price" min="0" onChange={this.handlePriceChanged} /><br />
                    </form>
                    <button onClick={this.addProductToMarket}>Add product</button>
                </div>
            </div>
        )
    }
}

export default AddProductToMarket;