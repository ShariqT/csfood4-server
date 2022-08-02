import { LitElement, html } from "lit";
import { bulmaStyles } from "./bulma.css";
export class CommonElement extends LitElement {
    static get styles() {
        return [bulmaStyles]
    }
}