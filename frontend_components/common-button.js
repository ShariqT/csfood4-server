import { html } from "lit";
import { CommonElement } from "./common-elem";
export class CommonButton extends CommonElement {
    
    static get properties() {
        return {
            buttonType: {type: String},
            buttonText: {type: String}
        }
    }

    constructor() {
        super();
        this.buttonType = 'normal'
        this.buttonText = 'Example'
    }

    render() {
        if (this.buttonType === 'submit') {
            return html`
                <input type="submit" value="${this.buttonText}" class="button is-primary" />
            `
        }
        return html`
        <button class="button">${this.buttonText}</button>
        `;
    }
}

window.customElements.define('common-button', CommonButton);