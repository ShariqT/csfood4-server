import { html } from "lit"
import "./common-cart"

export default {
    title: 'Common Cart',
    argTypes: {
        items: { control: 'text' }
    }
}
const Template = ({items}) => html`<common-cart items=${items}></common-cart>`;

export const Primary = Template.bind({})
Primary.storyName = "Cart"
Primary.args = {
    items: JSON.stringify([{Product: "Apple", Quantity: 3, Price: 12.00}])
}
