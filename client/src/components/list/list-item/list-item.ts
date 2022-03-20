import { ListItemProps } from './list-item.types';

export class ListItem {
	private props: ListItemProps;

	constructor(props: ListItemProps) {
		this.props = props;
	}

	render(root: HTMLElement) {
		const listItem = document.createElement('li');

		listItem.textContent = this.props.text;

		root.appendChild(listItem);
	}
}
