import { Text } from '../../text';

import { ProductBodyProps } from './product-body.types';

export class ProductBody {
	private props: ProductBodyProps;

	constructor(props: ProductBodyProps) {
		this.props = props;
	}

	render(root: HTMLElement) {
		const productBody = document.createElement('div');

		productBody.className = 'product_body';

		new Text('h2').render(productBody, this.props.title);
		new Text('p').render(productBody, this.props.description);
		new Text('h4').render(productBody, this.props.price.toString());

		root.appendChild(productBody);
	}
}
