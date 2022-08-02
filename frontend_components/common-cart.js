import { html } from "lit";
import { CommonElement } from "./common-elem";
export class CommonCart extends CommonElement {

    static properties = {
        items: { type: String, reflect: true },
        _totalCost: { state: true}

    }
    constructor(){
        super();
        this._totalCost = 0.0
    }
    render(){
        let itemHTML = []
        let itemJSON = JSON.parse(this.items)
        for(const element of itemJSON){
            let itemRow = html`<tr>
            <form class="form" action="/update_cart" method="POST>
                    <input type="hidden" name="product_id" value="${element.ID}" />
                    <td>${element.Product}</td>
                    <td><input class="input" type="number" value="${element.Quantity}" name="quantity" /></td>
                    <td>${parseFloat(element.Price)}</td>
                    <td>
                    
                        <input type="submit" class="button is-primary is-small" value="Edit" />
                        </form> 
            <form class="form" action="/update_cart" method="POST">
                        <input type="submit" class="button is-small is-danger" value="Remove" />
            </form>
                           </tr>
                `
            itemHTML.push(itemRow)

        }
        console.log(itemHTML)
        return html`
        <table class="table is-fullwidth is-hoverable">
            <thead>
                <tr>
                    <th>Item</th>
                    <th>Quantity</th>
                    <th>Price</th>
                    <th></th>
                </tr>
            </thead>
            <tbody>
                ${itemHTML}
            </tbody>
        </table>
        `
    }
}

window.customElements.define('common-cart', CommonCart);