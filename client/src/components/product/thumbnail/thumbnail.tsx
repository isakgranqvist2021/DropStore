export class Thumbnail {
	render(root: HTMLElement) {
		const thumbnail = document.createElement('div');

		thumbnail.className = 'product_thumbnail';

		root.appendChild(thumbnail);
	}
}
