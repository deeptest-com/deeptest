import { TreeModelSettings, TreeStatus, FoldingType } from './tree.types';
import { RenamableNode } from './renamable.node';

export interface TreeModel {
  id: number;
  type: string;
  value: string | RenamableNode;
  pid?: number;
  children?: TreeModel[];
  settings?: TreeModelSettings;
  _status?: TreeStatus;
  _foldingType?: FoldingType;
}
