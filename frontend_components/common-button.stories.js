import { html } from 'lit'
import './common-button.js'

export default {
    title: 'Common Button',
    argTypes: {
        buttonType: {
        options: ['normal', 'submit'],
        control: 'select'
        }
    }
}

const Template = ({buttonText, buttonType}) => html`<common-button buttonText=${buttonText} buttonType=${buttonType}></common-button>`;

export const Primary = Template.bind({})
Primary.storyName = "Primary Button"
Primary.args = {
    buttonText: 'Normal Btn',
    buttonType: 'normal'
}
