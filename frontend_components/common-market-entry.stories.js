import { html } from "lit"
import "./common-market-entry"

export default {
    title: 'Common Market Entry',
    argTypes: {
        market: { control: 'text' }
    }
}
const Template = ({ market }) => html`<common-market market=${market} ></common-market marets`;

export const Primary = Template.bind({})
Primary.storyName = "Market Entry"
Primary.args = {
    market: JSON.stringify({Name: "test", Photo: "https://firebasestorage.googleapis.com/v0/b/commonstock-dev.appspot.com/o/hollywood-market-peppers.jpg?alt=media&token=3d0bf207-2449-443d-90c7-95b8d2ed8bf7"})
}

