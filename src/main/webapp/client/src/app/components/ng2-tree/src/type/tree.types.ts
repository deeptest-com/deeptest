import * as _ from 'lodash';

import { TreeModel } from './tree.model';

export class FoldingType {
  public static Expanded: FoldingType = new FoldingType('node-expanded fa');
  public static Collapsed: FoldingType = new FoldingType('node-collapsed fa');
  public static Leaf: FoldingType = new FoldingType('node-leaf fa');

  public constructor(private _cssClass: string) {
  }

  public get cssClass(): string {
    return this._cssClass;
  }
}

export class TreeModelSettings {
  /**
   * "static" property when set to true makes it impossible to drag'n'drop tree or call a menu on it.
   * @name TreeModelSettings#static
   * @type boolean
   * @default false
   */

  public constructor(staticTree?: boolean) {
  }

  public static merge(sourceA: TreeModel, sourceB: TreeModel): TreeModelSettings {
    return _.defaults({}, _.get(sourceA, 'settings'), _.get(sourceB, 'settings'), {staticTree: false});
  }
}

export enum TreeStatus {
  New,
  Modified,
  IsBeingRenamed
}
