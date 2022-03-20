import { ListItem } from './list-item';
import { ListProps } from './list-types';

export class List {
	private props: ListProps;

	constructor(props: ListProps) {
		this.props = props;
	}

	render(root: HTMLElement) {
		const list = document.createElement('ul');

		this.props.listItems.forEach((text) => {
			new ListItem({ text }).render(list);
		});

		root.appendChild(list);
	}
}
