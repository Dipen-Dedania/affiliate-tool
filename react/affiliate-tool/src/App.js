import React, {Component} from 'react';
import './App.css';
import {Segment} from 'semantic-ui-react';
import Mathew from './images/matthew.png'

class App extends Component {
    constructor(){
        super();
        this.state = {
            searchString : ""
        };
        this.triggerSearch = this.triggerSearch.bind(this);
        this.searchStringChange = this.searchStringChange.bind(this);
        this.searchResult = this.searchResult.bind(this);
    }

    triggerSearch(){
        console.log('DEBUG', '17 triggerSearch', this.state.searchString);
    }

    searchStringChange(event){
        this.setState({ searchString : event.target.value});
    }

    searchResult(event){
        if (event.key === "Enter") {
            this.triggerSearch();
        }
    }

    getCardView(){
        let cardView = [];
        cardView.push(<div className="ui card" key="1">
            <img src={Mathew} className="ui image" alt="preview"/>
            <div className="content">
                <div className="header">Iphone</div>
                <div className="meta">
                    <span className="date">Apple Inc</span>
                </div>
                <div className="description">Sample description about the Apple Inc company</div>
            </div>
            <div className="extra content">
                <a>
                    <i aria-hidden="true" className="user icon"></i>Get Link</a>
            </div>
        </div>);
        return cardView
    }

    render() {
        return (
            <div className="App">
                <div className="ui container">
                    <Segment>
                        <h1>Affiliation Tool</h1>
                    </Segment>
                    <div className="ui icon input width-100">
                        <input type="text" placeholder="Search..." value={this.state.searchString}
                               onChange={this.searchStringChange} onKeyPress={this.searchResult}/>
                        <i aria-hidden="true" className="search circular link icon" onClick={this.triggerSearch}></i>
                    </div>
                    <hr/>
                    {this.state.searchString !== "" && this.getCardView()}
                </div>
            </div>
        );
    }
}

export default App;
