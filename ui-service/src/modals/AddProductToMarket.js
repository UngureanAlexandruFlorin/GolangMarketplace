import React from 'react';
import './Modal.css';

class AddProductToMarket extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            name: '',
            description: '',
            price: 0
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

            if (!this.state.name || !this.state.description || !this.state.price) {
                alert('Please fill all fields!');
                return;
            }

            response = await fetch('http://192.168.1.6:8082/create', {
                method: 'post',
                headers: {
                    'Access-Control-Request-Headers': 'Authorization',
                    'Authorization': localStorage.getItem('token')
                },
                body: JSON.stringify({
                    name: this.state.name,
                    description: this.state.description,
                    price: +this.state.price
                })
            });

            if (response.status !== 200) {
                alert('Failed to create the product!');
            }
            this.props.changeModalState(false);
        } catch (error) {
            alert(error);
            this.props.changeModalState(false);
            return;
        }
    }

    render() {
        return (
            <div className="modal" onClick={() => { this.props.changeModalState(false); }}>
                <div className="modalContent" onClick={(event) => { event.stopPropagation(); }} >
                    <p className="modalTitle">Add a new product</p>
                    <form>
                        <input type="text" placeholder="Name" onChange={this.handleNameChanged} /><br />
                        <input type="text" placeholder="Description" onChange={this.handleDescriptionChanged} /><br />
                        <input type="number" placeholder="Price" min="0" onChange={this.handlePriceChanged} /><br />
                    </form>
                    <button className="modalButton" onClick={this.addProductToMarket}>Add product</button>
                </div>
            </div>
        )
    }
}

export default AddProductToMarket;