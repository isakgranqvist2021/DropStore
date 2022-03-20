import { Product } from './components';

window.onload = () => {
	const root = document.getElementById('root');

	new Product().render(root);
};
