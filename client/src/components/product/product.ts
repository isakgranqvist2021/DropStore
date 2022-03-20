import { ProductBody } from './product-body';
import { Thumbnail } from './thumbnail';

const productBodyProps = {
	title: 'Amazing Shoes',
	description:
		'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum is simply dummy text of the printing and typesetting industry.Lorem Ipsum is simply dummy text of the printing and typesetting industry',
	price: '200 SEK',
};

export class Product {
	render(root: HTMLElement) {
		const product = document.createElement('div');

		new Thumbnail().render(product);
		new ProductBody(productBodyProps).render(product);

		root.appendChild(product);
	}
}
