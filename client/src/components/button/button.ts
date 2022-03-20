import { ButtonProps } from './button.types';

export class Button {
	private props: ButtonProps;

	constructor(props: ButtonProps) {
		this.props = props;
	}

	render(root: HTMLElement) {
		const button = document.createElement('button');

		button.textContent = this.props.text;

		if (this.props.type) {
			button.type = this.props.type;
		}

		root.appendChild(button);
	}
}
