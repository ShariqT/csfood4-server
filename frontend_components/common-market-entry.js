import { html } from "lit";
import { CommonElement } from "./common-elem";
export class CommonMarket extends CommonElement {
    static get properties() {
        return {
            market: { type: String },
        }
    }
    

    render() {
        let marketInfo = null
        if (this.market !== ""){
            let marketJson = JSON.parse(this.market)
            marketInfo = html`
            <div class="columns">
                <div class="column is-one-third">  
                    <img src="${marketJson.Photo}" />
                </div>
                <div class="column">
                    <h1 class="is-size-1">${marketJson.Name}</h1>
                </div>
            </div>
            `
        } else {
            marketInfo = html`<p>No data available</p>`
        }

        return html`
            <div class="card">
                <div class="card-content">
                    <div class="content">
                        ${marketInfo}
                    </div>
                </div>
            </div>
        `
    }
}
window.customElements.define('common-market', CommonMarket);