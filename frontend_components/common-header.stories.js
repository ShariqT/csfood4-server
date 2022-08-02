import { html } from "lit"
import './common-header'

export default {
    title: 'Common Header',
    argTypes: {
        isLoggedIn: { control: 'boolean'},
        numCartItems: { control: 'range', min:0, step:1, max:30}
    }
}

const Template = ({isLoggedIn, numCartItems}) => html`<common-header isloggedin=${isLoggedIn} numcartitems=${numCartItems}></common-header>`;

export const Primary = Template.bind({})
Primary.storyName = "Header"
Primary.args = {
    isLoggedIn: false,
    numCartItems: 0
}