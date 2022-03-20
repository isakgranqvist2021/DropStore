import { Button } from '../button';
import { Input } from '../input';

export class Form {
	render(root: HTMLElement) {
		const form = document.createElement('form');

		form.method = 'POST';
		form.action = '/checkout';

		new Input({
			min: '1',
			value: '1',
			type: 'number',
			name: 'quantity',
		}).render(form);

		new Button({
			text: 'Buy Now',
		}).render(form);

		root.appendChild(form);
	}
}
