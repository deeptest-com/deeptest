export enum NodeMenuItemAction {
  NewFolder,
  NewTag,
  NewFolderInner,
  NewTagInner,
  Rename,
  Remove
}

export enum NodeMenuAction {
  Close
}

export interface NodeMenuEvent {
  sender: HTMLElement;
  action: NodeMenuAction;
}

export interface NodeMenuItemSelectedEvent {
  nodeMenuItemAction: NodeMenuItemAction;
}
