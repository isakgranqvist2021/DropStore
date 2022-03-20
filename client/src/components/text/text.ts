export class Text {
	private variant: string = 'p';

	constructor(variant?: string) {
		if (variant) this.variant = variant;
	}

	render(root: HTMLElement, text: string) {
		const textElement = document.createElement(this.variant);

		textElement.textContent = text;

		root.appendChild(textElement);
	}
}
