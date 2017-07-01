import { Tree } from './tree';
import { RenamableNode } from './type/renamable.node';

export class NodeEvent {
  public constructor(public node: Tree) {
  }
}

export class NodeDraggingEvent {
  public constructor(public act: string) {

  }
}

export class NodeSelectedEvent extends NodeEvent {
  public constructor(node: Tree) {
    super(node);
  }
}

export class NodeDestructiveEvent extends NodeEvent {
  public constructor(node: Tree) {
    super(node);
  }
}

export class NodeMovedEvent extends NodeDestructiveEvent {
  public constructor(node: Tree, public srcTree: Tree, public options: any) {
    super(node);
  }
}
export class NodeMovedRemoteEvent extends NodeDestructiveEvent {
  public constructor(node: Tree, public srcTree: Tree, public options: any) {
    super(node);
  }
}

export class NodeRemovedEvent extends NodeDestructiveEvent {
  public constructor(node: Tree) {
    super(node);
  }
}
export class NodeRemovedRemoteEvent extends NodeDestructiveEvent {
  public constructor(node: Tree) {
    super(node);
  }
}

export class NodeCreatedEvent extends NodeDestructiveEvent {
  public constructor(node: Tree) {
    super(node);
  }
}
export class NodeCreatedRemoteEvent extends NodeDestructiveEvent {
  public constructor(node: Tree) {
    super(node);
  }
}

export class NodeRenamedEvent extends NodeDestructiveEvent {
  public constructor(node: Tree, public oldValue: string | RenamableNode, public newValue: string | RenamableNode) {
    super(node);
  }
}
export class NodeRenamedRemoteEvent extends NodeDestructiveEvent {
  public constructor(node: Tree, public oldValue: string | RenamableNode, public newValue: string | RenamableNode) {
    super(node);
  }
}
