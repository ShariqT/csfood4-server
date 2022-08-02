import { html } from "lit";
import { CommonElement } from "./common-elem";
export class CommonHeader extends CommonElement {
    static get properties() {
        return {
            isLoggedIn: { type: String },
            numCartItems: { type: Number }
        }
    }

    constructor() {
        super()
        this.numCartItems = null;
    }

    render() {
        let menuItems = null
        let cartItem = null
        if (this.isLoggedIn === 'true') {
           
            menuItems = html`<a class="navbar-item" href="/logout">Logout</a>`
            
        } else {
            menuItems = html`<a class="navbar-item" href="/login">Login</a><a class="navbar-item" href="/signup">SignUp</a>`
        }

        if (this.isLoggedIn && this.numCartItems == 0) {
            cartItem = html`<a class="navbar-item" href="/cart">Cart</a>`
        }
        if (this.isLoggedIn && this.numCartItems > 0) {
            cartItem = html`<a class="navbar-item" href="/cart">Cart<span class="tag is-dark">${this.numCartItems}</span></a>`
        } 
        return html`
        <nav class="navbar is-primary" role="navigation" aria-label="main navigation">
            <div class="navbar-menu">
                <div class="navbar-start">
                    <a class="navbar-item" href="/">
                        Home
                    </a>
                </div>
                <div class="navbar-end">

                    ${menuItems}
                    ${cartItem}
                
                </div>
            </div>
        </nav>
        `
    }
}
window.customElements.define('common-header', CommonHeader);