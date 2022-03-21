import { Form } from '../form';
import { List } from '../list';

import { ProductBody } from './product-body';
import { Thumbnail } from './thumbnail';
import { PRODUCT_BODY_PROPS, LIST_ITEMS } from './product.constants';

export class Product {
	render(root: HTMLElement) {
		const product = document.createElement('div');

		product.className = 'product';

		new Thumbnail().render(product);

		new ProductBody(PRODUCT_BODY_PROPS).render(product);

		new List({ listItems: LIST_ITEMS }).render(product);

		new Form().render(product);

		root.appendChild(product);
	}
}
