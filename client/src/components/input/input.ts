import { InputProps } from './input.types';

export class Input {
	private props: InputProps;

	constructor(props: InputProps) {
		this.props = props;
	}

	render(root: HTMLElement) {
		const input = document.createElement('input');

		input.name = this.props.name;
		input.min = this.props.min.toString();
		input.type = this.props.type;

		if (this.props.value !== undefined) {
			input.value = this.props.value.toString();
		}

		if (this.props.disabled) {
			input.disabled = this.props.disabled;
		}

		root.appendChild(input);
	}
}
